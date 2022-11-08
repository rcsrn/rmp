package test

import (
	"testing"
	"github.com/rcsrn/rmp/pkg/database"
	"github.com/rcsrn/rmp/pkg/miner"
	"fmt"
	"os"
)

var testBuilder *database.Builder
var testDataBase *database.DataBase

func initBuilder() {
	miner := miner.CreateNewMiner("/home/casarin/Escuela/Modelado/Proyectos/rmp/test/miner/TestRolas")
	miner.Traverse()
	miner.MineTags()
	rolas := miner.GetRolas()
	currentDirectory, _ := os.Getwd()
	testBuilder = database.CreateNewBuilder(rolas, currentDirectory)
}

func TestBuildDataBase(t *testing.T) {
	testBuilder = database.CreateNewBuilder(nil, "")
	dataBase, err := testBuilder.BuildDataBase()
	if err == nil || dataBase != nil {
		t.Errorf("this could not happen")
	}
	
	initBuilder()
	testDataBase, err = testBuilder.BuildDataBase()
	if err != nil || testDataBase == nil {
		t.Errorf("this could not happen: " + err.Error())
	}

	fmt.Println(testDataBase)

	query := "find-rolas-by-id"
	
	result, err := testDataBase.Query(query, 1)
	if err != nil {
		t.Errorf("this could not happen: " + err.Error())
	}
	fmt.Println(result)
}
