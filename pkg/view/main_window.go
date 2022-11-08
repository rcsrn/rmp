package view

import (
	"fyne.io/fyne/v2/app"
	_"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
)

func RunMainWindow() {
	myApp := app.New()
	myWindow := myApp.NewWindow("RMP")
	defaultSize := fyne.NewSize(800, 400)
	myWindow.Resize(defaultSize)
	
	myWindow.ShowAndRun()
}
