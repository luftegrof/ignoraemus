/* Handle Interrupts */

package main


import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for signal := range c {
			switch signal {
				case syscall.SIGHUP:
					fmt.Printf("\nNo, you %s first!\n", signal)
				case syscall.SIGINT:
					fmt.Printf("\nDo NOT %s me!\n", signal)
				case syscall.SIGTERM:
					fmt.Printf("\nI'll be back.\n")
			}
		}
	}()
	for true {
		fmt.Println("Just doing my thing...")
		time.Sleep(5 * time.Second)
	}
}
