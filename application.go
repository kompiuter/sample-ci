package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello AWS Beanstalk! %v", time.Now())
	})
	http.ListenAndServe(":5000", nil)
}
