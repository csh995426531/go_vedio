package main

import (
	"net/http"
	"vedio/defs"
	"vedio/session"
)

var Header_Field_Session = "X-Session-Id"
var Header_Field_Uname = "X-User-Name"

// ValidateUserSession 验证用户session
func ValidateUserSession(r *http.Request) bool {
	sid := r.Header.Get(Header_Field_Session)
	if len(sid) == 0 {
		return false
	}

	uname, exp := session.IsSessionExpired(sid)
	if exp {
		return false
	}

	r.Header.Add(Header_Field_Uname, uname)
	return true
}

// ValidateUser
func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(Header_Field_Uname)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}

	return true
}
