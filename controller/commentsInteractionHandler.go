package forum

import (
	model "forum/model"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// handles the click on the like or dislike button for a comment
func HandlerCommentsLikes(w http.ResponseWriter, req *http.Request) {
	if !model.AllData.IsLogged {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.URL.Path != "/commentlike/" && req.URL.Path != "/commentdislike/" {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		ErrorHandler(w, req, http.StatusMethodNotAllowed)
		return
	}
	if model.ValidateSession(req) != nil {
		// log.Print()
	}
	t, err := template.ParseFiles("../view/postpage.html")
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = model.CommentInteraction(req.FormValue("commentInteraction"), req.FormValue("removeInteraction"), req.URL.Path)
	if err != nil {
		ErrorHandler(w, req, http.StatusBadRequest)
		return
	}
	t.ExecuteTemplate(w, "postpage.html", model.AllData.Postpage)
}
