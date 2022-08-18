// Package production providers production configuration
package production

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
	dbReadUsername        = "prodlolhub_RPROD"
	dbReadPassword        = "DB_MONGO_PRODLOLHUB_RPROD"
	dbWriteUsername       = "prodlolhub_WPROD"
	dbWritePassword       = "DB_MONGO_PRODLOLHUB_WPROD"
	dbHost                = "DB_MONGO_PRODLOLHUB_ENDPOINT"
	dbName                = "prodlolhub"
	dbPort                = "DB_MONGO_PRODLOLHUB_PORT"
	dbMaxIdleConns        = 25
	dbMaxOpenConns        = 50
	dbConnMaxLifetimeSecs = 300
	dbReadTimeout         = 10
	dbWriteTimeout        = 10
	dbTimeout             = 10

	/*
		// execution lock
		executionLockName = "PROD_USER_LOCK"
		executionLockTTL  = 15

		// publisher
		topicDropoutNotificationCluster = "BIGQUEUE_TOPIC_PROD_DROPOUT_NOTIFICATION_CLUSTER_NAME"
		topicDropoutNotificationName    = "BIGQUEUE_TOPIC_PROD_DROPOUT_NOTIFICATION_TOPIC_NAME"

		topicCalendarCluster = "BIGQUEUE_TOPIC_PROD_CALENDAR_USER_DROPOUT_CLUSTER_NAME"
		topicCalendarName    = "BIGQUEUE_TOPIC_PROD_CALENDAR_USER_DROPOUT_TOPIC_NAME"

		topicKYCMetricsCluster = "BIGQUEUE_TOPIC_PROD_USER_RECOVERY_METRICS_CLUSTER_NAME"
		topicKYCMetricsName    = "BIGQUEUE_TOPIC_PROD_USER_RECOVERY_METRICS_TOPIC_NAME"
	*/
	riotBaseURL = "https://%s.api.riotgames.com"
)

//

type production struct {
}

//

// NewProductionConfiguration production configuration
func NewProductionConfiguration() *production {
	return &production{}
}

//

// BasePath returns the url base path of the application
func (*production) BasePath() string {
	return basePath
}

// MaxIdleConnsPerHost returns max idle conns per host for rest service
func (*production) MaxIdleConnsPerHost() int {
	return maxIdleConnsPerHost
}

// RiotURL returns Base URL for RiotRest container
func (*production) RiotURL() string {
	return riotBaseURL
}

// DefaultTimeout rest default timeout
func (*production) DefaultTimeout() time.Duration {
	return defaultTimeout
}

func (*production) DBUserName() string {
	switch application.Context().Role() {
	case entities.ReadRole:
		return os.Getenv(dbReadUsername)
	case entities.WriteRole, entities.WorkerRole:
		return os.Getenv(dbWriteUsername)
	default:
		return ""
	}
}

func (*production) DBPassword() string {
	switch application.Context().Role() {
	case entities.ReadRole:
		return os.Getenv(dbReadPassword)
	case entities.WriteRole, entities.WorkerRole:
		return os.Getenv(dbWritePassword)
	default:
		return ""
	}
}

func (*production) DBHost() string {
	return os.Getenv(dbHost)
}

func (*production) DBName() string {
	return os.Getenv(dbName)
}

func (*production) DBPort() string {
	return os.Getenv(dbPort)
}
