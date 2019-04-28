package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//     Workspace/Go/src/com.biubiu/soup/video_server
//     Workspace/Go/bin

type middleWareHandler struct {
	router *httprouter.Router
}

func (handler middleWareHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	//check session
	validateUserSession(request)

	handler.router.ServeHTTP(writer, request)
}

func NewMiddleWareHandler(router *httprouter.Router) http.Handler {
	middleHandler := middleWareHandler{}
	middleHandler.router = router
	return middleHandler
}

func main() {
	route := RegisterHandler()
	//劫持
	handler := NewMiddleWareHandler(route)

	http.ListenAndServe("0.0.0.0:8000", handler)
}

func RegisterHandler() *httprouter.Router {

	router := httprouter.New()

	router.POST("/user", CreateUserHandler)

	router.POST("/user/:user_name", LoginHandler)

	return router
}
