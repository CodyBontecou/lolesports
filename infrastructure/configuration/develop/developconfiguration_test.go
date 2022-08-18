// Package develop providers develop configuration
package develop

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewDevelopConfiguration(t *testing.T) {
	ass := assert.New(t)

	configuration := NewDevelopConfiguration()
	os.Setenv("LOCAL_RUN", "")

	basePath := configuration.BasePath()
	ass.NotEmpty(basePath)

	maxIdleConnsPerHost := configuration.MaxIdleConnsPerHost()
	ass.NotZero(maxIdleConnsPerHost)

	defaultTimeout := configuration.DefaultTimeout()
	ass.NotZero(defaultTimeout)

	ass.NotEmpty(configuration.RiotURL())
}

func Test_NewDevelopConfiguration_Fury(t *testing.T) {
	ass := assert.New(t)
	configuration := NewDevelopConfiguration()

	// Fury config
	riotURL := configuration.RiotURL()
	ass.Equal(riotURL, riotBaseURL)
}
