package UserController

import (
	"fmt"
	dbDns "go_crud_01/config"
	helper "go_crud_01/func"
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

func numRows(index int) int {
	return index + 1
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

	tmpl := template.Must(template.New("home.html").Funcs(template.FuncMap{"numRows": numRows}).ParseFiles("template/home.html"))
	tmpl.Execute(w, users)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		username := r.FormValue("username")
		password := r.FormValue("password")
		usertoken := helper.RandomString(10)

		conn, _ := dbDns.Connect()
		stmt, err := conn.Prepare("insert into users (UserName, UserPass, UserToken) values (?, ?, ?)")

		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(username, password, usertoken)
		if err != nil {
			log.Fatal(err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func ViewUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")

	conn, _ := dbDns.Connect()
	rows, err := conn.Query("SELECT UserID, UserName, UserPass, UserToken FROM users WHERE UserID = ?", userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	if rows.Next() {
		var user Users
		if err := rows.Scan(&user.UserID, &user.UserName, &user.UserPass, &user.UserToken); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "User ID: %s\n", user.UserID)
		fmt.Fprintf(w, "User Name: %s\n", user.UserName)
		fmt.Fprintf(w, "User Password: %s\n", user.UserPass)
		fmt.Fprintf(w, "User Token: %s\n", user.UserToken)
	} else {
		fmt.Fprintf(w, "User not found")
	}
}
