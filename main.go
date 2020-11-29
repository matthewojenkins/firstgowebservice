package main

import (
	"fmt"
	"log"
	"net/http"
	. "github.com/matthewojenkins/firstgowebservice/handlers"
)

func main() {
	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/program", handleProgramRequest)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Working")
}

func handleProgramRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("Handling program POST")
		HandleProgramPost(w, r)
	}
	if r.Method == http.MethodGet {
		fmt.Println("Handling program GET ")
		HandleProgramGet(w, r)
	}
}
