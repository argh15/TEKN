package controller

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

var Welcome = func(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome")
}


var Upload =  func(w http.ResponseWriter, r *http.Request){

	// reading from request
	input,header,readErr := r.FormFile("video")
	if readErr != nil{
		fmt.Fprint(w,readErr)
		return;
	}
	defer input.Close()

	//open a file and write to it
	out,fileErr := os.Create("media/"+header.Filename)
	if fileErr != nil{
		fmt.Fprint(w,fileErr)
		return;
	}
	defer out.Close()
	_,writeErr:= io.Copy(out,input)
	if writeErr != nil{
		fmt.Fprint(w,writeErr)
		return;
	}

	fmt.Fprint(w,"Successful")

	// Creating database entry
	

}
