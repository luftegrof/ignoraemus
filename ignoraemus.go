package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	for signal := range c {
		switch signal {
			case syscall.SIGHUP:
				fmt.Printf("\nNo, you %s first!\n", signal)
		}
	}
}
