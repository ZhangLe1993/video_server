package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := RegisterHandlers()

	http.ListenAndServe("0.0.0.0:9000", router)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:video_id", streamHandler)

	router.POST("/upload/:video_id", uploadHandler)

	return router
}
