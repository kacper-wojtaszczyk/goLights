package music

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/zmb3/spotify"
	"log"
	"net/http"
	"os"
)

var (
	auth  = spotify.NewAuthenticator(os.Getenv("SPOTIFY_REDIRECT_URI"), spotify.ScopeUserReadPrivate, spotify.ScopeUserReadCurrentlyPlaying, spotify.ScopeUserReadPlaybackState)
	ch    = make(chan *spotify.Client)
	state = uuid.NewString()
)
var refreshToken string

type SpotifyClient struct {
	client *spotify.Client
	state  string
	auth   spotify.Authenticator
}

func (client SpotifyClient) GetCurrentTrackAttributes() *spotify.AudioFeatures {
	return client.getTrackAudioAnalysis(client.getCurrentTrack().ID)
}

func (client SpotifyClient) Pause() {
	err := client.client.Pause()
	panic(err)
}

func (client SpotifyClient) getCurrentTrack() *spotify.FullTrack {
	country := "LV"
	playerState, err := client.client.PlayerStateOpt(&spotify.Options{Country: &country})
	if nil != playerState {
		return playerState.Item
	}
	fmt.Println(err)
	panic(err)
}

func (client SpotifyClient) getTrackAudioAnalysis(id spotify.ID) *spotify.AudioFeatures {
	audioAnalysis, _ := client.client.GetAudioFeatures(id)

	return audioAnalysis[0]
}

func CreateSpotifyClient() SpotifyClient {
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go http.ListenAndServe(":8080", nil)
	auth  = spotify.NewAuthenticator(os.Getenv("SPOTIFY_REDIRECT_URI"), spotify.ScopeUserReadPrivate, spotify.ScopeUserReadCurrentlyPlaying, spotify.ScopeUserReadPlaybackState)

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)

	return SpotifyClient{client: client, auth: auth, state: state}
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	refreshToken = tok.RefreshToken
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}