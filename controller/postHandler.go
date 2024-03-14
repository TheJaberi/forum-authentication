package forum

import (
	model "forum/model"
	"html/template"
	"net/http"
	"strconv"
)

// handles the creation of a post
func HandlerPost(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/post" {
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
	if model.ValidateSession(req) != nil {
		// log.Println()
	}
	var postCategories []int
	for i := 1; i <= len(model.AllCategories); i++ {
		categorytmp := req.FormValue(strconv.Itoa(i))
		if categorytmp != "" {
			postCategories = append(postCategories, i)
		}
	}
	err = model.CreatePost(req.FormValue("title"), req.FormValue("post"), postCategories)
	if err != nil {
		model.PostError2 = true
		model.AllData.PostErrorMsg = err.Error()
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "index.html", model.AllData)
}
