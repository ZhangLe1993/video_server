package main

import (
	"io"
	"net/http"
)

func sendErrorResponse(writer http.ResponseWriter, sc int, errMsg string) {
	writer.WriteHeader(sc)
	io.WriteString(writer, errMsg)
}
