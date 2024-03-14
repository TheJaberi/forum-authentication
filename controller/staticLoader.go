package forum

import "net/http"

var HTMLs []string

// loads all the html, css and js files
func StaticFileLoader() {
	if HTMLs == nil {
		HTMLs = []string{
			"../view/index.html",
			"../view/error.html",
			"../view/postpage.html",
		}
	}
	cssFiles := http.FileServer(http.Dir("../view/css"))
	jsFiles := http.FileServer(http.Dir("../view/js"))
	imgFiles := http.FileServer(http.Dir("../view/img"))
	fontFiles := http.FileServer(http.Dir("../view/fonts"))
	http.Handle("/css/", http.StripPrefix("/css/", cssFiles))
	http.Handle("/js/", http.StripPrefix("/js/", jsFiles))
	http.Handle("/img/", http.StripPrefix("/img/", imgFiles))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", fontFiles))
}
