package miner

import (
	"github.com/rcsrn/rmp/pkg/database"
)

type Miner struct {
	filePaths []string
	rolas []*database.Rola
}

func createNewMiner() *Miner {
	return Miner{make([]string, 0), make([]*Rolas, 0)}
}

func (miner *Miner) Traverse() {
	
}

func (miner *Miner) MineTags() {
	
}

func (miner *Miner) getRolas() []database.Rola {
	return miner.rolas
}
