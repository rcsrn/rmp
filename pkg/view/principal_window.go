package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

type PrincipalWindow struct {
	window      fyne.Window
	content     *fyne.Container
	top         *fyne.Container
	bottom      *fyne.Container
	left        *fyne.Container
	right       *fyne.Container
	newPlayList *widget.List
}

func createPrincipalWindow(window fyne.Window, content *fyne.Container, top *fyne.Container,
	bottom *fyne.Container, left *fyne.Container, right *fyne.Container) *PrincipalWindow {
	return &PrincipalWindow{
		window: window,
		content: content,
		top: top,
		bottom: bottom,
	}
}

func (principal *PrincipalWindow) UpdateDisplay(nameSongs *[]string) {
	principal.setNewPlayList(nameSongs) 
	newContent := container.NewBorder(principal.top, principal.bottom, nil, nil, principal.newPlayList)
	principal.window.SetContent(newContent)
}

func (principal *PrincipalWindow) setNewPlayList(nameSongs *[]string) {
	data := binding.BindStringList(nameSongs)
	
	playList := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})
	principal.newPlayList = playList
}

func (principal *PrincipalWindow) UpdateToGeneralPlayList() {
	principal.window.SetContent(principal.content)
}

func (principal *PrincipalWindow) OnSelect(action func (id int)) {
	principal.newPlayList.OnSelected = action
}
