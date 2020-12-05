package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	. "github.com/matthewojenkins/firstgowebservice/database"
	. "github.com/matthewojenkins/firstgowebservice/model"
)

func HandleProgramPost(w http.ResponseWriter, r *http.Request) {
	var program Program

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&program)
	if err != nil {
		panic(err)
	}

	err = SaveProgram(program.ProgramName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	w.WriteHeader(201)
}

func HandleHomePage(w http.ResponseWriter, r *http.Request) {

}

func HandleProgramRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("Handling program POST")
		HandleProgramPost(w, r)
	}
	if r.Method == http.MethodGet {
		fmt.Println("Handling program GET ")
		HandleProgramGet(w, r)
	}
}

func HandleStatusPage(w http.ResponseWriter, r *http.Request) {
	var dbCheck string
	if IsDatabaseUp() {
		dbCheck = "Database is alive"
	} else {
		dbCheck = "Database is down"
	}
	err := json.NewEncoder(w).Encode(dbCheck)
	if err != nil {
		fmt.Println(err)
	}
}

func HandleProgramGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(strings.Join(r.URL.Query()["ID"], ""))

	if id > 0 {
		program := GetProgram(id)
		fmt.Println("Program found : ", program)

		err := json.NewEncoder(w).Encode(program)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		programs := GetPrograms()
		err := json.NewEncoder(w).Encode(programs)
		if err != nil {
			fmt.Println(err)
		}
	}
}
