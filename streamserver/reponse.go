package main

import (
	"io"
	"net/http"
)

// sendErrorResponse 发送错误响应
func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}