package main

import (
	"fmt"
	"io"
	"os"
)

type Config struct {
	OutStream io.Writer
}

type Option func(cfg *Config) error

func WithOutputStream(stream io.Writer) Option {
	return func(cfg *Config) error {
		cfg.OutStream = stream
		return nil
	}
}

func NewConfig(opts ...Option) (Config, error) {
	c := Config{OutStream: os.Stdout}
	for _, opt := range opts {
		err := opt(&c)
		if err != nil {
			return Config{}, err
		}
	}
	return c, nil
}

func app(words []string, cfg Config) {
	for _, w := range words {
		if len(w)%2 == 0 {
			fmt.Fprintf(cfg.OutStream, "word %s is even\n", w)
		} else {
			fmt.Fprintf(cfg.OutStream, "word %s is odd\n", w)
		}
	}
}

func main() {
	words := os.Args[1:]
	config, err := NewConfig(WithOutputStream(os.Stdout))
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error in creating config %v", err)
	}
	app(words, config)
}
