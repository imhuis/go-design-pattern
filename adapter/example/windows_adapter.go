package example

import "fmt"

type Computer interface {
	InsertIntoLightningPort()
}

type Mac struct{}

func (m *Mac) mountUSB() {
	fmt.Println("Mac mounted USB")
}

type Windows struct{}

func (w *Windows) insertIntoUSB() {
	fmt.Println("USB inserted to Windows")
}

type WindowsAdapter struct {
	Windows
}

type MacAdapter struct {
	Mac
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	w.Windows.insertIntoUSB()
}

func (m *MacAdapter) name() {
	m.Mac.mountUSB()
}
