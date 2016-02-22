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
	fmt.Printf("Server running on port %d...\n", config.port)
	controller := controller.Middleware{}
	controller.Start_Game()
	controller.Raises_Routes()
	panic(http.ListenAndServe(config.port, nil))
}
