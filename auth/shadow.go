// Based on https://github.com/linuxdeepin/go-lib

package auth

// #include <stdlib.h>
// #include <shadow.h>
import "C"
import (
	"errors"
	"unsafe"
)

type ShadowPassword struct {
	Name       string // username
	Password   string // password
	LastChange int64  // last password change
	Min        int64  // minimum number of days between password changes
	Max        int64  // maximum number of days between password changes
	Warn       int64  // number of days before user warned about password expiration
	Inactive   int64  // number of days before user account disabled
	Expire     int64  // account expiration date
	Flag       uint64 // reserved for future use
}

func convertShadowPassword(spasswordC *C.struct_spwd) *ShadowPassword {
	return &ShadowPassword{
		Name:       C.GoString(spasswordC.sp_namp),
		Password:   C.GoString(spasswordC.sp_pwdp),
		LastChange: int64(spasswordC.sp_lstchg),
		Min:        int64(spasswordC.sp_min),
		Max:        int64(spasswordC.sp_max),
		Warn:       int64(spasswordC.sp_warn),
		Inactive:   int64(spasswordC.sp_inact),
		Expire:     int64(spasswordC.sp_expire),
		Flag:       uint64(spasswordC.sp_flag),
	}
}

// Go wrapper around `getspnam()`
// Get shadow password file entry corresponding to given username.
func GetShadowPasswordName(name string) (*ShadowPassword, error) {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	spasswordC, err := C.getspnam(nameC)

	// If matching shadow password record cannot be found, `getspnam()` should
	// return NULL and leave `errno` unchanged.
	if spasswordC == nil {
		if err == nil {
			return nil, errors.New("Shadow password record not found.")
		} else {
			return nil, err
		}
	}
	return convertShadowPassword(spasswordC), nil
}
