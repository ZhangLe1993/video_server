package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"video_server/api/defs"
	"video_server/api/orm"
	"video_server/api/session"
)

func CreateUserHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	res, _ := ioutil.ReadAll(request.Body)
	body := &defs.User{}
	if err := json.Unmarshal(res, body); err != nil {
		SendErrorResponse(writer, defs.ErrorRequestBodyParseFailed)
		return
	}
	if err := orm.AddUser(body.UserName, body.Password); err != nil {
		SendErrorResponse(writer, defs.ErrorDBExec)
		return
	}
	sid := session.GenerateSessionId(body.UserName)
	sign := defs.SignedUp{true, sid}
	if resp, err := json.Marshal(sign); err != nil {
		SendErrorResponse(writer, defs.ErrorInternalFaults)
	} else {
		SendNormalResponse(writer, string(resp), 201)
	}
}

func LoginHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	uname := params.ByName("user_name")

	io.WriteString(writer, uname)
}
