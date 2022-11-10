package database

import (
	"github.com/qustavo/dotsql"
	_"path/filepath"
	"os"
	_"fmt"
)

type Builder struct {
	rolas []*Rola
	dbPath string
	executer *dotsql.DotSql
}

func CreateNewBuilder(rolas []*Rola, dbPath string) *Builder {
	return &Builder{rolas, dbPath, nil}
}

func (builder *Builder) BuildDataBase() (*DataBase, error) {
	database, err := CreateNewDataBase(builder.dbPath)
	if err != nil {
		return nil, err
	}
	
	executer, err := getExecuter()
	if err != nil {
		return nil, err
	}
	builder.executer = executer
	
	if !database.fileExists {
		builder.buildDBFile(database)
		err = builder.poblateDataBase(database)
		if err != nil {
			return nil, err
		}
	}	

	err = database.Load()
	if err != nil {
		return nil, err
	}
	
	return database, nil
}

func (builder *Builder) buildDBFile(database *DataBase) error{
	CREATE := "create-"
	TABLE := "-table"
	nameTags := make([]string, 0)

	nameTags = append(nameTags, CREATE+"types-table")
	nameTags = append(nameTags, CREATE+"type0")
	nameTags = append(nameTags, CREATE+"type1")
	nameTags = append(nameTags, CREATE+"type2")
	nameTags = append(nameTags, CREATE+"performers"+TABLE)
	nameTags = append(nameTags, CREATE+"persons"+TABLE)
	nameTags = append(nameTags, CREATE+"groups"+TABLE)
	nameTags = append(nameTags, CREATE+"albums"+TABLE)
	nameTags = append(nameTags, CREATE+"rolas"+TABLE)
	nameTags = append(nameTags, CREATE+"in_group"+TABLE)

	for _, query := range nameTags {
		_, err := builder.executer.Exec(database.db, query)
		if err != nil {
			return err
		}
	}

	database.fileExists = true
	
	return nil
}

func (builder *Builder) poblateDataBase(database *DataBase) error {
	id := 1
	for _, rola := range(builder.rolas) {
		err := builder.InsertRola(database, id, rola)
		if err != nil {
			return err
		}
		
		err = builder.InsertPerformer(database, id, rola)
		if err != nil {
			return err
		}

		err = builder.InsertAlbum(database, id, rola)
		id++
	}
	return nil
}

func isUnknown(tag string) bool {
	return tag == "<Unknown>"
}


func (builder *Builder) InsertRola(database *DataBase, id int, rola *Rola) error {
	_, err := builder.executer.Exec(database.db, "insert-rola", id, id, id, rola.GetPath(), rola.GetTitle(), rola.GetTrack(), rola.GetYear(), rola.GetGenre())
	if err != nil {
		return err
	}
	return nil
}

func (builder *Builder) InsertPerformer(database *DataBase, id int, rola *Rola) error {
	if isUnknown(rola.GetPerformer()) {
		_, err := builder.executer.Exec(database.db, "insert-performer", id, 2, rola.GetPerformer())
		if err != nil {
			return err
		}	
	} else {
		_, err := builder.executer.Exec(database.db, "insert-performer", id, 0, rola.GetPerformer()) 
		if err != nil {
			return err
		}
	}
	return nil
}

func (builder *Builder) InsertAlbum(database *DataBase, id int, rola *Rola) error {
	_, err := builder.executer.Exec(database.db, "insert-album", id, rola.GetPath(), rola.GetAlbum(), rola.GetYear())
	if err != nil {
		return err
	}
	return nil
}


func getExecuter() (*dotsql.DotSql, error) {
	current, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	sqlPath := current + "/pkg/database/rmp.sql"
	executer, err := dotsql.LoadFromFile(sqlPath)
	if err != nil {
		return nil, err
	}
	return executer, nil
}
