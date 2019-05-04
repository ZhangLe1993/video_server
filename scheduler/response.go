package main

import (
	"io"
	"net/http"
)

func sendResponse(writer http.ResponseWriter, sc int, resp string) {
	writer.WriteHeader(sc)
	io.WriteString(writer, resp)
}
