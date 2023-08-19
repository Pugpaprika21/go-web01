package main

import (
	UserController "go_crud_01/Controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/", UserController.GetUsers)
	http.HandleFunc("/show", UserController.ViewUser)
	http.ListenAndServe(":8080", nil)
}
