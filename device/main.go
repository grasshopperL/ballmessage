package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	log.Println("<INFO> [ main ] = {start}")
	chanOS := make(chan os.Signal, 1)
	signal.Notify(chanOS, os.Interrupt, syscall.SIGTERM)
	e := <-chanOS
	log.Println("<INFO> [main] = {exit}", e)
}
