package apisends

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (app *FyneApp) CreateTransaction() {
	w := app.Window
	transForm := widget.NewForm(
		widget.NewFormItem("Account ID", widget.NewEntry()),
		widget.NewFormItem("Currency", widget.NewSelect([]string{"USD", "KZT", "RUB"}, nil)),
		widget.NewFormItem("Amount", widget.NewEntry()),
	)
	transForm.OnSubmit = func() {
		if sendCreateTransaction(transForm.Items[0].Widget.(*widget.Entry).Text,
			transForm.Items[1].Widget.(*widget.Select).Selected,
			transForm.Items[2].Widget.(*widget.Entry).Text,
			w) {
			app.UserHomePage()
		}
	}
	back := widget.NewButton("Back", func() {
		app.UserHomePage()
	})
	content := container.NewVBox(
		widget.NewLabel("Create Transaction"),
		transForm,
		back,
	)
	w.SetContent(content)
	w.Show()
}

func sendCreateTransaction(accountID, currency, amount string, w fyne.Window) bool {
	toAccId, err := strconv.Atoi(accountID)
	if err != nil {
		dialog.ShowError(err, w)
		return false
	}
	am, err := strconv.Atoi(amount)
	if err != nil {
		dialog.ShowError(err, w)
		return false
	}
	transferRequest := createTransactionRequest{
		FromAccountID: idFromCurrency(currency),
		ToAccountID:   int64(toAccId),
		Currency:      currency,
		Amount:        int64(am),
	}
	data, err := json.Marshal(transferRequest)
	if err != nil {
		return false
	}
	req, err := http.NewRequest(http.MethodPost, URLS.CreateTransferPOST, bytes.NewBuffer(data))
	if err != nil {
		return false
	}
	req.Header.Set("authorization", "Bearer "+UserCabinet.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 400: // Bad Request
		dialog.ShowInformation("Error", "Bad Request", w)
		return false
	case 500: // Internal Server Error
		dialog.ShowInformation("Error", "Internal Server Error", w)
		return false
	}
	return true
}

func idFromCurrency(currency string) int64 {
	for _, account := range Accounts {
		if account.Currency == currency {
			return account.ID
		}
	}
	return 0
}
