package database

import (
	"log"
	. "oakdene.net/firstgoservice/model"
)

func SaveProgram(program string) error {
	db := GetDB()
	_, err := db.Exec("INSERT INTO program (program_name) VALUES($1)", program)
	return err
}

func GetProgram(id int) Program {
	db := GetDB()

	sqlStatement := "SELECT program_id, program_name FROM program WHERE program_id=$1;"
	var programName string
	var programID int

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&programID, &programName)
	if err != nil {
		log.Fatal(err)
	}

	return Program{programID, programName}
}
