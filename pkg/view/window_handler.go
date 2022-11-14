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
	loadWindow.CenterOnScreen()

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter the directory")
	
	
	loadButton := widget.NewButton("Load Files", func() {
		if input.Text == "" || string(input.Text[0]) != "/" {
			handler.use()
		} else {
		handler.filePath = input.Text
		handler.ShowPrincipalWindow()
			loadWindow.Close()
		}
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

func (handler *WindowHandler) ShowPrincipalWindow() {
	principalWindow := handler.app.NewWindow("RMP")
	principalWindow.Resize(fyne.NewSize(800, 400))
	principalWindow.SetMaster()
	principalWindow.CenterOnScreen()

	searchBar := widget.NewEntry()
	searchBar.SetPlaceHolder("Search...")

	top := container.NewVBox(searchBar)
	
	content := container.NewBorder(top, nil, nil, nil)

	principalWindow.SetContent(content)
	
	principalWindow.Show()
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
