package router

import (
	"just-shopping-around-server/internal/handler"

	"github.com/gorilla/mux"
)

 

func CreateRouter() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/api/news", handler.GetNews).Methods("GET")
	return r
}