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

type WhiteChanged struct {
	white int
}

func (event WhiteChanged) EventType() string {
	return "WHITE"
}

func (event WhiteChanged) Value() string {
	return strconv.Itoa(event.white)
}

type CTChanged struct {
	ct int
}

func (event CTChanged) EventType() string {
	return "CT"
}

func (event CTChanged) Value() string {
	return strconv.Itoa(event.ct)
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
	white  int
	ct     int
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
	light.recordThat(TurnedOn{})
}

func (light *Light) TurnOff() {
	light.recordThat(TurnedOff{})
}

func (light *Light) SetCT(ct int) {
	light.recordThat(CTChanged{ct: ct})
}

func (light *Light) SetWhite(white int) {
	light.recordThat(WhiteChanged{white: white})
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
	case WhiteChanged:
		light.onWhiteChanged(e)
		break
	case CTChanged:
		light.onCtChanged(e)
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

func (light *Light) onWhiteChanged(event WhiteChanged) {
	light.white = event.white
}

func (light *Light) onCtChanged(event CTChanged) {
	light.ct = event.ct
}

func Create(name string) Light {
	return Light{name: name, hue: 0, power: false, white: 0, ct: 0}
}
