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