package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrypt(t *testing.T) {
	// test valid inputs
	password1 := "password1"
	salt1 := "aa"
	hash1, err := crypt(password1, salt1)
	assert.Equal(t, "aajfMKNH1hTm2", hash1)
	assert.Nil(t, err)

	// test invalid inputs
	password2 := "password2"
	salt2 := ""
	_, err = crypt(password2, salt2)
	assert.NotNil(t, err)
}

func TestGetPasswordHash(t *testing.T) {
	// TODO
}

func TestCheckPasswordHash(t *testing.T) {
	// test correct password
	password1 := "password1"
	hash1 := "xxj31ZMTZzkVA"
	correct1, err := CheckPasswordHash(password1, hash1)
	assert.True(t, correct1)
	assert.Nil(t, err)

	// test incorrect password
	password2 := "password2"
	hash2 := "yyj31ZMTZzkVA"
	correct2, err := CheckPasswordHash(password2, hash2)
	assert.False(t, correct2)
	assert.Nil(t, err)

	// test invalid inputs
	password3 := "password3"
	hash3 := ""
	_, err = CheckPasswordHash(password3, hash3)
	assert.NotNil(t, err)
}
