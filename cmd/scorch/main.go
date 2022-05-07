package main

import (
	"fmt"
	"log"

	"github.com/alwindoss/scorch"
	"github.com/alwindoss/scorch/internal/engine"
	"github.com/caarlos0/env/v6"
)

func main() {
	cfg := &scorch.Config{}
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	err := engine.Run(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
