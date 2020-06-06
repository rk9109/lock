package main

import (
	"fmt"
	"os"
)

func run() error {
    lockscreen, err := NewLockscreen()
    if err != nil {
        return err
    }
    err = lockscreen.Lock()
    if err != nil {
        return err
    }
    lockscreen.xEventLoop()

    return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
