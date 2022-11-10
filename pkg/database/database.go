package database

import (
	_"github.com/mattn/go-sqlite3"
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type DataBase struct {
	db *sql.DB
	dbPath string
	fileExists bool
}

func CreateNewDataBase(dbPath string) (*DataBase, error) {
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


//AddRola inserts a rola into the database using the idPerformer and idAlbum  //and returns the id associated to the rola added.
//If the rolas was already in the databse AddRola returns -1.
func (database *DataBase) AddRola(rola *Rola, idPerformer int64, idAlbum int64) (int64, error) {
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

	result, err := stmt.Exec(idPerformer, idAlbum, rola.GetPath(), rola.GetTitle(), rola.GetTrack(),
		rola.GetYear(), rola.GetGenre(), rola.GetTitle(), idPerformer, idAlbum, rola.GetGenre(), rola.GetPath())
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

func (database *DataBase) AddAlbum(rola *Rola) (int64, error) {
	idalbum, err := database.ExistsAlbum(filepath.Dir(rola.GetPath()), rola.GetAlbum())
	if err != nil {
		return -1, err
	}

	if idalbum > 0 {
		return idalbum, nil
	}
	
	stmtStr := `INSERT
                INTO albums (
                  path,
                  name,
                  year)
                SELECT ?, ?, ?
                WHERE NOT EXISTS
                (SELECT 1 FROM albums WHERE path = ? AND name = ?)`
	
	tx, stmt, err := database.PrepareStatement(stmtStr)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	
	id, err := stmt.Exec(filepath.Dir(rola.GetPath()), rola.GetAlbum(), rola.GetYear(), filepath.Dir(rola.GetPath()), rola.GetAlbum())
	if err != nil {
		return -1, err
	}
	tx.Commit()
	lastId, err := id.LastInsertId()
	if err != nil {
		return -1, err
	}
	return lastId, nil
}

func (database *DataBase) ExistsAlbum(albumPath, name string) (int64, error) {
	stmtStr := "SELECT id_album FROM albums WHERE albums.path = ? AND albums.name = ? LIMIT 1"
	tx, stmt, rows, err := database.PreparedQuery(stmtStr, albumPath, name)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	defer rows.Close()
	
	var id int64
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return -1, err
		}
	}
	err = rows.Err()
	if err != nil {
		return -1, err
	}
	tx.Commit()
	return id, nil
}

func (database *DataBase) AddPerformer(rola *Rola) (int64, error) {
	idp, err := database.ExistsPerformer(rola.GetPerformer())
	if err != nil {
		return -1, err
	}
	
	if idp > 0 {
		return idp, nil
	}

	stmtStr := `INSERT
                INTO performers (
                  id_type,
                  name)
                SELECT ?, ?
                WHERE NOT EXISTS
                (SELECT 1 FROM performers WHERE name = ?)`

	tx, stmt, err := database.PrepareStatement(stmtStr)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	id, err := stmt.Exec(2, strings.TrimSpace(rola.GetPerformer()), rola.GetPerformer())
	if err != nil {
		return -1, nil
	}
	tx.Commit()
	lastId, err := id.LastInsertId()
	if err != nil {
		return -1, nil
	}
	return lastId, nil
}

func (database *DataBase) ExistsPerformer(performerName string) (int64, error) {
	stmtStr := `SELECT
                  id_performer
                FROM performers
                WHERE performers.name = ?
                LIMIT 1`
	tx, stmt, rows, err := database.PreparedQuery(stmtStr, performerName)
	if err != nil {
		return -1, err
	}
	
	defer stmt.Close()
	defer rows.Close()
	
	var id int64
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return -1, err
		}
	}
	err = rows.Err()
	if err != nil {
		return -1, err
	}
	tx.Commit()
	return id, nil
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

func (database *DataBase) PreparedQuery(statement string, args ...interface{}) (*sql.Tx, *sql.Stmt, *sql.Rows, error) {
	tx, stmt, err := database.PrepareStatement(statement)
	if err != nil {
		return nil, nil, nil, err
	}
	
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, nil, nil, err
	}
	return tx, stmt, rows, nil
}
