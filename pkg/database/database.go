package database

import (	
	//"github.com/mattn/go-sqlite3"
	"github.com/qustavo/dotsql"
	"database/sql"
	"errors"
	"os"
)

type DataBase struct {
	db *sql.DB
	dbPath string
}

func CreateNewDataBase(dbPath string) (*DataBase, error) {
	os.Mkdir(dbPath, 0700)

	dbFileExists := true
	
	if _, err := os.Stat(dbPath + "/rmp.db"); os.IsNotExist(err) {
		dbFileExists = false
	}
	
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, errors.New("Could not open the database." + ": " + err.Error())
	}

	if !dbFileExists {
		if err = createDBFile(dbPath, db); err != nil {
			return nil, err
		}
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

func createDBFile(dbPath string, db *sql.DB) error {
	if _, err := os.Stat(dbPath + "/rmp.sql"); os.IsNotExist(err) {
		os.Create(dbPath + "/rmp.sql")
	}
	
	dot, err := dotsql.LoadFromFile(dbPath + "/rmp.sql")
	if err != nil {
		return errors.New("It is not possible to load rmp.sql" + ": " + err.Error())
	}
	
	nameTags := make ([]string, 0)

	nameTags = append(nameTags, "create-" + "performers" + "-table")
	nameTags = append(nameTags, "create-" + "persons" + "-table")
	nameTags = append(nameTags, "create-" + "groups" + "-table")
	nameTags = append(nameTags, "create-" + "albums" + "-table")
	nameTags = append(nameTags, "create-" + "rolas" + "-table")
	nameTags = append(nameTags, "create-" + "in_group" + "-table")

	for _, query := range nameTags {
		_, err = dot.Exec(db, query)
		if err != nil {
			return err
		}
	}
	return nil
}


