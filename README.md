a tiny learning project. using go to control an array of tasmota-based lightbulbs via mqtt
usage:
```
go run main.go command bulb1 bulb2 bulbN ...
```

where `bulb1`, `bulb2`, `bulbN` are the names of mqtt topics for appropriate devices and `command` is [something](https://github.com/kacper-wojtaszczyk/goLights/blob/7992cfb437eb24c31084956ba9954bbdc830909d/main.go#L25)

<img src="Y2Vr4gR.png" width="400">

- - [ ] extract the animations from main
  - - [ ] ideally into some comprehensible json or something
- - [ ] make spotify.PlayerState more JIT (spotify API doesn't provide callbacks afaik)
    - - [ ] fuck querying spotify, let's move playback mgmt here
        - - [ ] or maybe fuck spotify altogether and something like [supercollider](https://github.com/supercollider/supercollider) ? I have some previous experience with SC. Not a pleasant one, but also _you should try everything at least twice_.
- - [ ] optimise the MQTT traffic, because it's sometimes too much for the bulbs. the event approach actually paid off as only the changes are published, but the message could be better. e.g. setting saturation and Hue in the same message
    - - - - - [ ] note to self see whether it's not `mosquitto` on the rpi. 
