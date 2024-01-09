package dependencies

import (
	"github.com/Trafilea/internal/core/ports"
	"github.com/Trafilea/internal/core/service/numbersrv"
	"github.com/Trafilea/internal/handlers/numberhdl"
	"github.com/Trafilea/internal/repositories/numberrepo"
	"os"
)

type Definition struct {

	//
	// Repositories
	//

	NumberRepository ports.NumberRepository

	//
	// Core
	//

	NumberService ports.NumberService

	//
	// Handlers
	//

	NumberHandler *numberhdl.Handler
}

func NewByEnvironment() Definition {

	//
	// Obtains the environment
	//

	environment := "staging"
	scope := os.Getenv("GO_ENVIRONMENT")

	switch scope {
	case "production":
		environment = "production"
	}

	//
	// Obtains configs based on environment
	//

	_config := configs[environment]

	//
	// Initializes clients
	//

	storageClient := initStorage(_config.storage["number-repository"])

	//
	// Initializes dependencies
	//

	d := initDependencies(
		storageClient,
	)

	return d
}

func initDependencies(
	storageClient map[int]string,

) Definition {

	d := Definition{}

	//
	// Repositories
	//

	d.NumberRepository = numberrepo.New(storageClient)

	//
	// Core
	//

	d.NumberService = numbersrv.New(d.NumberRepository)

	//
	// Handlers
	//

	d.NumberHandler = numberhdl.New(d.NumberService)

	return d
}
