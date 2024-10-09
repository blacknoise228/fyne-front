package apisends

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type loginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type userResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}
type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

var UserCabinet *loginUserResponse

func FyneAuthUser(a fyne.App) {

	w := a.NewWindow("Welcome")

	loginForm := widget.NewForm(
		widget.NewFormItem("Username", widget.NewEntry()),
		widget.NewFormItem("Password", widget.NewPasswordEntry()),
	)

	loginForm.OnSubmit = func() {
		if ok := SendAPIAuth(loginForm.Items[0].Widget.(*widget.Entry).Text,
			loginForm.Items[1].Widget.(*widget.Entry).Text, w); ok {
			AuthOkNext(w)
		}
	}
	content := container.NewVBox(
		widget.NewLabel("Login"),
		loginForm,
	)
	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 400))
	w.ShowAndRun()
}
func SendAPIAuth(login string, password string, w fyne.Window) bool {
	user := loginUserRequest{Username: login, Password: password}
	data, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return false
	}
	req, err := http.NewRequest(http.MethodPost, URLS.LoginUserPOST, bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return false
	}
	req.Header.Set("content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	json.Unmarshal(body, &UserCabinet)
	log.Println("OK")
	if resp.StatusCode != http.StatusOK {
		dialog.ShowInformation("No Authorize", "Invalid password or username", w)
		return false
	}
	return true
}
