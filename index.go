package main

import (
	"fmt"
	"net/http"
	"code/router"
)

func main(){

	//attach auth middleware to this router
	router:= routerSetup.CreateRouter() 
	port:= "9090"

	err:= http.ListenAndServe(":"+port,router)
	if err != nil {
		fmt.Print(err)
	} else{
		fmt.Print("Listening on"+port)
	}

	

}