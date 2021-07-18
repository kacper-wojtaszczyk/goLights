package music

import (
	"context"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"log"
)

type SpotifyClient struct {
	client spotify.Client
}

func (client SpotifyClient) GetCurrentTrackAttributes() *spotify.AudioFeatures {
	//return client.getTrackAudioAnalysis(client.getCurrentTrack().ID).Track
	return client.getTrackAudioAnalysis("1bnEQF8BBhei6RmlSAjHSl")
}

func (client SpotifyClient) getCurrentTrack() *spotify.FullTrack {
	playerState, err := client.client.PlayerCurrentlyPlaying()
	if nil != playerState {
		return playerState.Item
	}

	panic(err)
}

func (client SpotifyClient) getTrackAudioAnalysis(id spotify.ID) *spotify.AudioFeatures {
	audioAnalysis, _ := client.client.GetAudioFeatures(id)

	return audioAnalysis[0]
}

func CreateSpotifyClient(clientId string, clientSecret string, redirectUri string) SpotifyClient {
	config := &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	return SpotifyClient{client: spotify.Authenticator{}.NewClient(token)}
}
