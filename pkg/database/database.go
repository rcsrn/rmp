package database

import (	
	//"github.com/mattn/go-sqlite3"
	"database/sql"
	"os/user"
	"errors"
	"os"
)

type DataBase struct {
	db *sql.DB
}

func CreateNewDataBase() (*DataBase, error) {
	user, err := user.Current()
	if err != nil {
		return nil, errors.New("Could  not retrieve the user" + ": " + err.Error())
	}
	
	dbPath := user.HomeDir + "./internal/rmpDB/"

	if _, err := os.Stat(dbPath + "/rmp.db"); os.IsNotExist(err) {
		//CREATE A THE DB FILE.
	}
	
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, errors.New("Could not open the database." + ": " + err.Error())
	}
	
	dbError := db.Ping()
	if dbError != nil {
		return nil, dbError
	}
	return &DataBase{db}, nil
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
