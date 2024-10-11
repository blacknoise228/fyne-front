package apisends

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var UserCabinet *loginUserResponse

// FyneAuthUser creates a new Fyne window with a login form. If the user enters
// a valid username and password, the window is closed and the AuthOkNext
// function is called. Otherwise, an error dialog is shown.
func (a *FyneApp) FyneAuthUser() {

	w := a.Window

	loginForm := widget.NewForm(
		widget.NewFormItem("Username", widget.NewEntry()),
		widget.NewFormItem("Password", widget.NewPasswordEntry()),
	)

	loginForm.OnSubmit = func() {
		if ok := sendAPIAuth(loginForm.Items[0].Widget.(*widget.Entry).Text,
			loginForm.Items[1].Widget.(*widget.Entry).Text, w); ok {
			a.UserHomePage()
		}
	}
	back := widget.NewButton("Back", func() {
		a.WelcomePage()
	})
	content := container.NewVBox(
		widget.NewLabel("Login"),
		loginForm,
		back,
	)
	w.SetContent(content)
	w.Show()
}

// sendAPIAuth sends a POST request to the /users/login API endpoint with the
// given username and password. If the request is successful, it unmarshals the
// response body into UserCabinet and returns true. If the request fails, it
// shows an error dialog and returns false.
func sendAPIAuth(username string, password string, w fyne.Window) bool {
	authRequest := loginUserRequest{
		Username: username,
		Password: password,
	}
	data, err := json.Marshal(authRequest)
	if err != nil {
		return false
	}
	req, err := http.NewRequest(http.MethodPost, URLS.LoginUserPOST, bytes.NewBuffer(data))
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	err = json.Unmarshal(body, &UserCabinet)
	if err != nil {
		return false
	}
	if resp.StatusCode != http.StatusOK {
		dialog.ShowInformation("Authentication Failed", "Invalid password or username", w)
		return false
	}
	return true
}
