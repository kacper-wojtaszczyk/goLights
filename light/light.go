package light

import "strconv"

type Event interface {
	EventType() string
	Value() string
}

type HueChanged struct {
	hue int
}

func (event HueChanged) EventType() string {
	return "HSBCOLOR1"
}

func (event HueChanged) Value() string {
	return strconv.Itoa(event.hue)
}

type TurnedOn struct {
}

func (event TurnedOn) EventType() string {
	return "POWER"
}

func (event TurnedOn) Value() string {
	return "ON"
}

type TurnedOff struct {
}

func (event TurnedOff) EventType() string {
	return "POWER"
}

func (event TurnedOff) Value() string {
	return "OFF"
}

type Light struct {
	name   string
	hue    int
	power  bool
	events []Event
}

func (light *Light) SetHue(hue int) {
	light.recordThat(HueChanged{hue: hue})
}

func (light *Light) IncrementHue(increment int) {
	light.recordThat(HueChanged{hue: light.hue + increment})
}

func (light *Light) TurnOn() {
	light.power = true
}

func (light *Light) TurnOff() {
	light.power = false
}

func (light Light) GetPower() bool {
	return light.power
}

func (light Light) GetHue() int {
	return light.hue
}

func (light Light) GetPowerString() string {
	if light.power {
		return "ON"
	}

	return "OFF"
}

func (light *Light) recordThat(event Event) {
	light.apply(event)
	light.events = append(light.events, event)
}

func (light *Light) apply(event Event) {
	switch e := event.(type) {
	case HueChanged:
		light.onHueChanged(e)
		break
	case TurnedOn:
		light.onTurnedOn()
		break
	case TurnedOff:
		light.onTurnedOff()
		break
	}
}

func (light *Light) onHueChanged(event HueChanged) {
	light.hue = event.hue % 360
}

func (light *Light) onTurnedOn() {
	light.power = true
}

func (light *Light) onTurnedOff() {
	light.power = false
}

func (light *Light) PopEvents() []Event {
	popped := light.events
	light.events = nil

	return popped
}

func (light Light) GetName() string {
	return light.name
}

func Create(name string) Light {
	return Light{hue: 0, power: false, name: name}
}
