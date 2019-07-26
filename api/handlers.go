package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/csh995426531/go_vedio/api/dbops"
	"github.com/csh995426531/go_vedio/api/defs"
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

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

}

// Login 用户登录
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "login")
}

// DeleteUser 用户注销
func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "delete user")
}
