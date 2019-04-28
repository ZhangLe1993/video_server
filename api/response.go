package main

import (
	"encoding/json"
	"io"
	"net/http"
	"video_server/api/defs"
)

func SendErrorResponse(writer http.ResponseWriter, errResp defs.ErrorResponse) {

	writer.WriteHeader(errResp.HttpSC)

	res, _ := json.Marshal(&errResp.Error)

	io.WriteString(writer, string(res))
}

func SendNormalResponse(writer http.ResponseWriter, resp string, sc int) {
	writer.WriteHeader(sc)
	io.WriteString(writer, resp)
}
