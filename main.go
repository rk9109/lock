package main

import (
	"fmt"
	"os"

	"github.com/rk9109/xgolock/graphics"
)

func run() error {
	lockscreen, err := graphics.NewLockscreen()
	if err != nil {
		return err
	}
	err = lockscreen.Lock()
	if err != nil {
		return err
	}
	lockscreen.EventLoop()

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
