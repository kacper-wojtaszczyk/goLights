package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func CreateMqttClient(uri string, username string, password string) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(uri)
	opts.SetClientID(username)
	opts.SetUsername(username)
	opts.SetPassword(password)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}
