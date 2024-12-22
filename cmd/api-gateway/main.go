package main

import (
	"fmt"
	"stugi/api-gateway/internal/app"
	"stugi/api-gateway/internal/config"
	"stugi/api-gateway/internal/service"
)

func main() {
	_ = config.New()
	gateway := service.New()
	app := app.New(&gateway)
	app.Run()
	fmt.Println("api-gateway started")
	select {}
}
