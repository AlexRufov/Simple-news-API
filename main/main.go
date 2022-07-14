package main

import (
	"Simple-news-API/news"
	"Simple-news-API/router"
)

func main() {
	r := router.New()
	a := news.New()
	go a.Serve()
	err := r.Run()
	if err != nil {
		return
	}
}
