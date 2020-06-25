package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var message = "Hello World"
		w.Write([]byte(message))
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
