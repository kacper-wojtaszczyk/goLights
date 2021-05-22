package main

import (
	"github.com/joho/godotenv"
	"github.com/kacper-wojtaszczyk/goLights/light"
	"github.com/kacper-wojtaszczyk/goLights/mqtt"
	"os"
	"time"
)

func main() {
	action := os.Args[1]
	bulbNames := os.Args[2:]
	godotenv.Load(".env.local", ".env")
	client := mqtt.CreateMqttClient(os.Getenv("MQTT_BROKER_URI"), os.Getenv("MQTT_USERNAME"), os.Getenv("MQTT_PASSWORD"))
	var lights []light.Light
	for i := 0; i < len(bulbNames); i++ {
		lights = append(lights, light.Create(bulbNames[i]))
	}
	lightRepository := mqtt.NewLightRepository(client)
	switch action {
	case "rainbowRotate":
		rainbowRotate(lightRepository, lights...)
		break
	case "turnOff":
		turnOff(lightRepository, lights...)
		break
	case "warmWhite":
		warmWhite(lightRepository, lights...)
		break
	default:
		panic("unrecognised command")
	}
}

func warmWhite(repository light.Repository, lights ...light.Light) {
	for i := 0; i < len(lights); i++ {
		lights[i].TurnOn()
		lights[i].SetWhite(100)
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
	time.Sleep(time.Second)
}

func rainbowRotate(repository light.Repository, lights ...light.Light) {
	var hue = 0
	for i := 0; i < len(lights); i++ {
		lights[i].TurnOn()
		lights[i].SetHue(hue)
		hue += 72
	}
	repository.Publish(lights...)

	for {
		time.Sleep(time.Millisecond * 150)
		for i := 0; i < len(lights); i++ {
			lights[i].IncrementHue(5)
		}
		repository.Publish(lights...)
	}
}
