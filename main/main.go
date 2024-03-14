package main

import (
	"fmt"
	controller "forum/controller"
	model "forum/model"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Whatever needs to load before the server starts (Files/APIs)
func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	controller.StaticFileLoader()
	model.DatabaseLoader()
	model.GetCategories()
	model.GetPosts()

}
func main() {
	const port = ":8080"
	var err error
	model.AllData.Port, err = strconv.Atoi(port[1:])
	if err != nil {
		log.Println(err.Error())
		os.Exit(0)
	}
	http.HandleFunc("/", controller.MainHandler)
	fmt.Println("http://localhost" + port)
	http.HandleFunc("/createcategory", controller.HandlerCreateCategory)
	http.HandleFunc("/postpage/", controller.HandlerPostPage)
	http.HandleFunc("/logout/", controller.HandlerLogout)
	http.HandleFunc("/comment", controller.HandlerComments)
	http.HandleFunc("/like/", controller.HandlerLikes)
	http.HandleFunc("/dislike/", controller.HandlerLikes)
	http.HandleFunc("/commentlike/", controller.HandlerCommentsLikes)
	http.HandleFunc("/commentdislike/", controller.HandlerCommentsLikes)
	http.HandleFunc("/register", controller.HandlerRegister)
	http.HandleFunc("/login", controller.HandlerLogin)
	http.HandleFunc("/post", controller.HandlerPost)
	log.Fatal(http.ListenAndServe(port, nil))
}
