package main

import (
	apisends "fynetest/fyne"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	app := apisends.NewFyneApp(a)
	app.StartApp()
}
