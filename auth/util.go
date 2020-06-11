package auth

// #cgo LDFLAGS: -lcrypt
// #include <stdlib.h>
// #include <crypt.h>
import "C"
import (
	"errors"
	"os"
	"unsafe"
)

// Go wrapper around `crypt()`
func crypt(key string, salt string) (string, error) {
	keyC := C.CString(key)
	saltC := C.CString(salt)
	defer C.free(unsafe.Pointer(keyC))
	defer C.free(unsafe.Pointer(saltC))

	hashC := C.crypt(keyC, saltC)
	if hashC == nil {
		return "", errors.New("Unable to encrypt password.")
	}
	return C.GoString(hashC), nil
}

// Get password hash associated to current user
func GetPasswordHash() (string, error) {
	uid := uint32(os.Getuid())
	password, err := GetPasswordUid(uid)
	if err != nil {
		return "", err
	}
	hash := password.Password

	// password hash in /etc/shadow
	if hash == "x" {
		spassword, err := GetShadowPasswordName(password.Name)
		if err != nil {
			return "", err
		}
		hash = spassword.Password
	}
	return hash, nil
}

// Check the provided password using `crypt()`
func CheckPasswordHash(password string, hash string) (bool, error) {
	passHash, err := crypt(password, hash)
	if err != nil {
		return false, err
	}
	return passHash == hash, nil
}
