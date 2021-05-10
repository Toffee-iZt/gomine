package main

import (
	"flag"

	"github.com/Toffee-iZt/gomine/server"
)

func main() {
	cfg := server.DefaultConfig

	flag.StringVar(&cfg.Host, "host", cfg.Host, "server binding address")
	flag.IntVar(&cfg.Port, "port", cfg.Port, "server binding port")
	flag.Parse()

	server.New(cfg).Start()
}
