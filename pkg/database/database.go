package database

import (
	_"github.com/mattn/go-sqlite3"
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"fmt"
)

type DataBase struct {
	db *sql.DB
	dbPath string
	fileExists bool
}

//CreateNewDataBase creates new a database handler.
//It verifies whether the .db file exists in "~/.local/rmp/". If the file
// does not exists then it is created opening a connection to it.
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


//Addrola inserts a rola into the database using the idPerformer and idAlbum
//and returns the id associated to the rola added.
//If the rolas was already in the databse AddRola returns -1.
func (database *DataBase) AddRola(rola *Rola, idPerformer int, idAlbum int) (int64, error) {
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


//AddAlbum insters the rola into the database using the rola associated to it.
//It returns the id of the album that has been added.
//If the album was already in the database AddAlbum returns -1.
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

//ExitsAlbum checks if an album already exists in the database.
//If the album is already in the database it returns its id
//and 0 if the album is not in the database.
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


//AddPerformer insters the rola into the database using the rola associated to it.
//It returns the id of the album that has been added.
//If the album was already in the database AddPerformer returns -1.
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


//ExistsPerformer checks if an album already exists in the database.
//If the album is already in the database it returns its id
//and 0 if the album is not in the database.
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

//Load pings the database to verify if the connnection is active.
func (database *DataBase) Load() error {
	err := database.db.Ping()
	if err != nil {
		return errors.New("Could not load the database: " + err.Error())
	}
	return nil
}


//PrepareStatement initializes a sqlite prepared statement from a string
//and returns the corresponding sql context and prepared statement.
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

// PreparedQuery executes a prepared query and returns the resulting rows,
// it handles the errors and returns the context and prepared statement
// for the user to close them.
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

//QueryPerformerById receives an id of a rola an returns the path associated to it.
//Prepares the SQL statement to query the path. It handles the errors and then returns
//the path.
func (database *DataBase) QueryPerformer(id int64) (int, string, error) {
	stmtStr := "SELECT " +
		" performers.id_type, " +
		" performers.name " +
		"FROM " +
		" performers " +
		"INNER JOIN types ON types.id_type = performers.id_type " +
		"WHERE " +
		" performers.id_performer = ?"

	tx, stmt, err := database.PrepareStatement(stmtStr)
	if err != nil {
		return -1, "", err
	}
	defer stmt.Close()
	
	rows, err := stmt.Query(id)
	if err != nil {
		return -1, "", err
	}
	defer rows.Close()

	var performerType int
	var name string
	for rows.Next() {
		err = rows.Scan(&performerType, &name)
		if err != nil {
			return -1, "", err
		}
	}
	err = rows.Err()
	if err != nil {
		return -1, "", err
	}
	tx.Commit()
	return performerType, name, nil
}


//QueryRola receives a idRola and prepares a statement to get the rola
//associated to the idRola. It is assumed that the rolas is in the database.
func (database *DataBase) QueryRola(idRola int64) (*Rola, error) {
	stmtStr :=  "SELECT " +
		" performers.name, " +
		" albums.name, " +
		" rolas.path, " +
		" rolas.title, " +
		" rolas.track, " +
		" rolas.year, " +
		" rolas.genre " +
		"FROM rolas " +
		"INNER JOIN performers ON performers.id_performer = rolas.id_performer " +
		"INNER JOIN albums ON albums.id_album = rolas.id_album " +
		"WHERE " +
		" rolas.id_rola = ?"
	
	tx, stmt, err := database.PrepareStatement(stmtStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	
	rows, err := stmt.Query(idRola)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var performer string
	var album string
	var path string
	var title string
	var track int
	var year int
	var genre string
	
	for rows.Next() {
		err = rows.Scan(&performer, &album, &path, &title, &track, &year, &genre)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return &Rola{idRola, performer, album, path, title, track, year, genre}, nil
}

//Query returns the rola's name associated to idRola.
func (database *DataBase) QuerySpecificColumnFromRolas(specificColumn string, condition string) (string, error) {
	stmtStr := fmt.Sprintf("SELECT %v FROM rolas WHERE %v", specificColumn, condition)
	_, stmt, err := database.PrepareStatement(stmtStr)
	if err != nil {
		return "", nil
	}
	defer stmt.Close()
	
	row, err := stmt.Query(specificColumn, condition)
	if err != nil {
		return "", err
	}

	var name string
	
	for row.Next() {
		err := row.Scan(&name)
		if err != nil {
			return "", err
		}
	}
	
	return name, nil
}

//QueryGeneralString takes a general string as a paramater an returns an array with 
//the id of all rolas that contain the string in its performer name, album name,
//title or genre.
func (database *DataBase) QueryGeneralString(general string) ([]int64, error) {
	result := make([]int64, 0)
	stmtStr := "SELECT " +
		" rolas.id_rola " +
		"FROM " +
		" rolas " +
		"INNER JOIN performers ON performers.id_performer = rolas.id_performer " +
		"INNER JOIN albums ON albums.id_album = rolas.id_album " +
		"WHERE " +
		" performers.name LIKE ? " +
		" OR albums.name LIKE ? " +
		" OR rolas.title LIKE ? " +
		" OR rolas.genre LIKE ?"

	general = "%" + strings.TrimSpace(general) + "%"
	tx, stmt, rows, err := database.PreparedQuery(stmtStr, general, general, general, general)

	if err != nil {
		return nil, err
	}
	
	defer stmt.Close()
	defer rows.Close()

	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		result = append(result, id)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return result, nil
}

