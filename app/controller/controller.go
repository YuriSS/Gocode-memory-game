package controller

import (
	"encoding/json"
	"github.com/YuriSS/memory_game/app/routes"
	"github.com/YuriSS/memory_game/app/stage"
	"github.com/YuriSS/memory_game/app/view"
	"log"
	"net/http"
)

type Middleware struct {
	api_prefix string
}

var scene = stage.Scene{}
var router = routes.Route{}

func (middleware *Middleware) Raises_Routes() {
	middleware.api_prefix = "/api/"

	http.HandleFunc("/", router.Make_Handler(index_handler))
	http.HandleFunc("/static/", router.Make_Handler(static_files))

	http.HandleFunc(middleware.api_prefix+"move/", router.Make_Handler(move))
}

func (router *Middleware) Start_Game() {
	scene.New_Scene()
}

func static_files(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, request.URL.Path[1:])
}

func index_handler(writer http.ResponseWriter, request *http.Request) {
	view.Render(writer, "index", view.Page{"Jogo da Memória"})
}

type Move_Response struct {
	Image string `json:"image"`
}

func move(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&scene.New_Play)
	handling_error(err, writer, "Decoder json ok")
	image := Move_Response{string(scene.Playing())}
	result, err := json.Marshal(image)
	handling_error(err, writer, image.Image)
	log.Println(string(result[:]))
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(result)
}

func handling_error(err error, w http.ResponseWriter, success_message string) bool {
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
		return false
	}
	log.Println(success_message)
	return true
}
