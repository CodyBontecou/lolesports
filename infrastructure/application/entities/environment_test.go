// Package entities providers entities_test
package entities_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"lolesports/api/infrastructure/application/entities"
)

func Test_EnvironmentContains_NilArray_NotFound(t *testing.T) {
	// given
	ass := assert.New(t)
	environments := make(entities.Environments, 0)
	develop := entities.DevelopEnvironment

	// when
	found := environments.Contains(develop)

	// then
	ass.False(found)
}

func Test_EnvironmentContains_Found(t *testing.T) {
	// given
	ass := assert.New(t)
	environments := mockedEnvironments()
	develop := entities.DevelopEnvironment

	// when
	found := environments.Contains(develop)

	// then
	ass.True(found)
}

func Test_EnvironmentContains_NotFound(t *testing.T) {
	// given
	ass := assert.New(t)
	environments := mockedEnvironments()
	invalid := entities.Environment("invalid")

	// when
	found := environments.Contains(invalid)

	// then
	ass.False(found)
}

func mockedEnvironments() entities.Environments {
	environments := make(entities.Environments, 3)
	environments[0] = entities.ProductionEnvironment
	environments[1] = entities.SandboxEnvironment
	environments[2] = entities.DevelopEnvironment

	// done
	return environments
}
