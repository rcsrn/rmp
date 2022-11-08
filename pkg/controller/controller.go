package controller

import (
	"github.com/rcsrn/rmp/pkg/database"
	"github.com/rcsrn/rmp/pkg/miner"
	"log"
	"fmt"
	"os/user"
	_"errors"
)

func Run() {
	obtainData()
}

func obtainData() {
	miner := miner.CreateNewMiner("/home/casarin/Escuela/Modelado/Proyectos/rmp/test/miner/TestRolas")

	err := miner.Traverse()	
	check(err)
	
	err = miner.MineTags()
	check(err)
	dbPath := getDBPath()
	rolas := miner.GetRolas()
	
	builder := database.CreateNewBuilder(rolas, dbPath)
	database, err := builder.BuildDataBase()
	check(err)
	fmt.Println(database)
}

func check(err error) {
	if err != nil {
	//it should show the error to user.
		log.Fatal(err)
	}
}

func getDBPath() string{
	user, err := user.Current()
	check(err)
	return user.HomeDir + "/.local/rmp"
}
