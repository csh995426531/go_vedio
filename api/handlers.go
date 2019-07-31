package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/csh995426531/go_vedio/api/dbops"
	"github.com/csh995426531/go_vedio/api/defs"
	"github.com/csh995426531/go_vedio/api/session"
	"github.com/julienschmidt/httprouter"
)

// CreateUser 创建用户
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// io.WriteString(w, "create user hander")

	res, _ := ioutil.ReadAll(r.Body)
	ubody := defs.UserCredential{}

	if err := json.Unmarshal(res, &ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFaild)
		return
	}

	pwd, err := dbops.GetUserCredential(ubody.Username)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	if len(pwd) > 0 {
		sendErrorResponse(w, defs.ErrorUserAlreadyExsitsError)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	sid, err := session.GenerateNewSessionID(ubody.Username)
	if err != nil {
		sendErrorResponse(w, defs.ErrorSessionError)
		return
	}

	su := defs.SignedUp{Success: true, SessionId: sid}

	if res, err = json.Marshal(&su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}
	sendNormalResponse(w, string(res), 201)
}

// Login 用户登录
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	uname := p.ByName("user_name")

	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFaild)
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	ubody := defs.Pwd{}

	if err := json.Unmarshal(res, &ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFaild)
		return
	}

	if pwd, err := dbops.GetUserCredential(uname); err != nil || pwd != ubody.Pwd {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}

	sid, err := session.GenerateNewSessionID(uname)
	if err != nil {
		sendErrorResponse(w, defs.ErrorSessionError)
		return
	}

	su := defs.SignedUp{Success: true, SessionId: sid}

	if res, err = json.Marshal(&su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}
	sendNormalResponse(w, string(res), 201)
}

// DeleteUser 用户注销
func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	uname := p.ByName("user_name")
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFaild)
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	ubody := defs.Pwd{}

	if err := json.Unmarshal(res, &ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFaild)
		return
	}

	if pwd, err := dbops.GetUserCredential(uname); err != nil || pwd != ubody.Pwd {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}

	if err := dbops.DeleteUserCredential(uname, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	sid := r.Header.Get(HeaderFieldSession)

	if len(sid) > 0 {
		session.DeleteSession(sid)
	}

	sendNormalResponse(w, "success", 200)
}
