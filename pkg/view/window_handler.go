package view

import (
	"github.com/rcsrn/rmp/internal/res"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
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
	loadWindow  fyne.Window
	loadHolder  *widget.Entry
	playButton  *widget.Button
	loadButton  *widget.Button
	muteButton  *widget.Button
	backButton  *widget.Button
	nextButton  *widget.Button
	loopButton  *widget.Button
	stopButton  *widget.Button
	volumeBar   *widget.Slider
	musicSlider *widget.Slider
	list        *widget.List
	playList    *[]string
}

func CreateNewWindowHandler() *WindowHandler {
	return &WindowHandler{
		filePath: "" ,
		app: app.New(),
		loadHolder: nil,
		loadWindow: nil,
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

func (handler *WindowHandler) SetPlayList(playList *[]string) {
	handler.playList = playList
}

func (handler *WindowHandler) InitializeLoadWindow() {
	load := handler.app.NewWindow("RMP")
	load.Resize(fyne.NewSize(800, 400))
	load.CenterOnScreen()
	
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter the directory")

	handler.loadHolder = input
	
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
	handler.loadWindow = load
}


func (handler *WindowHandler) RunApp() {
	handler.loadWindow.ShowAndRun()
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

	center := handler.createPlayList(handler.playList)
	
	content := container.NewBorder(top, bottom, nil, nil, center)
	
	principalWindow.SetContent(content)

	principalWindow.Show()
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

func (handler *WindowHandler) GetLoadText() string{
	return handler.loadHolder.Text
}


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


func (handler *WindowHandler) createPlayList(nameSongs *[]string) *widget.List {
	data := binding.BindStringList(nameSongs)
	
	playList := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})
	handler.list = playList
	
	return playList
}


func (handler *WindowHandler) Use() {
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
}

func (handler *WindowHandler) CloseLoadWindow() {
	handler.loadWindow.Close()
}

func (handler *WindowHandler) ChangePlayButtonIcon() int {
	if handler.playButton.Icon == theme.MediaPlayIcon() {
		handler.playButton.SetIcon(theme.MediaPauseIcon())
		return 0
	} else {
		handler.playButton.SetIcon(theme.MediaPlayIcon())
		return 1
	}
}

func (handler *WindowHandler) ChangeMuteButtonIcon() int {
	if handler.muteButton.Icon == theme.VolumeUpIcon() {
		handler.muteButton.SetIcon(theme.VolumeMuteIcon())
		return 0
	} else {
		handler.muteButton.SetIcon(theme.VolumeUpIcon())
		return 1
	}
}

func (handler *WindowHandler) ChangeLoopButtonIcon() int {
	var icon *fyne.StaticResource = res.ResourceRepeatLightPng
	
	if handler.loopButton.Icon == theme.MediaReplayIcon() {
		handler.loopButton.SetIcon(icon)
		return 0
	} else {
		handler.loopButton.SetIcon(theme.MediaReplayIcon())
		return 1
	}
	
}

func (handler *WindowHandler) OnLoad(action func()) {
	handler.loadButton.OnTapped = action
}

func (handler *WindowHandler) OnBack(action func()) {
	handler.backButton.OnTapped = action
}

func (handler *WindowHandler) OnPlay(action func()) {
	handler.playButton.OnTapped = action
}

func (handler *WindowHandler) OnNext(action func()) {
	handler.nextButton.OnTapped = action
}

func (handler *WindowHandler) OnMute(action func()) {
	handler.muteButton.OnTapped = action
}

func (handler *WindowHandler) OnLoop(action func()) {
	handler.loopButton.OnTapped = action
}

func (handler *WindowHandler) OnStop(action func()) {
	handler.stopButton.OnTapped = action
}

func (handler *WindowHandler) ChangePlayButton(isPlaying bool) {
	
}

func (handler *WindowHandler) ChangeMuteButton(isMuted bool) {
	
}

func (handler *WindowHandler) ChangeLoopButton(isLoop bool) {
	
}

func (handler *WindowHandler) OnSelect(action func(id int)) {
	handler.list.OnSelected = action
}

func (handler *WindowHandler) IsOnPlayButton() bool {
	return handler.playButton.Icon == theme.MediaPlayIcon()
}

func (handler *WindowHandler) SelectPreviousItem(currentItem string) string{
	var j int
	for i, rola := range (*handler.playList) {
		if rola == currentItem {
			j = i
		}
	}

	if j == 0 {
		handler.list.Select(j)
		return (*handler.playList)[j]
	} else {
		if j - 1 < 0 {
			return ""
		}
		handler.list.Select(j - 1)
		return (*handler.playList)[j - 1]
	}
}

func (handler *WindowHandler) SelectNextItem(currentItem string) string{
	var j int
	for i, rola := range (*handler.playList) {
		if rola == currentItem {
			j = i
		}
	}
	
	if j == len (*handler.playList) {
		handler.list.Select(j)
		return (*handler.playList)[j]
	} else {
		if j + 1 == len (*handler.playList) {
			return ""
		}
		handler.list.Select(j + 1)
		return (*handler.playList)[j + 1]
	}
}
