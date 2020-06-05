// Based on https://github.com/linuxdeepin/go-lib

package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPasswordName(t *testing.T) {
	// test root user
	name1 := "root"
	password1, err := GetPasswordName(name1)
	assert.Equal(t, password1.Uid, uint32(0))
	assert.Equal(t, password1.Home, "/root")
	assert.Nil(t, err)

	// test invalid user
	name2 := "invalid"
	password2, err := GetPasswordName(name2)
	assert.Nil(t, password2)
	assert.NotNil(t, err)
}

func TestGetPasswordUid(t *testing.T) {
	// test root user
	uid1 := uint32(0)
	password1, err := GetPasswordUid(uid1)
	assert.Equal(t, password1.Name, "root")
	assert.Equal(t, password1.Home, "/root")
	assert.Nil(t, err)

	// test invalid user
	uid2 := uint32(65535)
	password2, err := GetPasswordUid(uid2)
	assert.Nil(t, password2)
	assert.NotNil(t, err)
}
