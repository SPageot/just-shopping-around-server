package main

import (
	"fmt"
	"just-shopping-around-server/internal/router"
	"net/http"
)

func main(){
	r := router.CreateRouter()
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", r ); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
	
}



