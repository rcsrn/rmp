package test

import (
	"testing"
	"github.com/rcsrn/rmp/pkg/database"
	"github.com/rcsrn/rmp/pkg/miner"
	_"fmt"
	"os"
	_"github.com/lib/pq"
)

var testBuilder *database.Builder
var testDataBase *database.DataBase

func initBuilder() {
	miner := miner.CreateNewMiner("/home/rodrigo/Escuela/Modelado/Proyectos/rmp/test/miner/TestRolas")
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
		t.Errorf("Could not build the database: " + err.Error())
	}
}

func TestAddRola(t *testing.T) {
	rola := database.CreateNewRola()
	rola.SetTitle("AAAA")
	id_rola, err := testDataBase.AddRola(rola, 1, 1)
	if err != nil {
		t.Errorf("The song has not been added correctly." + err.Error())
	}
	if id_rola == -1 {
		t.Errorf("The song was not in the database")
	}
	
}
