package graphics

import (
	"errors"
	"time"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

type Lockscreen struct {
	conn   *xgb.Conn
	screen *xproto.ScreenInfo
	window xproto.Window
}

func NewLockscreen() (*Lockscreen, error) {
	conn, err := xgb.NewConn()
	if err != nil {
		return nil, err
	}

	setup := xproto.Setup(conn)
	screen := setup.DefaultScreen(conn)

	return &Lockscreen{
		conn:   conn,
		screen: screen,
	}, nil
}

func (l *Lockscreen) Lock() error {
	window, err := xproto.NewWindowId(l.conn)
	if err != nil {
		return err
	}
	l.window = window

	mask := 0
	values := make([]uint32, 3)

	mask |= xproto.CwBackPixel
	values[0] = 0xFFFFFF

	mask |= xproto.CwOverrideRedirect
	values[1] = 1

	mask |= xproto.CwEventMask
	values[2] = (xproto.EventMaskKeyPress | xproto.EventMaskKeyRelease |
		xproto.EventMaskExposure | xproto.EventMaskVisibilityChange |
		xproto.EventMaskStructureNotify)

	xproto.CreateWindow(
		l.conn,
		l.screen.RootDepth,
		l.window,
		l.screen.Root,
		0, 0, 100, 100, 0,
		xproto.WindowClassInputOutput,
		l.screen.RootVisual,
		uint32(mask),
		values,
	)

	err = l.GrabPointerKeyboard()
	if err != nil {
		return err
	}

	err = xproto.MapWindowChecked(l.conn, l.window).Check()
	if err != nil {
		return err
	}
	return nil
}

func (l *Lockscreen) Unlock() error {
	// TODO
	xproto.UngrabPointer(l.conn, 0)
	xproto.UngrabKeyboard(l.conn, 0)

	return nil
}

func (l *Lockscreen) GrabPointerKeyboard() error {
	var pointerGrab *xproto.GrabPointerReply
	var keyboardGrab *xproto.GrabKeyboardReply

	for i := 0; i < 10; i++ {
		if pointerGrab == nil || pointerGrab.Status != xproto.GrabStatusSuccess {
			pointerGrab, _ = xproto.GrabPointer(
				l.conn,
				false,
				l.screen.Root,
				0,
				xproto.GrabModeAsync,
				xproto.GrabModeAsync,
				0,
				0,
				0,
			).Reply()
		}

		if keyboardGrab == nil || keyboardGrab.Status != xproto.GrabStatusSuccess {
			keyboardGrab, _ = xproto.GrabKeyboard(
				l.conn,
				true,
				l.screen.Root,
				0,
				xproto.GrabModeAsync,
				xproto.GrabModeAsync,
			).Reply()
		}

		if pointerGrab.Status == xproto.GrabStatusSuccess &&
			keyboardGrab.Status == xproto.GrabStatusSuccess {
			break
		}

		time.Sleep(100 * time.Millisecond)
	}

	if pointerGrab.Status != xproto.GrabStatusSuccess ||
		keyboardGrab.Status != xproto.GrabStatusSuccess {
		return errors.New("Unable to grab pointer and keyboard.")
	}

	return nil
}
