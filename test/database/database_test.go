package test

import (
	"testing"
	"github.com/rcsrn/rmp/pkg/database"
)

var testDataBase *database.DataBase

func TestCreateNewDataBase(t *testing.T) {
	testDataBase, err := database.CreateNewDataBase()
}
