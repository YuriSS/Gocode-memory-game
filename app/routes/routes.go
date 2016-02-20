package routes

import (
	"encoding/json"
	"fmt"
	"github.com/YuriSS/memory_game/app/view"
	"io/ioutil"
	"log"
	"net/http"
)

type Route struct{}

func (route *Route) Make_Handler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fn(writer, request)
	}
}

func (route *Route) Static_Files(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, request.URL.Path[1:])
}

func (route *Route) Index_Handler(writer http.ResponseWriter, request *http.Request) {
	view.Render(writer, "index", view.Page{"Jogo da Memória"})
}

type test_struct struct {
	Index string
}

func (route *Route) Move(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}
