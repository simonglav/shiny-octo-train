package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/simonglav/http-rest-api/internal/app/apiserver"
)

// Path to config via flag in cmd
var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// Run server
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
