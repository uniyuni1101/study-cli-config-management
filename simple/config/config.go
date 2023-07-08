package config

type Theme string
type TickRate int

type Config struct {
	Theme    Theme
	TickRate TickRate
}

func NewDefault() Config {
	return Config{
		Theme:    "simple",
		TickRate: 20,
	}
}
