package database

import (	
	_"github.com/mattn/go-sqlite3"
	_"github.com/qustavo/dotsql"
	"database/sql"
	"errors"
	"os"
	_"path/filepath"
	"fmt"
)

type DataBase struct {
	db *sql.DB
	dbPath string
}

func CreateNewDataBase(dbPath string) (*DataBase, error) {
	if _, err := os.Stat(dbPath + "/rmp.db"); os.IsNotExist(err) {
	 	createDBFile(dbPath)
	 }
	
	
	db, err := sql.Open("sqlite3", dbPath + "/rmp.db")
	if err != nil {
		fmt.Println("ACA1")
		return nil, errors.New("Could not open the database." + ": " + err.Error())
	}

	//it verifies if database connection works.
	dbError := db.Ping()
	if dbError != nil {
		return nil, dbError
	}
	return &DataBase{db, dbPath}, nil
}

func (database *DataBase) AddRola() {
	
}

func (database *DataBase) AddAlbum() {
	
}

func (database *DataBase) AddPerformer() {
	
}

func (database *DataBase) AddGroup() {
	
}

func (database *DataBase) AddPerson() {
	
}

func createDBFile(dbPath string) error {
return nil
}


