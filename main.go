package main

import (
	"github.com/joho/godotenv"
	"github.com/kacper-wojtaszczyk/goLights/light"
	"github.com/kacper-wojtaszczyk/goLights/mqtt"
	"github.com/kacper-wojtaszczyk/goLights/music"
	"os"
	"time"
)

func main() {
	action := os.Args[1]
	bulbNames := os.Args[2:]
	godotenv.Load(".env.local", ".env")
	client := mqtt.CreateMqttClient(os.Getenv("MQTT_BROKER_URI"), os.Getenv("MQTT_USERNAME"), os.Getenv("MQTT_PASSWORD"))
	spotify := music.CreateSpotifyClient(
		os.Getenv("SPOTIFY_CLIENT_ID"),
		os.Getenv("SPOTIFY_CLIENT_SECRET"),
		os.Getenv("SPOTIFY_REDIRECT_URI"),
		)
	var lights []light.Light
	for i := 0; i < len(bulbNames); i++ {
		lights = append(lights, light.Create(bulbNames[i]))
	}
	lightRepository := mqtt.NewLightRepository(client)
	switch action {
	case "rainbowRotate":
		rainbowRotate(lightRepository, lights...)
		break
	case "redshift":
		redshift(lightRepository, spotify, lights...)
		break
	case "turnOff":
		turnOff(lightRepository, lights...)
		break
	case "warmWhite":
		warmWhite(lightRepository, lights...)
		break
	default:
		redshift(lightRepository, spotify, lights...)
	}
}

func warmWhite(repository light.Repository, lights ...light.Light) {
	for i := 0; i < len(lights); i++ {
		lights[i].TurnOn()
		lights[i].SetWhite(10)
		lights[i].SetCT(450)
	}
	repository.Publish(lights...)
	time.Sleep(time.Second)
}

func turnOff(repository light.Repository, lights ...light.Light) {
	for i := 0; i < len(lights); i++ {
		lights[i].TurnOff()
	}
	repository.Publish(lights...)
}

func rainbowRotate(repository light.Repository, lights ...light.Light) {
	var hue = 0
	for i := 0; i < len(lights); i++ {
		lights[i].TurnOn()
		lights[i].SetHue(hue)
		lights[i].SetBrightness(33)

		hue += 72
	}
	repository.Publish(lights...)

	for {
		time.Sleep(time.Minute / 45)
		for i := 0; i < len(lights); i++ {
			lights[i].IncrementHue(144)
		}
		repository.Publish(lights...)
	}
}

func redshift(repository light.Repository, spotify music.SpotifyClient, lights ...light.Light) {
	attributes := spotify.GetCurrentTrackAttributes()
	tempo := int(attributes.Tempo)
	for i := 0; i < len(lights); i++ {
		lights[i].SetBrightness(100)
		lights[i].SetSat(100)
		lights[i].SetFadeSpeed(2)
	}
	shift := 0
	step := 15
	x := 0
	for {
		for i := 0; i < len(lights); i++ {
			lights[i].SetHue(shift + x*step)
			x = (x+1)%6
		}

		repository.Publish(lights...)
		shift = (shift + step) % 360
		time.Sleep(time.Minute/time.Duration(tempo))
	}
}
