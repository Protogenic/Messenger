package main

import (
	"fmt"
	"messenger/internal/app"
	"messenger/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.Load()

	fmt.Println(cfg)

	Application := app.New(cfg.GRPC.Port)

	go Application.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	Application.Stop()
	fmt.Println("Application stopped")
}
