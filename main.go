package main

import (
	"fmt"
	"github.com/YuriSS/memory_game/app/controller"
	"net/http"
)

var config = struct {
	port string
}{
	":3333",
}

func main() {
	fmt.Println("Server running on port %d...", config.port)
	controller := controller.Middleware{}
	controller.Raises_Routes()
	controller.Start_Game()
	panic(http.ListenAndServe(config.port, nil))
}
