package main

import (
	"flag"
	"fmt"
	"uri-shortener/cmd/modes"
)

func main() {
	app := modes.AppMode{}

	pathToConfig := flag.String("path-to-config", "", "Path to config file")
	configFilename := flag.String("config-file", "", "Config filename")
	flag.Parse()
	fmt.Println("path to config =", *pathToConfig, "config file name =", *configFilename)

	err := app.ParseConfig(*pathToConfig, *configFilename)
	if err != nil {
		return
	}

	app.Run()
}
