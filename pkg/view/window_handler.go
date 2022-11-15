package view

import (
	_"github.com/rcsrn/rmp/internal/res"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2"
	_"log"
	_"fmt"
)

type WindowHandler struct {
	filePath    string
	app         fyne.App
	window      fyne.Window
	playButton  *widget.Button
	loadButton  *widget.Button
	muteButton  *widget.Button
	backButton  *widget.Button
	nextButton  *widget.Button
	loopButton  *widget.Button
	stopButton  *widget.Button
	volumeBar   *widget.Slider
	musicSlider *widget.Slider
}

func CreateNewWindowHandler() *WindowHandler {
	return &WindowHandler{
		filePath: "" ,
		app: app.New(),
		window: nil,
		playButton: nil,
		loadButton: nil,
		muteButton: nil,
		backButton: nil,
		nextButton: nil,
		loopButton: nil,
		stopButton: nil,
		volumeBar: nil,
		musicSlider: nil,
	}
}

func (handler *WindowHandler) GetFilePath() string {
	return handler.filePath
}


func (handler *WindowHandler) InitializeWindow() {
	load := handler.app.NewWindow("RMP")
	load.Resize(fyne.NewSize(800, 400))
	load.CenterOnScreen()
	
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter the directory")
	
	loadButton := widget.NewButton("Load Files", func() {})

	content := container.NewGridWithRows(3,
		layout.NewSpacer(),
		container.NewGridWithColumns(3,
			layout.NewSpacer(),
			container.NewVBox(input, loadButton),
			layout.NewSpacer()),
		layout.NewSpacer())
	load.SetContent(content)

	handler.loadButton = loadButton
	handler.window = load
}



// func (handler *WindowHandler) ShowLoadWindow() {
// 	loadWindow := handler.app.NewWindow("RMP")
// 	loadWindow.Resize(fyne.NewSize(800, 400))
// 	loadWindow.CenterOnScreen()

// 	input := widget.NewEntry()
// 	input.SetPlaceHolder("Enter the directory")
	
// 	loadButton := widget.NewButton("Load Files", func() {
// 		if input.Text == "" || string(input.Text[0]) != "/" {
// 			handler.use()
// 		} else {
// 			handler.filePath = input.Text
// 			handler.ShowPrincipalWindow()
// 			loadWindow.Close()
// 		}
// 	})

// 	handler.loadButton = loadButton
	
// 	content := container.NewGridWithRows(3,
// 		layout.NewSpacer(),
// 		container.NewGridWithColumns(3,
// 			layout.NewSpacer(),
// 			container.NewVBox(input, loadButton),
// 			layout.NewSpacer()),
// 		layout.NewSpacer())

// 	loadWindow.SetContent(content)
	
// 	fmt.Println(loadButton.OnTapped)
	
// 	loadWindow.Show()
// }


func (handler *WindowHandler) RunApp() {
	handler.app.Run()
}


func (handler *WindowHandler) InitializePrincipalWindow() {
	principalWindow := handler.app.NewWindow("RMP")
	principalWindow.Resize(fyne.NewSize(800, 400))
	principalWindow.SetMaster()
	principalWindow.CenterOnScreen()

	searchBar := widget.NewEntry()
	searchBar.SetPlaceHolder("Search...")

	top := container.NewVBox(searchBar)

	controls := handler.createControls()
	musicBar := handler.createMusicBar()

	bottom := container.NewGridWithRows(2,
		controls,
		musicBar)

	content := container.NewBorder(top, bottom, nil, nil)
	
	principalWindow.SetContent(content)
}


func (handler *WindowHandler) createControls() *fyne.Container {
	handler.playButton =  widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {})
	handler.muteButton = widget.NewButtonWithIcon("", theme.VolumeUpIcon(), func() {})
	handler.backButton = widget.NewButtonWithIcon("", theme.MediaSkipPreviousIcon(), func() {})
	handler.nextButton = widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), func() {})
	handler.loopButton = widget.NewButtonWithIcon("", theme.MediaReplayIcon(), func() {})
	handler.stopButton = widget.NewButtonWithIcon("", theme.MediaStopIcon(), func() {})

	volumeBar := handler.createVolumeBar()
	
	return container.NewGridWithColumns(7,
		handler.backButton,
		handler.playButton,
		handler.nextButton,
		handler.muteButton,
		handler.stopButton,
		handler.loopButton,
		volumeBar,
	)
}


// func (handler *WindowHandler) ShowPrincipalWindow() {

// 	principalWindow := handler.app.NewWindow("RMP")
// 	principalWindow.Resize(fyne.NewSize(800, 400))
// 	principalWindow.SetMaster()
// 	principalWindow.CenterOnScreen()

// 	searchBar := widget.NewEntry()
// 	searchBar.SetPlaceHolder("Search...")

// 	top := container.NewVBox(searchBar)

// 	var playButton *widget.Button
// 	var muteButton *widget.Button
// 	var loopButton *widget.Button
	
// 	playButton = widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
// 		if playButton.Icon == theme.MediaPlayIcon() {
// 			playButton.SetIcon(theme.MediaPauseIcon())
// 		} else {
// 			playButton.SetIcon(theme.MediaPlayIcon())
// 		}
// 	})

// 	backButton := widget.NewButtonWithIcon("", theme.MediaSkipPreviousIcon(), func() {

// 	})

// 	nextButton := widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), func() {
		
// 	})

// 	muteButton = widget.NewButtonWithIcon("", theme.VolumeUpIcon(), func() {
// 		if muteButton.Icon == theme.VolumeUpIcon() {
// 			muteButton.SetIcon(theme.VolumeMuteIcon())
// 		} else {
// 			muteButton.SetIcon(theme.VolumeUpIcon())
// 		}
// 	})
// 	stopButton := widget.NewButtonWithIcon("", theme.MediaStopIcon(), func() {})
// 	loopButton = widget.NewButtonWithIcon("", theme.MediaReplayIcon(), func() {
// 		var icon *fyne.StaticResource = res.ResourceRepeatLightPng

// 		if loopButton.Icon == theme.MediaReplayIcon() {
// 			loopButton.SetIcon(icon)
// 		} else {
// 			loopButton.SetIcon(theme.MediaReplayIcon())
// 		}
// 	})

// 	handler.playButton = playButton
// 	handler.muteButton = muteButton
// 	handler.backButton = backButton
// 	handler.nextButton = nextButton
// 	handler.loopButton = loopButton
// 	handler.stopButton = stopButton

// 	musicBar := handler.createMusicBar()
// 	volumeBar := handler.createVolumeBar()
	
// 	bottom := container.NewGridWithRows(2,
// 		container.NewGridWithColumns(7,
// 			backButton,
// 			playButton,
// 			nextButton,
// 			muteButton,
// 			stopButton,
// 			loopButton,
// 			volumeBar),
// 		musicBar)
	
// 	content := container.NewBorder(top, bottom, nil, nil)

// 	principalWindow.SetContent(content)
	
// 	principalWindow.Show()
// }

func (handler *WindowHandler) createMusicBar() *fyne.Container {
	progressSlider := widget.NewSlider(0, 100)
	currentTime := widget.NewLabel("00:00")
	endTime := widget.NewLabel("00:00")

	sliderHolder := container.NewBorder(nil, nil, currentTime, endTime, progressSlider)

	handler.musicSlider = progressSlider
	
	return sliderHolder
}

func (handler *WindowHandler) createVolumeBar() *widget.Slider {
	volumeSlider := widget.NewSlider(0, 100)
	volumeSlider.Orientation = widget.Horizontal
	handler.volumeBar = volumeSlider
	return volumeSlider
}


func (handler *WindowHandler) createDisplay(information []string) *fyne.Container {
	return nil
}


func (handler *WindowHandler) use() {
	use := widget.NewLabel("please enter a valid path directory")
	
	useWindow := handler.app.NewWindow("Warning")
	useWindow.Resize(fyne.NewSize(10, 5))
	useWindow.CenterOnScreen()

	content := container.NewGridWithRows(3,
		layout.NewSpacer(),
		container.NewGridWithColumns(1,
			use),
		layout.NewSpacer())
	
	useWindow.SetContent(content)
	useWindow.Show()
}

func (handler *WindowHandler) ShowError(error string) {
	errorMessage := widget.NewLabel(error)
	
	errorWindow :=  handler.app.NewWindow("ERROR")
	errorWindow.Resize(fyne.NewSize(10, 5))
	errorWindow.CenterOnScreen()
	
	content := container.NewGridWithRows(3,
		layout.NewSpacer(),
		container.NewGridWithColumns(1,
			errorMessage),
		layout.NewSpacer())

	errorWindow.SetContent(content)

	errorWindow.Show()

	handler.app.Quit()
}


func (handler *WindowHandler) ChangePlayButtonIcon() int {
	if handler.playButton.Icon == theme.MediaPlayIcon() {
		handler.playButton.SetIcon(theme.MediaPauseIcon())
	} else {
		handler.playButton.SetIcon(theme.MediaPlayIcon())
	}
}

func (handler *WindowHandler) OnPlay(action func()) {
	handler.playButton.OnTapped = action
}
