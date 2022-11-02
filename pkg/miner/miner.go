package miner

import (
	"github.com/rcsrn/rmp/pkg/database"
)

type Miner struct {
	filePaths []string
	songs []*database.Rola
}

func createNewMiner() *Miner {
	return Miner{make([]string, 0), make([]*Rolas, 0)}
}

func(miner *Miner) Traverse() {
	
}

func(miner *Miner) MineTags() {
	
}
