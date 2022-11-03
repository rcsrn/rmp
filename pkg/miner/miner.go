package miner

import (
	"github.com/rcsrn/rmp/pkg/database"
	"path/filepath"
	"log"
	"os"
	"strings"
	"errors"
	"id3v2"
)

type Miner struct {
	directoryPath string
	filePaths []string
	rolas []*database.Rola
}

//CreateNewMiner creates a new Miner with two empty slices to be filled later and a directoryPath.
func CreateNewMiner(directoryPath string) *Miner {
	return &Miner{directoryPath, make([]string, 0), make([]*database.Rola, 0)}
}

//Traverse traverses the miner's directoryPath searching for mp3 files. When a mp3 file is found
//its path is stored in the miner's filePaths.
func (miner *Miner) Traverse() {
	err := filepath.Walk(miner.directoryPath, func(path string, info os.FileInfo, err error) error {
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
		miner.directoryPath)
	}
}

//MineTags searchs for ID3v2.4 tags of each mp3 file through its path stored in the miner's filePaths.
//Once all tags of a file are obtained a Rola is created with that information.
func (miner *Miner) MineTags() error {
	for _, path := range miner.filePaths {
		tag, err := id3v2.Open(path, id3v2.Options{Parse: true})
		if err != nil {
			return errors.New("Error while obtaining file tags of '%v'. %v",
			path, err)
		}
		rola := database.createNewRola()
		rola.setPath(path)
		
		if tag.Artist() != "" {
			rola.SetPerformer(tag.Artist())
		}
		if tag.Title() != "" {
			rola.SetTitle(tag.Artist())
		}
		if tag.Album() != "" {
			rola.SetAlbum(tag.Artist())
		}
		if tag.Track() != 0 {
			rola.SetTrack(tag.Artist())
		}
		if tag.Year() != 0 {
			rola.SetYear(tag.Artist())
		}
		if tag.Genre() != "" {
			rola.SetGenre(tag.Artist())
		}
	}
}

func (miner *Miner) GetRolas() []*database.Rola {
	return miner.rolas
}

func (miner *Miner) GetFilePaths() []string {
	return miner.filePaths
}
