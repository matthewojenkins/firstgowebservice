package database

import (
	"log"

	. "github.com/matthewojenkins/firstgowebservice/model"
)

func SaveProgram(program string) error {
	db := GetDB()
	_, err := db.Exec("INSERT INTO program (program_name) VALUES($1)", program)
	return err
}

func IsDatabaseUp() bool {
	db := GetDB()
	_, err := db.Exec("select 1")
	if err != nil {
		return false
	} else {
		return true
	}
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

	return Program{ProgramID: programID, ProgramName: programName}
}

func GetPrograms() []Program {
	db := GetDB()

	sqlStatement := "SELECT program_id, program_name FROM program order by program_name;"
	var programName string
	var programID int
	var program Program
	var programs []Program

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&programID, &programName)
		if err != nil {
			log.Fatal(err)
		}
		program = Program{ProgramID: programID, ProgramName: programName}
		programs = append(programs, program)
	}
	return programs
}
