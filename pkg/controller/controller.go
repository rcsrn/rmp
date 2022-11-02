package controller

import (
	//"github.com/rcsrn/rmp/pkg/database"
	"github.com/rcsrn/rmp/pkg/miner"
	//"fmt"
)

func Run() {
	obtainData()
}

func obtainData() {
	miner := miner.CreateNewMiner("Escuela/Modelado/Rolas/")
	miner.Traverse()
}
