// Package configuration providers configuration environment
package configuration

import (
	"time"

	"lolesports/infrastructure/application"
	"lolesports/infrastructure/application/entities"
	"lolesports/infrastructure/configuration/develop"
	"lolesports/infrastructure/configuration/production"
)

// Configuration interface for configuration service
type Configuration interface {
	BasePath() string

	// rest
	MaxIdleConnsPerHost() int
	DefaultTimeout() time.Duration
	RiotURL() string

	// db
	DBUserName() string
	DBPassword() string
	DBHost() string
	DBName() string
	DBPort() string
}

//

// GetConfiguration returns the configuration service depending on the scope environment variable
func GetConfiguration() Configuration {
	var configuration Configuration

	switch application.Context().Environment() {
	case entities.ProductionEnvironment:
		configuration = production.NewProductionConfiguration()
	/*case entities.SandboxEnvironment:
	configuration = sandbox.NewSandboxConfiguration()*/
	case entities.DevelopEnvironment:
		fallthrough
	default:
		configuration = develop.NewDevelopConfiguration()
	}

	// done
	return configuration
}
