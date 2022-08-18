// Package configuration providers configuration environment _test
package configuration_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"lolesports/api/infrastructure/application"
	"lolesports/api/infrastructure/application/entities"
	"lolesports/api/infrastructure/configuration"
)

func Test_GetConfiguration(t *testing.T) {
	ass := assert.New(t)

	application.InitCustomTestContext(entities.Environment("production"), entities.Role("write"), "custom_application_id", "custom_version")

	conf := configuration.GetConfiguration()
	ass.NotNil(conf)
	ass.Implements((*configuration.Configuration)(nil), conf)
}

func Test_Configuration_Sandbox(t *testing.T) {
	ass := assert.New(t)

	application.InitCustomTestContext(entities.Environment("sandbox"), entities.Role("write"), "custom_application_id", "custom_version")

	conf := configuration.GetConfiguration()

	ass.NotNil(conf)
	ass.Implements((*configuration.Configuration)(nil), conf)
}

func Test_Configuration_Dev(t *testing.T) {
	ass := assert.New(t)

	application.InitCustomTestContext(entities.Environment("develop"), entities.Role("write"), "custom_application_id", "custom_version")

	conf := configuration.GetConfiguration()

	ass.NotNil(conf)
	ass.Implements((*configuration.Configuration)(nil), conf)
}
