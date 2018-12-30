package main

import (
	"fmt"
	"log"
	"net/http"

	"flexible-sample/util"
)

func main() {
	http.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, util.RESPONSE)
}
