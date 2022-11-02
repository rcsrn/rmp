package miner

import (
	"github.com/rcsrn/rmp/pkg/database"
	"path/filepath"
	"os/user"
	"log"
	"os"
	"strings"
	"fmt"
)

type Miner struct {
	directoryPath string
	filePaths []string
	rolas []*database.Rola
}

func CreateNewMiner(directoryPath string) *Miner {
	return &Miner{directoryPath, make([]string, 0), make([]*database.Rola, 0)}
}

func (miner *Miner) Traverse() {
	user, err := user.Current()
	if err != nil {
		log.Fatal("Something went wrong while retrieving user. %v",
			err)
	}
	musicPath := user.HomeDir + "/" + miner.directoryPath

	err = filepath.Walk(musicPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("It is not possible to acces",
				path, err)
			return err
		}
		
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".mp3") {
			miner.filePaths = append(miner.filePaths, path)
		}
		return nil
	})

	if err != nil {
		log.Fatal("Error while traversing the music path '%v'",
		musicPath)
	}

	fmt.Println(miner.filePaths)
}

func (miner *Miner) MineTags() {
	
}

func (miner *Miner) GetRolas() []*database.Rola {
	return miner.rolas
}
