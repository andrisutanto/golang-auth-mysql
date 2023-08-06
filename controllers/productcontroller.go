package controllers

import (
	"html/template"
	"net/http"

	"appgo/config"
)

func ProductIndex(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	// jika sessionnya kosong, maka redirect ke halaman login
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			data := map[string]interface{}{
				"nama_lengkap": session.Values["nama_lengkap"],
			}

			temp, _ := template.ParseFiles("views/products/index.html")
			temp.Execute(w, data)
		}
	}
}
