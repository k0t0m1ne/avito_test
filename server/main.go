package main

import (
	"avitoTest/initializer"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var db = initializer.DbConnectionInit()

func main() {
	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
		log.Println("Connection to database closed successfully")
	}()

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)
		<-sigChannel
		close(sigChannel)
		cancel()
	}()

	Serve(ctx)
}
