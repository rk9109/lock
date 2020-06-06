package main

import (
	"errors"
	"fmt"

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
		0, 0, 500, 500, 0,
		xproto.WindowClassInputOutput,
		l.screen.RootVisual,
		uint32(mask),
		values,
	)

	err = xproto.MapWindowChecked(l.conn, l.window).Check()
	if err != nil {
		return err
	}
	return nil
}

func (l *Lockscreen) Unlock() error {
	// TODO
    return nil
}

func grabPointerKeyboard() error {
	// TODO
	return nil
}

func ungrabPointerKeyboard() error {
	// TODO
	return nil
}

func (l *Lockscreen) xEventLoop() error {
	for {
		event, xerr := l.conn.WaitForEvent()
		if event == nil && xerr == nil {
			return errors.New("X event loop error")
		}
		if event != nil {
			fmt.Printf("Event: %s\n", event)
		}
		if xerr != nil {
			fmt.Printf("Error: %s\n", xerr)
		}
	}
}
