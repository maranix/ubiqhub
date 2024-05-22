package cfg

import (
	"fmt"
	"os"
)

type Config struct {
	Server Server `json:"server"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

const (
	host = "HOST"
	port = "PORT"
)

func create(addr string, p string) *Config {
	return &Config{
		Server: Server{
			Host: addr,
			Port: p,
		},
	}
}

func envVar(key string, args map[string]string) string {
	if _, ok := args[key]; !ok {
		fmt.Fprintf(os.Stderr, "Could not find %s in env", key)
		os.Exit(1)
	}

	return args[key]
}

func FromArgs(args map[string]string) *Config {
	host := envVar(host, args)
	port := envVar(port, args)

	return create(host, port)
}
