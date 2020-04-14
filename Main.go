package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("start webservice")
	http.HandleFunc("/", webserviceHandler)
	http.ListenAndServe(":5000", nil)
}

func webserviceHandler(w http.ResponseWriter, r *http.Request) {
	webside, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(webside))
}
