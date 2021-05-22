package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/kacper-wojtaszczyk/goLights/light"
)

type LightRepository struct {
	client mqtt.Client
}

func (repository LightRepository) Publish(lights ...light.Light) {
	for i := 0; i < len(lights); i++ {
		repository.processEvents(lights[i].GetName(), lights[i].PopEvents())
	}
}

func (repository LightRepository) sendMessage(deviceName string, command string, value string) {
	repository.client.Publish("cmnd/"+deviceName+"/"+command, 0, false, value)
}

func (repository LightRepository) processEvents(name string, events []light.Event) {
	for i := 0; i < len(events); i++ {
		repository.sendMessage(name, events[i].EventType(), events[i].Value())
	}
}

func NewLightRepository(client mqtt.Client) LightRepository {
	return LightRepository{client: client}
}
