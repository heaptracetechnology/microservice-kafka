package route

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	kafka "github.com/heaptracetechnology/microservice-kafka/kafka"
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
		"consume",
		"POST",
		"/consume",
		kafka.Consume,
	},
	Route{
		"produce",
		"POST",
		"/consume",
		kafka.ProduceStream, 
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
