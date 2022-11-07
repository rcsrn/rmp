package database

import (
	"github.com/qustavo/dotsql"
	"path/filepath"
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

	parent := filepath.Dir(database.dbPath)
	
	dot, err := dotsql.LoadFromFile(parent)
	if err != nil {
		return nil, err
	}

	nameTags := make([]string, 0)
	
	nameTags = append(nameTags, "create-" + "types-table")
	nameTags = append(nameTags, "create-" + "type0")
	nameTags = append(nameTags, "create-" + "type1")
	nameTags = append(nameTags, "create-" + "type2")
	nameTags = append(nameTags, "create-" + "performers" + "-table")
	nameTags = append(nameTags, "create-" + "persons" + "-table")
	nameTags = append(nameTags, "create-" + "groups" + "-table")
	nameTags = append(nameTags, "create-" + "albums" + "-table")
	nameTags = append(nameTags, "create-" + "rolas" + "-table")
	nameTags = append(nameTags, "create-" + "in_group" + "-table")

	for _, query := range nameTags {
		_, err = dot.Exec(database.db, query)
		if err != nil {
			return nil, err
		}
	}

	return database, nil
}

