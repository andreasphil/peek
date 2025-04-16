package cli

import (
	"flag"

	"github.com/charmbracelet/log"
)

type Config struct {
	Filename    string
	Port        string
	AllowUnsafe bool
}

func ParseFlags() *Config {
	port := flag.String("port", "8080", "the port for serving the application")
	allowUnsafe := flag.Bool("allow-unsafe", true, "render inline HTML")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatalf("filename is required")
	} else if len(args) > 1 {
		log.Warn("ignored additional arguments", "args", args[1:])
	}

	return &Config{
		Filename:    args[0],
		Port:        *port,
		AllowUnsafe: *allowUnsafe,
	}
}
