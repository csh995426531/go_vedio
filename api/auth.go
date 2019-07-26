package main

import (
	"net/http"

	"github.com/csh995426531/go_vedio/api/defs"
	"github.com/csh995426531/go_vedio/api/session"
)

//HeaderFieldSession session标识
var HeaderFieldSession = "X-Session-Id"

//HeaderFieldUname user标识
var HeaderFieldUname = "X-User-Name"

// ValidateUserSession 验证用户session
func ValidateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HeaderFieldSession)
	if len(sid) == 0 {
		return false
	}

	uname, exp := session.IsSessionExpired(sid)
	if exp {
		return false
	}

	r.Header.Add(HeaderFieldUname, uname)
	return true
}

// ValidateUser 验证用户
func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HeaderFieldUname)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}

	return true
}
