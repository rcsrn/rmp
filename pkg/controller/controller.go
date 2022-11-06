package controller

import (
	"github.com/rcsrn/rmp/pkg/database"
	"github.com/rcsrn/rmp/pkg/miner"
	"log"
	"fmt"
	"os/user"
	"errors"
)

func Run() {
	obtainData()
}

func obtainData() {
	miner := miner.CreateNewMiner("/home/casarin/Escuela/Modelado/Proyectos/rmp/test/miner/TestRolas")

	err := miner.Traverse()	
	if err != nil {
		fatal(err)
	}
	
	err = miner.MineTags()
	if err != nil {
		fatal(err)
	}
	dbPath := getDBPath()
	rolas := miner.GetRolas()
	
	builder := database.CreateNewBuilder(rolas, dbPath)
	database, err := builder.BuildDataBase()
	if err != nil {
		fatal(err)
	}
	fmt.Println(database)
}

func fatal(err error) {
	//it should show the error to user.
	log.Fatal(err)	
}

func getDBPath() string{
	user, err := user.Current()
	if err != nil {
		fatal(errors.New("Could not retrieve the current user."))
	}
	return user.HomeDir + "/.local/rmpDB/rmp.sql"
}
