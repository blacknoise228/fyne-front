package apisends

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var Accounts []Account

// UserHomePage creates a new Fyne window with a scrollable list showing the
// user's accounts. The list items show the account ID, currency, and balance.
// The window also has a button at the bottom to create a new account and transaction.
func (app *FyneApp) UserHomePage() {
	w := app.Window
	listAccounts(w)
	list := widget.NewList(
		func() int {
			return len(Accounts)
		},
		func() fyne.CanvasObject {
			label := widget.NewLabel("Account Info")
			label.Resize(fyne.NewSize(300, 100))
			return label
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(
				fmt.Sprintf(
					"Account ID: %d Currency: %s Balance: %d",
					Accounts[i].ID,
					Accounts[i].Currency,
					Accounts[i].Balance,
				))
		},
	)
	refresh := widget.NewButton("Refresh", func() {
		app.UserHomePage()
	})
	createAccount := widget.NewButton("Create Account", func() {
		app.CreateAccount()
	})
	createTransfer := widget.NewButton("Create Transaction", func() {
		app.CreateTransaction()
	})
	scroll := container.NewVScroll(list)
	scroll.SetMinSize(fyne.NewSize(400, 200))
	content := container.NewVBox(
		widget.NewLabel(fmt.Sprintf("User %v", UserCabinet.User.FullName)),
		refresh,
		scroll,
		createTransfer,
		createAccount,
	)
	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 400))
	w.Show()
}

// listAccounts sends a GET request to the /accounts API endpoint with the
// access token from the UserCabinet. If the request is successful, it
// unmarshals the response body into the Accounts variable and logs a success
// message. Otherwise, it logs an error message and shows an error dialog.
func listAccounts(w fyne.Window) {
	req, err := http.NewRequest(http.MethodGet, URLS.ListAccountGET, nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("authorization", "Bearer "+UserCabinet.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(body, &Accounts)
	if err != nil {
		log.Println(err)
		return
	}
	if len(Accounts) == 0 {
		dialog.ShowInformation("Error", "No accounts found", w)
	}
	log.Println("OK")
}
