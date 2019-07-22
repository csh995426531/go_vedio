package main

import (
	"encoding/json"
	"io"
	"net/http"
	"vedio/defs"
)

// sendErrorResponse 设置错误响应
func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse) {
	w.WriteHeader(errResp.HttpSC)

	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

// sendNormalResponse 设置正常响应
func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
