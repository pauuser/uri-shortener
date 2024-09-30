package cmd

import (
	"flag"
	"uri-shortener/cmd/modes"
)

func main() {
	app := modes.AppMode{}

	pathToConfig := flag.String("path-to-config", "", "Path to config file")
	configFilename := flag.String("config-file", "", "Config filename")
}
