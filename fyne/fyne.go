package apisends

import "fyne.io/fyne/v2"

type APIAdress struct {
	CreateUserPOST     string
	LoginUserPOST      string
	CreateAccountPOST  string
	GetAccountGET      string
	ListAccountGET     string
	CreateTransferPOST string
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

func StartApp(app fyne.App) {
	InitURLS()
	FyneAuthUser(app)
}
