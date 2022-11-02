package controller

import (
	"github.com/rcsrn/rmp/pkg/database"
	"github.com/rcsrn/rmp/pkg/miner"
	"fmt"
)

func Run() {
	obtainData()
}

func obtainData() {
	miner := miner.CreateNewMiner("")
	miner.Traverse()
	miner.MineTags()
	rolas := miner.GetRolas()

	builder := database.CreateNewBuilder()
	builder.SetRolas(rolas)
	database := builder.BuildDataBase()
	fmt.Println("it is all right here.", database)
}
