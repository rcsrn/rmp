package view

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2"
	_"log"
)

type WindowHandler struct {
	filePath string
	app fyne.App
}

func CreateNewWindowHandler() *WindowHandler {
	return &WindowHandler{
		filePath: "" ,
		app: app.New(),
	}
}

func (master *WindowHandler) GetFilePath() string {
	return master.filePath
}

func (master *WindowHandler) ShowLoadWindow() {
	loadWindow := master.app.NewWindow("RMP")
	loadWindow.Resize(fyne.NewSize(800, 400))
	loadWindow.SetMaster()

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter the directory")
	
	
	loadButton := widget.NewButton("Load Files", func() {
		master.filePath = input.Text
	})

	content := container.NewGridWithRows(3, layout.NewSpacer(), container.NewGridWithColumns(3, layout.NewSpacer(), container.NewVBox(input, loadButton), layout.NewSpacer()), layout.NewSpacer())

	loadWindow.SetContent(content)
	
	loadWindow.Show()
}


func (master *WindowHandler) RunApp() {
	master.app.Run()
}

