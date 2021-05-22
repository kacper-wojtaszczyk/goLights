package light

type Repository interface {
	Publish(lights ...Light)
}
