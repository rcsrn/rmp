package database

import (
	_"github.com/qustavo/dotsql"
	_"path/filepath"
	_"os"
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
		database.CreateDBFile()
	}	

	err = database.Load()
	if err != nil {
		return nil, err
	}
	
	return database, nil
}
