package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/scheduler/orm"
)

func videoDelHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	vid := params.ByName("video_id")

	if len(vid) == 0 {
		sendResponse(writer, 400, "video_id is should not be empty")
		return
	}
	err := orm.AddVideoDeletionRecord(vid)
	if err != nil {
		sendResponse(writer, 500, "Internal Server Error")
		return
	}
	sendResponse(writer, 200, "Success")
	return
}
