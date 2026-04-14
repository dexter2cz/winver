package main

import (
	"bytes"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func getWindowsVersion() string {
	cmd := exec.Command("cmd", "/C", "ver")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "Chyba při zjišťování verze"
	}
	return out.String()
}

func main() {
	a := app.New()
	w := a.NewWindow("Windows verze")

	label := widget.NewLabel("Klikni na tlačítko")

	button := widget.NewButton("Zjistit verzi", func() {
		label.SetText(getWindowsVersion())
	})

	w.SetContent(container.NewVBox(label, button))
	w.Resize(fyne.NewSize(300, 200))

	w.ShowAndRun()
}
