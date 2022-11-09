package database

import (
	"strings"
)

type Rola struct {
	id int64
	performer string
	album string
	path string
	title string
	track int
	year int
	genre string
}

func CreateNewRola() *Rola {
	return &Rola{
		id: 0,
		performer: "<Unknown>",
		album: "<Unknown>",
		path: "<Unknown>",
		title: "<Unknown>",
		track: 0,
		year: 0,
		genre: "<Unknown>",
	}
}

func (rola *Rola) GetID() int64 {
	return rola.id
}

func (rola *Rola) GetPerformer() string {
	return rola.performer
}

func (rola *Rola) GetAlbum() string {
	return rola.album
}

func (rola *Rola) GetPath() string {
	return rola.path
}

func (rola *Rola) GetTitle() string {
	return rola.title
}

func (rola *Rola) GetTrack() int {
	return rola.track
}

func (rola *Rola) GetYear() int {
	return rola.year
}

func (rola *Rola) GetGenre() string {
	return rola.genre
}

func (rola *Rola) SetPerformer(performer string) {
	rola.performer = strings.TrimSpace(performer)
}

func (rola *Rola) SetTitle(title string) {
	rola.title = strings.TrimSpace(title)
}

func (rola *Rola) SetAlbum(album string) {
	rola.album = strings.TrimSpace(album)
}

func (rola *Rola) SetTrack(track int) {
	rola.track = track
}

func (rola *Rola) SetYear(year int) {
	rola.year = year
}

func (rola *Rola) SetGenre(genre string) {
	rola.genre = strings.TrimSpace(genre)
}

func (rola *Rola) SetPath(path string) {
	rola.path = strings.TrimSpace(path)
}

func (rola *Rola) SetID(id int64) {
	rola.id = id
}

