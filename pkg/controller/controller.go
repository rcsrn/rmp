package controller

import (
	"github.com/rcsrn/rmp/pkg/database"
	"github.com/rcsrn/rmp/pkg/miner"
	"github.com/rcsrn/rmp/pkg/view"
	"log"
	"fmt"
	"os/user"
	"os"
	_"errors"
)

func Run() {
	filePath := startView()
	obtainData(filePath)
}

func obtainData(filePath string) {
	miner := miner.CreateNewMiner(filePath)

	err := miner.Traverse()	
	check(err)
	
	err = miner.MineTags()
	check(err)
	dbPath := getDBPath()
	rolas := miner.GetRolas()
	
	builder := database.CreateNewBuilder(rolas, dbPath)
	database, err := builder.BuildDataBase()
	check(err)

	fmt.Println(database.QueryGeneralString(""))
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

func startView() string {
	handler := view.CreateNewWindowHandler()
	handler.ShowLoadWindow()
	handler.RunApp()
	
	return handler.GetFilePath()
}


func getCurrentPath() string {
	current, _ := os.Getwd()
	return current
}

func getMusicPath() string {
	return ""
}
