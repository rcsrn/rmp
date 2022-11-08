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
}

func CreateNewBuilder(rolas []*Rola, dbPath string) *Builder {
	return &Builder{rolas, dbPath}
}

func (builder *Builder) BuildDataBase() (*DataBase, error) {
	database, err := CreateNewDataBase(builder.dbPath)
	if err != nil {
		return nil, err
	}

	if !database.fileExists {
		CreateDBFile(database)
	}	

	err = database.Load()
	if err != nil {
		return nil, err
	}
	
	return database, nil
}

func CreateDBFile(database *DataBase) error{
	current, err := os.Getwd()
	if err != nil {
		return err
	}

	sqlPath := current + "/pkg/database/rmp.sql"
	
	
	dot, err := dotsql.LoadFromFile(sqlPath)
	if err != nil {
		return err
	}

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
	setup = append(setup, CREATE+"rmp"+TABLE)
	setup = append(setup, CREATE+"in_group"+TABLE)

	for _, query := range setup {
		_, err = dot.Exec(database.db, query)
		if err != nil {
			return err
		}
	}

	database.fileExists = true
	
	return nil
}
