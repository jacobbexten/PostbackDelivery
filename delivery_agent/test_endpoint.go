package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "TEST")
	})

	fmt.Printf("Starting server at port 8100\n")
	if err := http.ListenAndServe(":8100", nil); err != nil {
		log.Fatal(err)
	}
}
