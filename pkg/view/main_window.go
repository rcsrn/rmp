package view

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2"
	"log"
)

func RunMainWindow() {
	myApp := app.New()
	myWindow := myApp.NewWindow("RMP")
	myWindow.SetMaster()
	defaultSize := fyne.NewSize(800, 400)
	myWindow.Resize(defaultSize)

	
	input := widget.NewEntry()
	input.SetPlaceHolder("Search...")
	
	content := container.NewVBox(input, widget.NewButton("Save", func() {
		log.Println("Content was:", input.Text)
	}))

	myWindow.SetContent(content)
	
	myWindow.ShowAndRun()
}


func createEntry() {
}
