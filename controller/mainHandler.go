package forum

import (
	m "forum/model"
	"html/template"
	"net/http"
)

// handles the main page
func MainHandler(w http.ResponseWriter, req *http.Request) {
	if m.ValidateSession(req) != nil {
		resp := m.RemoveSession(req)
		if resp != 0 {
			ErrorHandler(w, req, resp)
			return
		}
		http.SetCookie(w, m.BlankCookie)
	}
	if req.URL.Path != "/" {
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
	m.GetCategories()
	m.GetPosts()
	m.AllData.AllPosts = m.RSort(m.AllPosts)
	m.AllData.AllCategories = m.AllCategories
	m.AllData.CategoryCheck = true
	if req.FormValue("sortby") != "" {
		err = m.SortPosts(req.FormValue("sortby"))
		if err != nil {
			ErrorHandler(w, req, http.StatusNotFound)
			return
		}
	}

	if req.FormValue("category") != "" {
		err = m.FilterByCategory(req.FormValue("category"))
		if err != nil {
			ErrorHandler(w, req, http.StatusNotFound)
			return
		}
	}
	if m.LoginError2 {
		m.AllData.LoginError = true
	} else {
		m.AllData.LoginError = false
	}
	if m.PostError2 {
		m.AllData.PostError = true
	} else {
		m.AllData.PostError = false
	}
	t.ExecuteTemplate(w, "index.html", m.AllData)
	m.AllData.LoginError = false
	m.LoginError2 = false
	m.AllData.PostError = false
	m.PostError2 = false
	m.AllData.LoginErrorMsg = ""
}
