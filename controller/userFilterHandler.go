package forum

import (
	model "forum/model"
	"html/template"
	"net/http"
)

// handles all the users filters
func HandlerMyFilter(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/mylikes/" && req.URL.Path != "/myposts/" && req.URL.Path != "/mydislikes/" {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		ErrorHandler(w, req, http.StatusMethodNotAllowed)
		return
	}
	if model.ValidateSession(req) != nil {
		// log.Println()
	}
	t, err := template.ParseFiles("../view/index.html")
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError)
		return
	}
	// err = model.FilterUserData(req.FormValue("user"), req.URL.Path) // depending on the url path that is displayed when each of the buttons is clicked the data well be filtered
	err = model.FilterUserData(req.URL.Path) // depending on the url path that is displayed when each of the buttons is clicked the data well be filtered
	if err != nil {
		ErrorHandler(w, req, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "index.html", model.AllData)
}
