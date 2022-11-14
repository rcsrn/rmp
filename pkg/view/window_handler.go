package view

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2"
	_"log"
	_"fmt"
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

func (handler *WindowHandler) GetFilePath() string {
	return handler.filePath
}

func (handler *WindowHandler) ShowLoadWindow() {
	loadWindow := handler.app.NewWindow("RMP")
	loadWindow.Resize(fyne.NewSize(800, 400))
	loadWindow.SetMaster()

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter the directory")
	
	
	loadButton := widget.NewButton("Load Files", func() {
		handler.filePath = input.Text
	})

	content := container.NewGridWithRows(3,
		layout.NewSpacer(),
		container.NewGridWithColumns(3,
			layout.NewSpacer(),
			container.NewVBox(input, loadButton),
			layout.NewSpacer()),
		layout.NewSpacer())

	loadWindow.SetContent(content)
	
	loadWindow.Show()
}


func (handler *WindowHandler) RunApp() {
	handler.app.Run()
}



