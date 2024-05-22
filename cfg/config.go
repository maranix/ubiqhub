package cfg

import (
	"fmt"
	"os"
)

type Config struct {
	Env    string `json:"env"`
	Server Server `json:"server"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

const (
	host = "HOST"
	port = "PORT"
	env  = "ENV"
)

func create(env, addr, p string) *Config {
	return &Config{
		Env: env,
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
	env := envVar(env, args)
	host := envVar(host, args)
	port := envVar(port, args)

	return create(env, host, port)
}
