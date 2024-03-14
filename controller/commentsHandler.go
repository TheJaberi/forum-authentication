package forum

import (
	model "forum/model"
	"html/template"
	"net/http"
)

// handles the creation of a comment
func HandlerComments(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/comment" {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.Method != "POST" {
		ErrorHandler(w, req, http.StatusMethodNotAllowed)
		return
	}
	if model.ValidateSession(req) != nil {
		// log.Println()
	}
	t, err := template.ParseFiles("../view/postpage.html")
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError)
		return
	}
	postData, err := model.CreateComment(req.FormValue("commentContent"), req.FormValue("postid"))
	if err != nil {
		ErrorHandler(w, req, http.StatusBadRequest)
		return
	}
	err = model.GetPosts()
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError)
		return
	}
	postData.LoggedUser = true
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "postpage.html", model.AllData.Postpage)
}
