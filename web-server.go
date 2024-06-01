package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
http.NotFound(w, r)
return
	}
	fmt.Fprintf(w, "hello world!")
}