package src

import (
	"encoding/json"
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
	Route{"Health", "POST", "/email", MyPostEndpoint},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _,route := range routes {
		router.Name(route.Name).Methods(route.Method).Path(route.Pattern).Handler(route.HandleFunc)
	}
	return router
}

func buildResponse(w http.ResponseWriter, httpStatus int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(response)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	tools.Log()
	fmt.Fprintf(w, "pong")
}


func MyPostEndpoint(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		fmt.Println(key, ":", value)
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var object interface{}
	err := decoder.Decode(&object)
	if err != nil {
		buildResponse(w, http.StatusServiceUnavailable, nil )
		return
	}
	fmt.Println(object)
	buildResponse(w, http.StatusCreated, object )
}