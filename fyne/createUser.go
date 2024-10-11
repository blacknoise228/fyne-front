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

func (app *FyneApp) CreateUser() {
	w := app.Window
	registerForm := widget.NewForm(
		widget.NewFormItem("Username", widget.NewEntry()),
		widget.NewFormItem("FullName", widget.NewEntry()),
		widget.NewFormItem("Email", widget.NewEntry()),
		widget.NewFormItem("Password", widget.NewPasswordEntry()),
	)
	registerForm.OnSubmit = func() {
		if sendAPIRegistration(registerForm.Items[0].Widget.(*widget.Entry).Text,
			registerForm.Items[1].Widget.(*widget.Entry).Text,
			registerForm.Items[2].Widget.(*widget.Entry).Text,
			registerForm.Items[3].Widget.(*widget.Entry).Text,
			w) {
			app.FyneAuthUser()
		}
	}
	back := widget.NewButton("Back", func() {
		app.WelcomePage()
	})
	content := container.NewVBox(
		widget.NewLabel("Register"),
		registerForm,
		back,
	)
	w.SetContent(content)
	w.Show()
}

func sendAPIRegistration(username string, fullName string, email string, password string, w fyne.Window) bool {

	authRequest := createUserRequest{
		Username: username,
		Password: password,
		FullName: fullName,
		Email:    email,
	}
	data, err := json.Marshal(authRequest)
	if err != nil {
		return false
	}
	req, err := http.NewRequest(http.MethodPost, URLS.CreateUserPOST, bytes.NewBuffer(data))
	if err != nil {
		return false
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 403: // Unauthorized
		dialog.ShowInformation("Registration Failed", "User already exists", w)
		return false
	case 500: // Internal Server Error
		dialog.ShowInformation("Registration Failed", "Internal Server Error", w)
		return false
	case 400: // Bad Request
		dialog.ShowInformation("Registration Failed", "Bad Request", w)
		return false
	}

	dialog.ShowInformation("Registration Successful", "User created", w)
	return true
}
