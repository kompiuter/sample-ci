package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", indexHandler)

	log.Println("Listening on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v", r.Form)
	num, _ := strconv.Atoi(r.FormValue("num"))
	fmt.Fprintf(w, "%d is %s", num, size(num))
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
