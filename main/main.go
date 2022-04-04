package main

import (
	"MyNewsAPI/news"
	"MyNewsAPI/router"
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
