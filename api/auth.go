package main

import (
	"net/http"
	"video_server/api/defs"
	"video_server/api/session"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

/**
session校验
*/
func validateUserSession(request *http.Request) bool {
	sid := request.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSessionExpire(sid)
	if ok {
		return false
	}

	request.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

func validateUser(writer http.ResponseWriter, request *http.Request) bool {
	uname := request.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		SendErrorResponse(writer, defs.ErrorNotAuthorUser)
		return false
	}
	return true
}
