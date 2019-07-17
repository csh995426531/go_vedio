package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CreateUser 创建用户
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "create user hander")
}

// Login 用户登录
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Param) {
	io.WriteString(w, "login")
}

// DeleteUser 用户注销
func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Param) {
	io.WriteString(w, "delete user")
}
