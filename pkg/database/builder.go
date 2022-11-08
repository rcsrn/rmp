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

	setup := make([]string, 0)

	setup = append(setup, CREATE+"types-table")
	setup = append(setup, CREATE+"type0")
	setup = append(setup, CREATE+"type1")
	setup = append(setup, CREATE+"type2")
	setup = append(setup, CREATE+"performers"+TABLE)
	setup = append(setup, CREATE+"persons"+TABLE)
	setup = append(setup, CREATE+"groups"+TABLE)
	setup = append(setup, CREATE+"albums"+TABLE)
	setup = append(setup, CREATE+"rolas"+TABLE)
	setup = append(setup, CREATE+"in_group"+TABLE)

	for _, query := range setup {
		_, err := builder.executer.Exec(database.db, query)
		if err != nil {
			return err
		}
	}

	database.fileExists = true
	
	return nil
}

func (builder *Builder) poblateDataBase(database *DataBase) error {
	id := 0
	for _, rola := range(builder.rolas) {
		_, err := builder.executer.Exec(database.db, "insert-rola", id, rola.GetPerformer(), rola.GetAlbum(), rola.GetPath(), rola.GetTitle(), rola.GetTrack(), rola.GetYear(), rola.GetGenre())
		id++
		if err != nil {
			return err
		}
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
