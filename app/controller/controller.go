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

	http.HandleFunc("/", make_handler(index_handler))
	http.HandleFunc("/static/", make_handler(static_files))
}

func make_handler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fn(writer, request)
	}
}

func static_files(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, request.URL.Path[1:])
}

func index_handler(writer http.ResponseWriter, request *http.Request) {
	view.Render(writer, "index", view.Page{"Jogo da Memória"})
}
