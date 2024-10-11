package apisends

import (
	"bytes"
	"encoding/json"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// CreateAccount creates a new Fyne window with a form for creating a new account.
// When the form is submitted, the sendCreateAccount function is called with the
// entered currency and the current window. If sendCreateAccount returns an error,
// an error dialog is shown. Otherwise, the window is closed and the AuthOkNext
// function is called.
func (app *FyneApp) CreateAccount() {
	w := app.Window
	registerForm := widget.NewForm(
		widget.NewFormItem("Currency", widget.NewSelect([]string{"USD", "KZT", "RUB"}, nil)),
	)
	registerForm.OnSubmit = func() {
		if err := sendCreateAccount(registerForm.Items[0].Widget.(*widget.Select).Selected,
			w); err != nil {
			dialog.ShowError(err, w)
		}
		app.UserHomePage()
	}
	back := widget.NewButton("Back", func() {
		app.UserHomePage()
	})
	content := container.NewVBox(
		widget.NewLabel("Create Account"),
		registerForm,
		back,
	)
	w.SetContent(content)
	w.Show()
}

// sendCreateAccount sends a POST request to the /accounts API endpoint with the
// given currency and the access token from the UserCabinet. If the request is
// successful, it shows a success dialog and returns nil. If the request fails,
// it shows an error dialog and returns the error.
func sendCreateAccount(currency string, w fyne.Window) error {
	authRequest := createAccountRequest{
		Currency: currency,
	}
	data, err := json.Marshal(authRequest)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, URLS.CreateAccountPOST, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("authorization", "Bearer "+UserCabinet.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 400: // Bad Request
		dialog.ShowInformation("Error", "Bad Request", w)
		return err
	case 500: // Internal Server Error
		dialog.ShowInformation("Error", "Internal Server Error", w)
		return err
	case 403: // Forbidden
		dialog.ShowInformation("Error", "Account already exists", w)
		return err
	}
	dialog.ShowInformation("Success", "Account created", w)
	return nil
}
