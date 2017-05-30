package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		num, _ := strconv.Atoi(r.FormValue("num"))
		fmt.Fprintf(w, "Hello, %d is %s", num, size(num))
	})

	log.Println("Listening on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func size(a int) string {
	switch {
	case a < 0:
		return "negative"
	case a == 0:
		return "zero"
	case a < 10:
		return "small"
	case a < 100:
		return "big"
	case a < 1000:
		return "huge"
	}
	return "enormous"
}
