package test

import (
	"testing"
	"github.com/rcsrn/rmp/pkg/database"
	"github.com/rcsrn/rmp/pkg/miner"
	"fmt"
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
	rola.SetTitle("Rola1")
	idRola, err := testDataBase.AddRola(rola, 1, 1)
	if err != nil {
		t.Errorf("The song has not been added correctly." + err.Error())
	}
	if idRola == -1 {
		t.Errorf("The song was not in the database")
	}
	
	rolaObtained, err := testDataBase.QueryRola(idRola)

	if err != nil {
		t.Error(err)
	}

	if rolaObtained == nil {
		t.Error("Could not get the song recently inserted")
	}
	
	if rolaObtained.GetTitle() != "Rola1" {
		t.Error("The rola has not been added correctly")
	}
}

func TestAddPerformer(t *testing.T) {
	rola := database.CreateNewRola()
	rola.SetPerformer("PerformerX")

	idPerformer, err := testDataBase.AddPerformer(rola)
	if err != nil {
		t.Errorf("The performer has not been added correctly" + err.Error())
	}
	
	if idPerformer == -1 {
		t.Errorf("The performer was not in the database")
	}

	_, performerObtained, err := testDataBase.QueryPerformer(idPerformer)
	if err != nil {
		t.Errorf(err.Error())
	}
	
	if performerObtained != "PerformerX" {
		t.Errorf("The performer obtained is not correctly.")
	}
	
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAddAlbum(t *testing.T) {
	rola := database.CreateNewRola()
	rola.SetAlbum("album1")

	idAlbum, err := testDataBase.AddAlbum(rola) 
	if err != nil {
		t.Errorf("The album has not been added correctly" + err.Error())
	}
	
	if idAlbum == -1 {
		t.Errorf("The album was not in the database")
	}
}

func TestExistsPerformer(t *testing.T) {
	rola := database.CreateNewRola()
	performerName := "PerformerY"
	rola.SetPerformer(performerName)
	testDataBase.AddPerformer(rola)

	idPerformer, err := testDataBase.ExistsPerformer(performerName)

	if err != nil {
		t.Errorf("Could not retrieve the performer" + err.Error())
	}

	if idPerformer == 0 {
		t.Errorf("The performer is already in the database")
	}
}

func TestExistAlbum (t *testing.T) {
	rola := database.CreateNewRola()
	albumName := "AlbumX"
	rola.SetAlbum(albumName)
	rola.SetPath("")
	testDataBase.AddAlbum(rola)

	idAlbum, err := testDataBase.ExistsAlbum("", albumName)

	if err != nil {
		t.Errorf("Could not retrieve the album" + err.Error())
	}

	fmt.Println(idAlbum)
	
	if idAlbum == 0 {
		t.Errorf("The album is already in the database")
	}
}
