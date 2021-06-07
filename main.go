package main

import (
	"fmt"
	"my_heroku_app/src"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	router := src.NewRouter()
	port := getPort()

	go http.ListenAndServe(port, router)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func getPort() string {
	port := os.Getenv("PORT")
	fmt.Println("port : ", port)
	if port == "" {
		return ":8080"
	}
	return ":" + port

}