package main

import (
	. "github.com/matthewojenkins/firstgowebservice/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleHomePage)
	http.HandleFunc("/status", HandleStatusPage)
	http.HandleFunc("/program", HandleProgramRequest)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
