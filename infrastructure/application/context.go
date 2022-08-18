// Package application configuration application
package application

import (
	"fmt"
	"os"
	"strings"

	"lolesports/infrastructure/application/entities"
)

type (
	// Contexter interface for application context
	Contexter interface {
		// ApplicationID getter for application id
		ApplicationID() string
		// Environment getter for environment
		Environment() entities.Environment
		// Role getter for role
		Role() entities.Role

		Operator() string
		// Version returns the build version of the application
		Version() string
	}

	context struct{}
)

var (
	contextInitialized bool

	applicationID string
	environment   entities.Environment
	role          entities.Role
	operator      string

	version string
)

const (
	applicationIDKey      string = "SECRET_APPLICATION_ID"
	applicationVersionEnv string = "VERSION"
)

// InitContext initialize application context
func InitContext() {
	if !contextInitialized {
		scope := os.Getenv("SCOPE")

		parts := strings.Split(strings.ToLower(scope), "-")
		if len(parts) <= 1 {
			panic(fmt.Sprintf("invalid scope %s", scope))
		}

		environment = entities.Environment(parts[0])
		if !entities.PossibleEnvironments.Contains(environment) {
			panic(fmt.Sprintf("invalid environment %s", environment))
		}

		role = entities.Role(parts[1])
		if !entities.PossibleRoles.Contains(role) {
			panic(fmt.Sprintf("invalid role %s", role))
		}

		applicationID = os.Getenv(applicationIDKey)
		if applicationID == "" {
			panic("application_id cannot be obtained from secrets")
		}

		operator = os.Getenv("OPERATOR")

		version = os.Getenv(applicationVersionEnv)

		contextInitialized = true
	}

	// done
}

// Context getter for application context
func Context() Contexter {
	if !contextInitialized {
		panic("application context not initialized")
	}

	return &context{}
}

func (*context) ApplicationID() string {
	return applicationID
}

func (*context) Environment() entities.Environment {
	return environment
}

func (*context) Role() entities.Role {
	return role
}

func (*context) Operator() string {
	return operator
}

func (*context) Version() string {
	return version
}

// InitCustomTestContext initialize application context for tests
func InitCustomTestContext(e entities.Environment, r entities.Role, a string, v string) {
	environment = e
	role = r
	applicationID = a
	version = v

	contextInitialized = true

	// done
}

// IsLocal used to allow local run for engine
func IsLocal() bool {
	return strings.EqualFold(os.Getenv("LOCAL_RUN"), "true")
}

// Operator used to allow local run for engine
func Operator() string {
	return os.Getenv("OPERATOR")
}
