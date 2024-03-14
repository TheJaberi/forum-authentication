package forum

import (
	"html/template"
	"net/http"

	m "forum/model"
)

// handles the user logging in
func HandlerLogin(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/login" {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.Method != "POST" {
		ErrorHandler(w, req, http.StatusMethodNotAllowed)
		return
	}
	t, err := template.ParseFiles("../view/index.html")
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError)
		return
	}

	cookie, err := m.UserLogin(req.FormValue("email"), req.FormValue("password"))
	if err != nil {
		m.LoginError2 = true
	} else {
		m.LoginError2 = false
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "index.html", m.AllData)
}
