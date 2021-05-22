package light

type Light struct {
	name  string
	hue   int
	power bool
}

func (light *Light) SetHue(hue int) {
	light.hue = hue % 360
}

func (light *Light) IncrementHue(increment int) {
	light.hue = (light.hue + increment) % 360
}

func (light *Light) TurnOn() {
	light.power = true
}

func (light *Light) TurnOff() {
	light.power = false
}

func (light Light) GetName() string {
	return light.name
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

func Create(name string) Light {
	return Light{name: name, hue: 0, power: false}
}
