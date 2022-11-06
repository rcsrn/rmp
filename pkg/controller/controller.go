package controller

import (
	"github.com/rcsrn/rmp/pkg/database"
	"github.com/rcsrn/rmp/pkg/miner"
	"log"
	"fmt"
	"path/filepath"
	"os/user"
	"errors"
)

func Run() {
	obtainData()
}

func obtainData() {
	miner := miner.CreateNewMiner("/home/casarin/m/")
	
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
		errors.New("Could not retrieve the current user.") 
	}
	return filepath.Dir(filepath.Dir(user.HomeDir)) + "/internal/db"
}
