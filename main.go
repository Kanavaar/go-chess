package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func createGrid() *fyne.Container {
	grid := container.NewGridWithColumns(8)

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			bg := canvas.NewRectangle(color.Gray{0x30})
			if y%2 == x%2 {
				bg.FillColor = color.Gray{0xE0}
			}

			grid.Add(bg)
		}
	}
	return grid
}

func main() {
	a := app.New()
	w := a.NewWindow("chess")

	grid := createGrid()
	w.SetContent(grid)
	w.Resize(fyne.NewSize(480, 480))
	w.ShowAndRun()
}
