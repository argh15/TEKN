package routerSetup

import (	
	"github.com/gorilla/mux"
	"code/controller"
)

func CreateRouter() *mux.Router {

	// use authentication middleware with this router 
	router:= mux.NewRouter()

	// Welcome route
	router.HandleFunc("/",controller.Welcome)

	//upload route
	router.HandleFunc("/upload",controller.Upload).Methods("POST")

	return router
}