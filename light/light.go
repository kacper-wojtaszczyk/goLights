package light

type Light struct {
	name   string
	hue    int
	white  int
	ct     int
	power  bool
	events []Event
}

func (light *Light) SetHue(hue int) {
	light.recordThat(HueChanged{hue: hue%360})
}

func (light *Light) IncrementHue(increment int) {
	light.recordThat(HueChanged{hue: (light.hue + increment)%360})
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
