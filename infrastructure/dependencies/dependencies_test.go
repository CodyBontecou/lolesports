// Package dependencies provider dependencies
package dependencies

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"lolesports/infrastructure/application"
	"lolesports/infrastructure/application/entities"
)

func Test_StartDependencies(t *testing.T) {
	ass := assert.New(t)
	application.InitTestContext()

	// Should start as is_local to avoid using the rest service on get configuration
	os.Setenv("LOCAL_RUN", "true")

	container := StartDependencies()
	testContainer(ass, container)
}

func Test_StartDependencies_ReadRole(t *testing.T) {
	ass := assert.New(t)

	application.InitCustomTestContext(entities.TestEnvironment, entities.WorkerRole, "1234567890", "0.0.1")

	// Should start as is_local to avoid using the rest service on get configuration
	os.Setenv("LOCAL_RUN", "true")

	container := StartDependencies()
	testContainer(ass, container)
}

func testContainer(ass *assert.Assertions, container *Container) {
	ass.NotNil(container)
	ass.NotNil(container.Clock())
	// ass.NotNil(container.ExecutionLock())
}
