// Package application configuration application
package application

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"lolesports/infrastructure/application/entities"
)

func Test_InitContext_InvalidScope(t *testing.T) {
	// Given
	ass := assert.New(t)

	originalScope := os.Getenv("SCOPE")
	defer os.Setenv("SCOPE", originalScope)

	os.Setenv("SCOPE", "invalid_scope")
	prepareEnvironment()

	// Then
	ass.PanicsWithValue("invalid scope invalid_scope", func() { InitContext() })
}

func Test_InitContext_InvalidEnvironment(t *testing.T) {
	// Given
	ass := assert.New(t)

	originalScope := os.Getenv("SCOPE")
	defer os.Setenv("SCOPE", originalScope)

	os.Setenv("SCOPE", "invalid_environment-read")
	prepareEnvironment()

	// Then
	ass.PanicsWithValue("invalid environment invalid_environment", func() { InitContext() })
}

func Test_InitContext_InvalidRole(t *testing.T) {
	// Given
	ass := assert.New(t)

	originalScope := os.Getenv("SCOPE")
	defer os.Setenv("SCOPE", originalScope)

	os.Setenv("SCOPE", "production-invalid_role")
	prepareEnvironment()

	// Then
	ass.PanicsWithValue("invalid role invalid_role", func() { InitContext() })
}

func Test_InitContext_Fail_ApplicationID(t *testing.T) {
	// Given
	ass := assert.New(t)

	originalScope := os.Getenv("SCOPE")
	defer os.Setenv("SCOPE", originalScope)

	os.Setenv("SCOPE", "production-worker")
	prepareEnvironment()

	// Then
	ass.PanicsWithValue("application_id cannot be obtained from secrets", func() { InitContext() })
}

func Test_InitContext_Success(t *testing.T) {
	// Given
	ass := assert.New(t)

	originalScope := os.Getenv("SCOPE")
	defer os.Setenv("SCOPE", originalScope)

	os.Setenv("SCOPE", "production-worker")
	prepareEnvironment()

	// Then
	ass.NotPanics(func() { InitContext() })
}

func Test_Context_WithoutInit_Panics(t *testing.T) {
	// Given
	ass := assert.New(t)

	originalScope := os.Getenv("SCOPE")
	defer os.Setenv("SCOPE", originalScope)

	os.Setenv("SCOPE", "production-worker")
	prepareEnvironment()

	// Then
	ass.PanicsWithValue("application context not initialized", func() { Context() })
}

func Test_Environment(t *testing.T) {
	// Given
	ass := assert.New(t)

	originalScope := os.Getenv("SCOPE")
	defer os.Setenv("SCOPE", originalScope)

	os.Setenv("SCOPE", "production-worker")
	prepareEnvironment()

	// When
	InitContext()
	environment := Context().Environment()

	// Then
	ass.EqualValues(entities.ProductionEnvironment, environment)
}

func Test_Role(t *testing.T) {
	// Given
	ass := assert.New(t)

	originalScope := os.Getenv("SCOPE")
	defer os.Setenv("SCOPE", originalScope)

	os.Setenv("SCOPE", "production-worker")
	prepareEnvironment()

	// When
	InitContext()
	role := Context().Role()

	// Then
	ass.EqualValues(entities.WorkerRole, role)
}

func Test_ApplicationID(t *testing.T) {
	// Given
	ass := assert.New(t)

	originalScope := os.Getenv("SCOPE")
	defer os.Setenv("SCOPE", originalScope)

	os.Setenv("SCOPE", "production-worker")
	prepareEnvironment()

	// When
	applicationID := Context().ApplicationID()

	// Then
	ass.Equal("1234567890", applicationID)
}

func prepareEnvironment() {
	contextInitialized = false

	applicationID = ""
	environment = ""
	role = ""
}

func Test_InitTestContext_Success(t *testing.T) {
	// Given
	ass := assert.New(t)

	// When
	prepareEnvironment()
	InitTestContext()
	applicationContext := Context()

	// Then
	ass.Equal("test", string(applicationContext.Environment()))
	ass.Equal("worker", string(applicationContext.Role()))
	ass.Equal("1234567890", applicationContext.ApplicationID())
	ass.Equal("0.0.0", applicationContext.Version())
}

func Test_InitCustomTestContext_Success(t *testing.T) {
	// Given
	ass := assert.New(t)

	// When
	prepareEnvironment()
	InitCustomTestContext(entities.Environment("custom_environment"), entities.Role("custom_role"), "custom_application_id", "custom_version")
	applicationContext := Context()

	// Then
	ass.Equal("custom_environment", string(applicationContext.Environment()))
	ass.Equal("custom_role", string(applicationContext.Role()))
	ass.Equal("custom_application_id", applicationContext.ApplicationID())
	ass.Equal("custom_version", applicationContext.Version())
}

func Test_IsLocal_False(t *testing.T) {
	// Given
	ass := assert.New(t)

	// When
	prepareEnvironment()
	_ = os.Setenv("LOCAL_RUN", "")

	// Then
	ass.Equal(false, IsLocal())
}

func Test_IsLocal_True(t *testing.T) {
	// Given
	ass := assert.New(t)

	// When
	prepareEnvironment()
	_ = os.Setenv("LOCAL_RUN", "true")

	// Then
	ass.Equal(true, IsLocal())
}
