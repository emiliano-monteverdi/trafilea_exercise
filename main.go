package main

import (
	"github.com/Trafilea/cmd/dependencies"
	"github.com/Trafilea/cmd/httpserver"
)

func main() {

	//
	// Creates dependencies
	//

	d := dependencies.NewByEnvironment()

	//
	// Injects dependencies
	//

	httpserver.Start(d)
}
