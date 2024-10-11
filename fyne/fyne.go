package apisends

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type APIAdress struct {
	CreateUserPOST     string
	LoginUserPOST      string
	CreateAccountPOST  string
	GetAccountGET      string
	ListAccountGET     string
	CreateTransferPOST string
}
type FyneApp struct {
	App    fyne.App
	Window fyne.Window
}

// NewFyneApp creates a new FyneApp, which is a wrapper around the standard Fyne App type.
// It provides some additional methods for creating and managing user accounts.
func NewFyneApp(a fyne.App) *FyneApp {
	w := a.NewWindow("SimpleBank")
	return &FyneApp{App: a, Window: w}
}

var URLS *APIAdress

func InitURLS() {
	URLS = &APIAdress{
		CreateUserPOST:     "http://localhost:8080/users",
		LoginUserPOST:      "http://localhost:8080/users/login",
		CreateAccountPOST:  "http://localhost:8080/accounts",
		GetAccountGET:      "http://localhost:8080/accounts",
		ListAccountGET:     "http://localhost:8080/accounts?page_id=1&page_size=5",
		CreateTransferPOST: "http://localhost:8080/transfers",
	}
}

func (app *FyneApp) StartApp() {
	InitURLS()
	app.WelcomePage()
	app.Window.ShowAndRun()

}

func (app *FyneApp) WelcomePage() {
	w := app.Window
	login := widget.NewButton("Login", func() {
		app.FyneAuthUser()
	})
	register := widget.NewButton("Register", func() {
		app.CreateUser()
	})
	text := widget.NewLabel("Welcome to SimpleBank")
	content := container.NewVBox(text, login, register)
	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 400))
	w.Show()
}
