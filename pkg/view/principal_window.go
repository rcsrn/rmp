package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

type PrincipalWindow struct {
	window  fyne.Window
	content *fyne.Container
	top     *fyne.Container
	bottom  *fyne.Container
	left    *fyne.Container
	right   *fyne.Container
	center  *widget.List
}

func createPrincipalWindow(window fyne.Window, content *fyne.Container, top *fyne.Container, bottom *fyne.Container, left *fyne.Container, right *fyne.Container, center *widget.List) *PrincipalWindow {
	return &PrincipalWindow{window, content, top, bottom, left, right, center}
}

func (principal *PrincipalWindow) UpdateDisplay(nameSongs *[]string) {
	principal.setPlayList(nameSongs)
	newContent := container.NewBorder(principal.top, principal.bottom, nil, nil, principal.center)
	principal.window.SetContent(newContent)
}

func (principal *PrincipalWindow) setPlayList(nameSongs *[]string) {
	data := binding.BindStringList(nameSongs)
	
	playList := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})
	principal.center = playList
}

