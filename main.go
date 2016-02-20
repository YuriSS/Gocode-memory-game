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
	controller := controller.Router{}
	controller.Raises_Routes()
	http.ListenAndServe(config.port, nil)
}
