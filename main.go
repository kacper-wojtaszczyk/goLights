package main

import (
	"github.com/joho/godotenv"
	"github.com/kacper-wojtaszczyk/goLights/light"
	"github.com/kacper-wojtaszczyk/goLights/mqtt"
	"os"
	"time"
)

func main() {
	bulbNames := os.Args[1:]
	godotenv.Load(".env.local", ".env")
	client := mqtt.CreateMqttClient(os.Getenv("MQTT_BROKER_URI"), os.Getenv("MQTT_USERNAME"), os.Getenv("MQTT_PASSWORD"))
	var lights []light.Light
	for i := 0; i < len(bulbNames); i++ {
		lights = append(lights, light.Create(bulbNames[i]))
	}
	lightRepository := mqtt.NewLightRepository(client)
	rainbowRotate(lightRepository, lights...)
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
