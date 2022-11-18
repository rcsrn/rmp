package controller

import (
	"github.com/rcsrn/rmp/pkg/database"
	"github.com/rcsrn/rmp/pkg/miner"
	"github.com/rcsrn/rmp/pkg/view"
	
	"github.com/hajimehoshi/oto/v2"
	"github.com/hajimehoshi/go-mp3"

	"time"
	"os"
	"log"
	_"fmt"
	"os/user"
	_"errors"
)

type MainApp struct {
	handler *view.WindowHandler
	database *database.DataBase
	filePath string
	miner *miner.Miner
	isPlaying bool
	player oto.Player
	idCurrentRola int
	errorThrown bool
}

func createMainApp() *MainApp {

	return &MainApp{
		handler: view.CreateNewWindowHandler(),
		database: nil,
		filePath: "",
		miner: nil,
		errorThrown: false,
	}
}

func Run() {
	main := createMainApp()
	main.startView()	
}

func (main *MainApp) obtainData()  {
	main.errorThrown = false
	miner := miner.CreateNewMiner(main.filePath)
	main.miner = miner

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
	
	main.handler.InitializeLoadWindow()
	main.addLoadEvent()

	main.handler.RunApp()
}


func (main *MainApp) addLoadEvent() {
	main.handler.OnLoad(func() {
		format := main.handler.GetLoadText()
		
		if !main.isDirectoryPathFormat(format) {
			main.handler.Use()
		} else {
			main.filePath = format
			main.obtainData()
			if !main.errorThrown {
				main.handler.CloseLoadWindow()
				
				playList := main.obtainPlayList()
				main.handler.SetPlayList(playList)
				
				main.handler.InitializePrincipalWindow()
				main.addPrincipalEvents()
			}
		}
	})
}

func (main *MainApp) addPrincipalEvents() {
	main.handler.OnBack(func() {
		
	})
	
	main.handler.OnPlay(func() {
		
	})

	main.handler.OnNext(func() {
		
	})
	
	main.handler.OnMute(func() {
		
	})
	
	main.handler.OnLoop(func() {
		
	})

	main.handler.OnStop(func() {
		
	})

	main.handler.OnSelect(func(id int) {
		rola, err := main.database.QueryRola(int64(id))	
		main.check(err)

		idRola := int(rola.GetID())

		if main.isPlaying {
			main.player.Pause()
			if  idRola == main.idCurrentRola {
				return
			}
		}

		if main.handler.IsOnPlayButton() {
			main.handler.ChangePlayButtonIcon()
		}

		main.idCurrentRola = idRola
		
		file, err := os.Open(rola.GetPath())
		main.check(err)

		go main.playSong(file)
		
	})
}


func (main *MainApp) isDirectoryPathFormat(format string) bool {
	if format == "" || string(format[0]) != "/" {
		return false
	}
	return true
}


func (main *MainApp) getDBPath() string{
	user, err := user.Current()
	main.check(err)
	return user.HomeDir + "/.local/rmp"
}

func (main *MainApp) check(err error) {
	if err != nil {
		main.handler.ShowError(err.Error())
		main.errorThrown = true
		log.Print(err)		
	}
}

func (main *MainApp) obtainPlayList() *[]string {
	playList := make([]string, 0)

	rolas := main.miner.GetRolas()

	for _, rola := range(rolas) {
		title := rola.GetTitle()
		playList = append(playList, title)	
	}
	return &playList
}

func (main *MainApp) playSong(file *os.File) {
	decoder, err := mp3.NewDecoder(file)
	main.check(err)
	
	context, ready, err := oto.NewContext(decoder.SampleRate(), 2, 2)
		main.check(err)
	
	<- ready
	
	main.player = context.NewPlayer(decoder)
		defer main.player.Close()
	
	main.player.Play()

	main.isPlaying = true
	
	for {
		time.Sleep(time.Second)
		if !main.player.IsPlaying() {
			main.isPlaying = false
			break
		}
	}
	
}
