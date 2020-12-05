package window

type Config struct {
	Debug bool
	Frameless bool
}

func DefaultConfig() Config {
	return Config{
		Debug: false,
		Frameless: false,
	}
}
