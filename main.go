package main

import (
	apisends "fynetest/fyne"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	apisends.StartApp(app)
}
