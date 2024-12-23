package main

import (
	"fmt"
	"stugi/api-gateway/internal/app"
	"stugi/api-gateway/internal/config"
	"stugi/api-gateway/internal/service/comments"
	"stugi/api-gateway/internal/service/news"
)

func main() {
	_ = config.New()
	srvComments := comments.New("http://localhost:8082")
	srvNews := news.New("http://localhost:8081", srvComments)
	app := app.New(srvNews, srvComments)
	app.Run()
	fmt.Println("api-gateway started")
	select {}
}
