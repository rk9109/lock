package graphics

import (
	"errors"
	"fmt"
)

func (l *Lockscreen) EventLoop() error {
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
