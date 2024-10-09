package apisends

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

var Accounts []Account

// Next window after autorization is a accounts a user
func AuthOkNext(w fyne.Window) {
	listAccounts()
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
	scroll := container.NewVScroll(list)
	scroll.SetMinSize(fyne.NewSize(400, 200))
	content := container.NewVBox(
		widget.NewLabel(fmt.Sprintf("User %v", UserCabinet.User.FullName)),
		scroll,
	)
	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 400))
	w.Show()
}
func listAccounts() {
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
	log.Println(req)
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
	log.Println("OK")
}
