package miner

import (
	"github.com/rcsrn/rmp/pkg/database"
	"path/filepath"
	"log"
	"os"
	"strings"
	"errors"
	"github.com/dhowden/tag"
	"fmt"
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
//Once all available tags of a file are obtained a Rola is created with that information.
//If a file does not contain title tag then MineTags defines its name as default title.
func (miner *Miner) MineTags() error {
	for _, path := range miner.filePaths {
		file, err := os.Open(path)
		if err != nil {
			errorStr := fmt.Sprintf("Error while opening '%v'. %v",
			file, err)
			return errors.New(errorStr)
		}

		tag, err := tag.ReadFrom(file)
		
		if err != nil {
			errorStr := fmt.Sprintf("Error while obtaining file tags of '%v'. %v",
			path, err)
			return errors.New(errorStr)
		}
		
		rola := database.CreateNewRola()
		
		rola.SetPath(path)
		if performer := tag.Artist(); performer != "" {
			rola.SetPerformer(performer)
		}
		if title := tag.Title(); title != "" {
			rola.SetTitle(title)
			fmt.Println("llega aca")
		} else {
			title = defaultTitle(path)
			rola.SetTitle(title)
		}
		if album := tag.Album(); album != "" {
			rola.SetAlbum(album)
		}
		if track, _ := tag.Track(); track != 0 {
			rola.SetTrack(track)
		}
		if year := tag.Year(); year != 0 {
			rola.SetYear(year)
		}
		if genre := tag.Genre(); genre != "" {
			rola.SetGenre(genre)
		}

		fmt.Println(rola.GetTitle())
	}
	return nil
}

//GetRolas returns the rolas.
func (miner *Miner) GetRolas() []*database.Rola {
	return miner.rolas
}

//GetFilepaths returns the file paths.
func (miner *Miner) GetFilePaths() []string {
	return miner.filePaths
}

//defaultTitle returns the file name as default title for a file with
//no title tag.
func defaultTitle(path string) string {
	var index int
	path = strings.Trim(path, ".mp3")
	for i := len(path) - 1; i > 0; i-- {
		if string(path[i]) == "/" {
			index = i + 1
			break
		}
	}
	return path[index:]
}
