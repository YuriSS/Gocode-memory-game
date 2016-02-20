package controller

import (
	"github.com/YuriSS/memory_game/app/view"
	"net/http"
)

type Router struct {
	api_prefix string
}

func (router *Router) Raises_Routes() {
	router.api_prefix = "/api/"

	http.HandleFunc("/", makeHandler(indexHandler))
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fn(writer, request)
	}
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	view.Render(writer, "index", view.Page{"Jogo da Memória"})
}
