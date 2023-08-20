package main

import (
	UserController "go_crud_01/Controllers"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", UserController.GetUsers)
	http.HandleFunc("/create", UserController.Create)
	http.HandleFunc("/edit", UserController.Edit)
	http.HandleFunc("/update", UserController.Update)
	http.HandleFunc("/delete", UserController.Delete)
	http.ListenAndServe(":8080", nil)
}
