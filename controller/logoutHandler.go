package forum

import (
	m "forum/model"
	"html/template"
	"log"
	"net/http"
)

// handles the logout process
func HandlerLogout(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/logout/" {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		ErrorHandler(w, req, http.StatusMethodNotAllowed)
		return
	}
	t, err := template.ParseFiles("../view/index.html")
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp := m.RemoveSession(req)
	if resp != 0 {
		log.Println(resp)
	}
	m.AllData.LoggedUser = m.Empty
	m.AllData.IsLogged = false
	m.AllData.TypeAdmin = false
	m.LiveSession = m.EmptySession

	http.SetCookie(w, m.BlankCookie)
	t.ExecuteTemplate(w, "index.html", m.AllData)
}
