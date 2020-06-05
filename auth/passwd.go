// Based on https://github.com/linuxdeepin/go-lib

package auth

// #include <stdlib.h>
// #include <pwd.h>
import "C"
import (
	"errors"
	"unsafe"
)

type Password struct {
	Name     string // username
	Password string // user password
	Uid      uint32 // user ID
	Gid      uint32 // group ID
	Comment  string // user information
	Home     string // home
	Shell    string // shell program
}

func convertPassword(passwordC *C.struct_passwd) *Password {
	return &Password{
		Name:     C.GoString(passwordC.pw_name),
		Password: C.GoString(passwordC.pw_passwd),
		Uid:      uint32(passwordC.pw_uid),
		Gid:      uint32(passwordC.pw_gid),
		Comment:  C.GoString(passwordC.pw_gecos),
		Home:     C.GoString(passwordC.pw_dir),
		Shell:    C.GoString(passwordC.pw_shell),
	}
}

// Go wrapper around `getpwnam`
// Get password file entry corresponding to given username.
func GetPasswordName(name string) (*Password, error) {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	passwordC, err := C.getpwnam(nameC)

	// If matching passwd record cannot be found, `getpwnam` should return
	// NULL and leave `errno` unchanged.
	if passwordC == nil {
		if err == nil {
			return nil, errors.New("Password record not found.")
		} else {
			return nil, err
		}
	}
	return convertPassword(passwordC), nil
}

// Go wrapper around `getpwuid`
// Get password file entry corresponding to given UID.
func GetPasswordUid(uid uint32) (*Password, error) {
	uidC := C.__uid_t(uid)
	passwordC, err := C.getpwuid(uidC)

	// If matching passwd record cannot be found, `getpwuid` should return
	// NULL and leave `errno` unchanged.
	if passwordC == nil {
		if err == nil {
			return nil, errors.New("Password record not found.")
		} else {
			return nil, err
		}
	}
	return convertPassword(passwordC), nil
}
