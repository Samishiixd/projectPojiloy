package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Samishiixd/http_rest_api/internal/app/apiserver.go"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.Decodefile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)

	if err := Start(); err != nil {
		log.Fatal(err)
	}
}
