package route

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heaptracetechnology/microservice-kafka/message"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Cron",
		"POST",
		"/consume",
		message.consume,
	},
	Route{
		"Cron",
		"POST",
		"/consume",
		message.produce,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
