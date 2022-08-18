// Package production providers production configuration
package production

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewProductionConfiguration(t *testing.T) {
	ass := assert.New(t)

	configuration := NewProductionConfiguration()

	basePath := configuration.BasePath()
	ass.NotEmpty(basePath)

	maxIdleConnsPerHost := configuration.MaxIdleConnsPerHost()
	ass.NotZero(maxIdleConnsPerHost)

	defaultTimeout := configuration.DefaultTimeout()
	ass.NotZero(defaultTimeout)

	ass.NotEmpty(configuration.RiotURL())
}

func Test_NewProductionConfiguration_Fury(t *testing.T) {
	ass := assert.New(t)
	configuration := NewProductionConfiguration()

	// Fury config
	riotURL := configuration.RiotURL()
	ass.Equal(riotURL, riotBaseURL)
}
