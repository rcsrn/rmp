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

type MainApp struct {
	handler *view.WindowHandler
	database *database.DataBase
	filePath string
}

func createMainApp() *MainApp {
	return &MainApp{
		handler: view.CreateNewWindowHandler(),
		database: nil,
		filePath: "",
	}
}

func Run() {
	main := createMainApp()
	main.startView()
	main.obtainData()
}

func (main *MainApp) obtainData() {
	miner := miner.CreateNewMiner(main.filePath)

	err := miner.Traverse()	
	main.check(err)
	
	err = miner.MineTags()
	main.check(err)
	
	dbPath := main.getDBPath()
	rolas := miner.GetRolas()
	
	builder := database.CreateNewBuilder(rolas, dbPath)
	database, err := builder.BuildDataBase()
	main.check(err)

	main.database = database

}

func (main *MainApp) startView() {
	main.handler.ShowLoadWindow()
	
	main.handler.RunApp()
	
	main.obtainFilePath()
	fmt.Print(main.filePath)
	
}

func (main *MainApp) obtainFilePath() {
	main.filePath = main.handler.GetFilePath()
}

func (main *MainApp) getDBPath() string{
	user, err := user.Current()
	main.check(err)
	return user.HomeDir + "/.local/rmp"
}

func getCurrentPath() string {
	current, _ := os.Getwd()
	return current
}

func getMusicPath() string {
	return ""
}

func (main *MainApp) check(err error) {
	if err != nil {
		main.handler.ShowError(err.Error())
		log.Print(err)
		os.Exit(1)
	}
}
