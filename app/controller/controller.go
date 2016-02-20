package controller

import (
	"github.com/YuriSS/memory_game/app/routes"
	"github.com/YuriSS/memory_game/app/stage"
	"net/http"
)

type Middleware struct {
	api_prefix string
}

var scene = stage.Scene{}
var router = routes.Route{}

func (middleware *Middleware) Raises_Routes() {
	middleware.api_prefix = "/api/"

	http.HandleFunc("/", router.Make_Handler(router.Index_Handler))
	http.HandleFunc("/static/", router.Make_Handler(router.Static_Files))

	http.HandleFunc(middleware.api_prefix+"move/", router.Make_Handler(router.Move))
}

func (router *Middleware) Start_Game() {
	scene.New_Scene()
}
