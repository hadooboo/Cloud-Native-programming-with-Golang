package main

import (
	"flag"
	"log"

	"jaehonam.com/ev/apiserver"
	"jaehonam.com/ev/config"
	"jaehonam.com/ev/database/dblayer"
)

var (
	configPath = flag.String("config", `./application.yaml`, "flag to set the path to the configuration yaml file")
)

func main() {
	flag.Parse()

	log.Println("parsing config file...")
	config, err := config.NewConfig(*configPath)
	if err != nil {
		panic(err)
	}

	log.Println("initiating database handler...")
	databaseHandler, err := dblayer.NewDatabaseLayer(&config.Database)
	if err != nil {
		panic(err)
	}

	log.Println("starting api server...")
	log.Fatal(apiserver.Serve(&config.Apiserver, databaseHandler))
}
