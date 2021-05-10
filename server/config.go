package server

// Config is a server config.
type Config struct {
	Host string
	Port int
}

// DefaultConfig ...
var DefaultConfig = Config{
	"0.0.0.0",
	25565,
}
