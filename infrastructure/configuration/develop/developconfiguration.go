// Package develop providers develop configuration
package develop

import (
	"os"
	"time"

	"lolesports/infrastructure/application"
	"lolesports/infrastructure/application/entities"
)

const (
	basePath = "/back/lol_hub"

	// rest
	maxIdleConnsPerHost = 50
	defaultTimeout      = 180 * time.Second

	// Database connection
	dbReadUsername        = "devlolhub_RPROD"
	dbReadPassword        = "DB_MONGO_DEVLOLHUB_RPROD"
	dbWriteUsername       = "devlolhub_WPROD"
	dbWritePassword       = "DB_MONGO_DEVLOLHUB_WPROD"
	dbHost                = "DB_MONGO_DEVLOLHUB_ENDPOINT"
	dbName                = "devlolhub"
	dbPort                = "DB_MONGO_DEVLOLHUB_PORT"
	dbMaxIdleConns        = 25
	dbMaxOpenConns        = 50
	dbConnMaxLifetimeSecs = 300
	dbReadTimeout         = 10
	dbWriteTimeout        = 10
	dbTimeout             = 10

	riotBaseURL = "https://%s.api.riotgames.com"
)

//

type develop struct {
}

//

// NewDevelopConfiguration develop configuration
func NewDevelopConfiguration() *develop {
	return &develop{}
}

// BasePath returns the url base path of the application
func (*develop) BasePath() string {
	return basePath
}

// MaxIdleConnsPerHost returns max idle conns per host for rest service
func (*develop) MaxIdleConnsPerHost() int {
	return maxIdleConnsPerHost
}

// RiotURL returns Base URL for RiotRest container
func (*develop) RiotURL() string {
	return riotBaseURL
}

// DefaultTimeout rest deffault timeout
func (*develop) DefaultTimeout() time.Duration {
	return defaultTimeout
}

func (*develop) DBUserName() string {
	switch application.Context().Role() {
	case entities.ReadRole:
		return os.Getenv(dbReadUsername)
	case entities.WriteRole, entities.WorkerRole:
		return os.Getenv(dbWriteUsername)
	default:
		return ""
	}
}

func (*develop) DBPassword() string {
	switch application.Context().Role() {
	case entities.ReadRole:
		return os.Getenv(dbReadPassword)
	case entities.WriteRole, entities.WorkerRole:
		return os.Getenv(dbWritePassword)
	default:
		return ""
	}
}

func (*develop) DBHost() string {
	return os.Getenv(dbHost)
}

func (*develop) DBName() string {
	return os.Getenv(dbName)
}

func (*develop) DBPort() string {
	return os.Getenv(dbPort)
}
