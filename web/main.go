package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", homeHandler)
	router.POST("/", homeHandler)
	router.GET("/user_home", userHomeHandler)
	router.POST("/user_home", userHomeHandler)
	router.POST("/api", apiHandler)
	router.POST("/upload/:video_id", proxyHandler)
	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))
	return router
}

func main() {
	router := RegisterHandler()
	http.ListenAndServe(":8080", router)
}
