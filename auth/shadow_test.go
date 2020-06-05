// Based on https://github.com/linuxdeepin/go-lib

package auth

import (
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetShadowPasswordName(t *testing.T) {
	// test root user
	name := "root"
	spassword, err := GetShadowPasswordName(name)

	// `getspnam()` requires root privileges
	if err != nil {
		assert.Equal(t, err, syscall.EACCES)
		assert.Nil(t, spassword)
	} else {
		assert.Equal(t, spassword.Name, "root")
	}
}
