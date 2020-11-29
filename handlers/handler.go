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

func HandleProgramGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(strings.Join(r.URL.Query()["ID"], ""))
	program := GetProgram(id)
	fmt.Println("Program found : ", program)

	json.NewEncoder(w).Encode(program)
}
