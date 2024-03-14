package forum

import (
	model "forum/model"
	"html/template"
	"net/http"
	"strconv"
)

// handler postpage handles the post clicked on in the main page
func HandlerPostPage(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/postpage/" {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.Method != "GET" || model.AllData.AllPosts == nil {
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

	postID, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		ErrorHandler(w, req, http.StatusBadRequest)
		return
	}
	if postID > len(model.AllPosts) {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	model.AllData.Postpage = model.AllPosts[postID-1]
	model.AllData.Postpage.LoggedUser = model.AllData.IsLogged
	model.AllData.Postpage.Port = model.AllData.Port
	if model.AllData.IsLogged {
		for i := 0; i < len(model.AllData.Postpage.Comments); i++ { // for loop to set all comments in the post to logged to show the like buttons
			model.AllData.Postpage.Comments[i].CommentLoggedUser = model.AllData.IsLogged
		}
	}
	w.WriteHeader(http.StatusOK)
	model.GetUserCommentInteractions()
	t.ExecuteTemplate(w, "postpage.html", model.AllData.Postpage)
}
