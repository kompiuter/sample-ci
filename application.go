package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Beanstalk from Codeship!! %v", time.Now())
	})
	log.Println("Listening on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
