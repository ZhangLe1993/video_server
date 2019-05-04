package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/scheduler/task"
)

func main() {

	//c := make(chan int)
	go task.Start()
	router := RegisterHandlers()
	http.ListenAndServe("0.0.0.0:9001", router)
	//<- c
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video/delete/video_id", videoDelHandler)
	return router
}
