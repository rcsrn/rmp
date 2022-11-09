package database

import (
	_"github.com/mattn/go-sqlite3"
	"database/sql"
	"errors"
	"os"
	_"path/filepath"
	"fmt"
)

type DataBase struct {
	db *sql.DB
	dbPath string
	fileExists bool
}

func CreateNewDataBase(dbPath string) (*DataBase, error) {
	fmt.Println(dbPath)
	os.Mkdir(dbPath, 0700)
	
	fileExists := true
	if _, err := os.Stat(dbPath + "/rmp.db"); os.IsNotExist(err) {
	 	fileExists = false
	}
	
	db, err := sql.Open("sqlite3", dbPath + "/rmp.db")
	if err != nil {
		return nil, errors.New("Could not open the database." + ": " + err.Error())
	}

	return &DataBase{db, dbPath, fileExists}, nil
}

func (database *DataBase) AddRola(rola *Rola, idperformer int64, idalbum int64) (int64, error) {
	stmtStr := `INSERT
                INTO rolas (
                  id_performer,
                  id_album,
                  path,
                  title,
                  track,
                  year,
                  genre)
                SELECT ?, ?, ?, ?, ?, ?, ?
                WHERE NOT EXISTS
                (SELECT 1 FROM rolas WHERE (title = ?
                  AND id_performer = ?
                  AND id_album = ?
                  AND genre = ?)
				  OR path = ?)`
	
	tx, stmt, err := database.PrepareStatement(stmtStr)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(idperformer, idalbum, rola.GetPath(), rola.GetTitle(), rola.GetTrack(),
		rola.GetYear(), rola.GetGenre(), rola.GetTitle(), idperformer, idalbum, rola.GetGenre(), rola.GetPath())
	if err != nil {
		return -1, errors.New("could not execute insert:" + err.Error())
	}
	rowsAdded, err := result.RowsAffected()
	if err != nil {
		return -1, errors.New("could not retrieve number of affected rows:" + err.Error())
	}
	tx.Commit()
	if rowsAdded > 0 {
		id, err := result.LastInsertId()
		if err != nil {
			errors.New("Could not retrieve last inserted id:" + err.Error())
		}
		return id, nil
	}
	return -1, nil
}

func (database *DataBase) AddAlbum() {
	
}

func (database *DataBase) AddPerformer() {
	
}

func (database *DataBase) AddGroup() {
	
}

func (database *DataBase) AddPerson() {
	
}

func (database *DataBase) Load() error {
	err := database.db.Ping()
	if err != nil {
		return errors.New("Could not load the database: " + err.Error())
	}
	return nil
}

func (database *DataBase) PrepareStatement(statement string) (*sql.Tx, *sql.Stmt, error) {
	tx, err := database.db.Begin()
	if err != nil {
		return nil, nil, errors.New("could not begin transaction: " + err.Error())
	}
	stmt, err := tx.Prepare(statement)
	if err != nil {
		return nil, nil, errors.New("could not prepare statement: " + err.Error())
	}
	return tx, stmt, nil
}

func (database *DataBase) Query(query string, args ...any) (*sql.Rows, error) {
	return database.db.Query(query, args)
}
