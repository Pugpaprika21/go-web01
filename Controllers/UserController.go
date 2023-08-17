package UserController

import (
	dbDns "go_crud_01/config"
	"html/template"
	"log"
	"net/http"
)

type Users struct {
	UserID    int
	UserName  string
	UserPass  string
	UserToken string
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	conn, _ := dbDns.Connect()
	rows, err := conn.Query("select UserID, UserName, UserPass, UserToken from users")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var users []Users

	for rows.Next() {
		var user Users
		err := rows.Scan(&user.UserID, &user.UserName, &user.UserPass, &user.UserToken)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	tmpl := template.Must(template.ParseFiles("template/home.html"))
	tmpl.Execute(w, users)
}
