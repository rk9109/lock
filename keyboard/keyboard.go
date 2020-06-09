package keyboard

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
// #include <X11/Xutil.h>
//
// static int
// LookupString(Display* display,
//              unsigned int state,
//              unsigned int keycode,
//              char* buffer,
//              int buffer_len,
//              long unsigned int* keysym)
// {
//     int num;
//     XKeyEvent event;
//
//     event.display = display;
//     event.state = state;
//     event.keycode = keycode;
//
//     num = XLookupString(&event, buffer, buffer_len, keysym, 0);
//
//     return num;
// }
import "C"
import (
	"errors"
	"unsafe"

	"github.com/BurntSushi/xgb/xproto"
)

var display *C.struct__XDisplay

// Go wrapper around `XOpenDisplay()`
// Open a connection to the X server using Xlib.
func InitX() error {
	display = C.XOpenDisplay(nil)
	if display == nil {
		return errors.New("Unable to open X connection.")
	}
	return nil
}

// Go wrapper around `XLookupString()`
// Get string representation and keysym associated to a keypress event.
func LookupString(event *xproto.KeyPressEvent) (string, xproto.Keysym) {
	var keysym C.ulong
	buf := make([]byte, 32)

	numC := C.LookupString(
		display,
		C.uint(event.State),
		C.uint(event.Detail),
		(*C.char)(unsafe.Pointer(&buf[0])),
		C.int(len(buf)),
		&keysym,
	)
	num := int(numC)

	return string(buf[:num]), xproto.Keysym(keysym)
}
