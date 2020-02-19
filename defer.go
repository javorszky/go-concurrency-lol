package main

import (
	"fmt"
	"time"
)

func deferred(what string) {
	fmt.Println(fmt.Sprintf("[deferred-%s] Starting deferred function... sleeping for 3 seconds", what))
	time.Sleep(3 * time.Second)
	fmt.Println(fmt.Sprintf("[deferred-%s] Bye", what))
}

func main() {
	fmt.Println("[main] Setting up deferred function for main")
	defer deferred("main")
	go func() {
		fmt.Println("[gofunc] Setting up the deferred function for gofunc")
		defer deferred("gofunc")
	}()
	fmt.Println("[main] Just printing something before returning in 3 seconds...")
	time.Sleep(3 * time.Second)
	fmt.Println("[main] It is time")
	return
}
