package controller

import (
	"github.com/rcsrn/rmp/pkg/database"
	"github.com/rcsrn/rmp/pkg/miner"
	"log"
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
	
	rolas := miner.GetRolas()
	builder := database.CreateNewBuilder()
	builder.SetRolas(rolas)
	database := builder.BuildDataBase()
}

func fatal(err error) {
		//it should show the error to user.
	log.Fatal(err)	
}
