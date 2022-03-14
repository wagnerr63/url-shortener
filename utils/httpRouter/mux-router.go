package httpRouter

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) POST(uri string, f http.HandlerFunc) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) GET(uri string, f http.HandlerFunc) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) PUT(uri string, f http.HandlerFunc) {
	muxDispatcher.HandleFunc(uri, f).Methods("PUT")
}

func (*muxRouter) PATCH(uri string, f http.HandlerFunc) {
	muxDispatcher.HandleFunc(uri, f).Methods("PATCH")
}

func (*muxRouter) DELETE(uri string, f http.HandlerFunc) {
	muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
}

func (*muxRouter) SERVE(port string) {
	fmt.Println("Server started at http://" + os.Getenv("HOST") + ":" + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(port, muxDispatcher))
}
