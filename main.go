package main

import (
	"log"
	"net/http"
	"fmt"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API HOME PAGE")
	fmt.Println("Accessing API HOME PAGE")
}

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":3030", nil))
}