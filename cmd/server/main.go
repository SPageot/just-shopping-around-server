package main

import (
	"just-shopping-around-server/internal/router"
)

func main(){
	r := router.CreateRouter()
	r.Logger.Fatal(r.Start(":8080"))
	
	
}



