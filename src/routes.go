package src

import (
	"fmt"
	"github.com/gorilla/mux"
	"my_heroku_app/src/tools"
	"net/http"
)

type Route struct {
	Name 		string
	Method 		string
	Pattern 	string
	HandleFunc 	http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route{"Health", "GET", "/ping", Ping},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _,route := range routes {
		router.Name(route.Name).Methods(route.Method).Path(route.Pattern).Handler(route.HandleFunc)
	}
	return router
}

func Ping(w http.ResponseWriter, r *http.Request) {
	tools.Log()
	fmt.Fprintf(w, "pong")
}