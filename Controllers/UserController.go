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
	rows, err := conn.Query("SELECT UserID, UserName, UserPass, UserToken FROM users")
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

	tmpl := template.Must(template.New("home.gohtml").Funcs(template.FuncMap{"numRows": numRows}).ParseFiles("templates/layout/header.gohtml", "templates/layout/footer.gohtml", "templates/home.gohtml"))
	tmpl.Execute(w, users)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		username := r.FormValue("username")
		password := r.FormValue("password")
		usertoken := helper.RandomString(10)

		conn, _ := dbDns.Connect()
		stmt, err := conn.Prepare("INSERT INTO users (UserName, UserPass, UserToken) VALUES (?, ?, ?)")

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

func Edit(w http.ResponseWriter, r *http.Request) {
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

		detail := struct {
			MakePass string
		}{
			MakePass: helper.RandomString(10),
		}

		userDetail := struct {
			ID       int
			Username string
			Password string
			Token    string
			Detail   struct {
				MakePass string
			}
		}{
			ID:       user.UserID,
			Username: user.UserName,
			Password: user.UserPass,
			Token:    user.UserToken,
			Detail:   detail,
		}

		tmpl := template.Must(template.New("edit.gohtml").ParseFiles("templates/layout/header.gohtml", "templates/layout/footer.gohtml", "templates/edit.gohtml"))
		tmpl.Execute(w, userDetail)

	} else {
		fmt.Fprintf(w, "User not found")
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("userID")
	username := r.FormValue("username")
	password := r.FormValue("password")

	conn, _ := dbDns.Connect()
	_, err := conn.Exec("UPDATE users SET UserName = ?, UserPass = ? WHERE UserID = ?", username, password, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")

	conn, _ := dbDns.Connect()
	_, err := conn.Exec("DELETE FROM users WHERE UserID = ?", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
