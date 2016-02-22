package routes

import (
	"net/http"
)

type Route struct{}

func (route *Route) Make_Handler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fn(writer, request)
	}
}
