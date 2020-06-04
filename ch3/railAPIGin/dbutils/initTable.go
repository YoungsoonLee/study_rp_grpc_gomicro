package dbutils

import (
	"database/sql"
	"log"
)

func Initialize(dbDriver *sql.DB) {
	statement, driveError := dbDriver.Prepare(train)
	if driveError != nil {
		log.Println(driveError)
	}

	_, statementError := statement.Exec()
	if statementError != nil {
		log.Println("Table already exists!")
	}

	statement, driveError = dbDriver.Prepare(station)
	if driveError != nil {
		log.Println(driveError)
	}

	_, statementError = statement.Exec()
	if statementError != nil {
		log.Println("Table already exists!")
	}

	statement, driveError = dbDriver.Prepare(schedule)
	if driveError != nil {
		log.Println(driveError)
	}

	_, statementError = statement.Exec()
	if statementError != nil {
		log.Println("Table already exists!")
	}

	log.Println("All tables created/initialized successfully")

}
