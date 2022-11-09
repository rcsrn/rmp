package test

import (
	"testing"
	"github.com/rcsrn/rmp/pkg/database"
	"github.com/rcsrn/rmp/pkg/miner"
	"fmt"
	"os"
	"github.com/lib/pq"
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
	
	query := `SELECT * FROM rolas WHERE id_rolas = ANY($1)`
	
	rows, err := testDataBase.Query(query, pq.Array([]int{1}))
	if err != nil {
		t.Errorf("Could not get query: " + err.Error())
	}
	fmt.Println(rows )
}
