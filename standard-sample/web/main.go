package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"

	"standard-sample/util"
)

func main() {
	http.HandleFunc("/", rootHandler)

	appengine.Main()
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, util.RESPONSE)
}
