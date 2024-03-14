package forum

import (
	model "forum/model"
	"html/template"
	"net/http"
)

// handles the registration process for the user
func HandlerRegister(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/register" {
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
	w.WriteHeader(http.StatusOK)

	var NewApplicant = model.Applicant{
		Username: req.FormValue("username"),
		Password: []byte(req.FormValue("password")),
		Email:    req.FormValue("email"),
	}
	err = model.UserRegisteration(NewApplicant, model.DB)
	if err != nil {
		model.LoginError2 = true
		model.AllData.LoginErrorMsg = err.Error()
	}
	t.ExecuteTemplate(w, "index.html", model.AllData)
}
