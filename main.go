// Following along with this tutorial
// https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	http.HandleFunc("/", handler).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
