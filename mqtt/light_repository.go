package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/kacper-wojtaszczyk/goLights/light"
	"strconv"
)

type LightRepository struct {
	client mqtt.Client
}

func (repository LightRepository) Publish(lights ...light.Light) {
	for i := 0; i < len(lights); i++ {
		repository.refreshPower(lights[i])
		repository.refreshHue(lights[i])
	}
}

func (repository LightRepository) refreshPower(light light.Light) {
	repository.sendMessage(light.GetName(), "POWER", light.GetPowerString())
}

func (repository LightRepository) refreshHue(light light.Light) {
	repository.sendMessage(light.GetName(), "HSBCOLOR1", strconv.Itoa(light.GetHue()))
}

func (repository LightRepository) sendMessage(deviceName string, command string, value string) {
	repository.client.Publish("cmnd/"+deviceName+"/"+command, 0, false, value)
}

func NewLightRepository(client mqtt.Client) LightRepository {
	return LightRepository{client: client}
}
