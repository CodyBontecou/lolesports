// Package entities providers entities_test
package entities_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"lolesports/infrastructure/application/entities"
)

func Test_RolesContains_NilArray_NotFound(t *testing.T) {
	// given
	ass := assert.New(t)
	roles := make(entities.Roles, 0)
	read := entities.WorkerRole

	// when
	found := roles.Contains(read)

	// then
	ass.False(found)
}

func Test_RolesContains_Found(t *testing.T) {
	// given
	ass := assert.New(t)
	roles := mockedRoles()
	read := entities.WorkerRole

	// when
	found := roles.Contains(read)

	// then
	ass.True(found)
}

func Test_RolesContains_NotFound(t *testing.T) {
	// given
	ass := assert.New(t)
	roles := mockedRoles()
	invalid := entities.Role("invalid")

	// when
	found := roles.Contains(invalid)

	// then
	ass.False(found)
}

func mockedRoles() entities.Roles {
	roles := make(entities.Roles, 1)
	roles[0] = entities.WorkerRole

	// done
	return roles
}
