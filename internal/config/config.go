package config

// Config ...
type Config struct {
	HTTPAddr string `env:"HTTP_ADDR" envDefault:":8888"`
}
