package main

import (
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/simplehttpserver/internal/runner"
)

func main() {
	// Parse the command line flags and read config files
	options := runner.ParseOptions()
	runner, err := runner.New(options)
	if err != nil {
		gologger.Fatal().Msgf("Could not create runner: %s\n", err)
	}

	runner.Run()
	runner.Close()
}
