package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	router  *httprouter.Router
	limiter *ConnLimiter
}

func (middle middleWareHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if !middle.limiter.GetConn() {
		sendErrorResponse(writer, http.StatusTooManyRequests, "Too Many Request")
	}
	middle.router.ServeHTTP(writer, request)
	defer middle.limiter.ReleaseConn()
}
func NewMiddleWareHandler(router *httprouter.Router, conCurrent int) http.Handler {
	middle := middleWareHandler{}
	middle.router = router
	middle.limiter = InitConnLimiter(conCurrent)
	return middle
}

func main() {
	router := RegisterHandlers()
	//中间劫持
	middleRoute := NewMiddleWareHandler(router, 2)
	http.ListenAndServe("0.0.0.0:9000", middleRoute)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:video_id", streamHandler)

	router.POST("/upload/:video_id", uploadHandler)

	router.GET("/upload/form", skipToUpLoadPageHandler)

	return router
}
