package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHander struct {
	r *httprouter.Router
	l *ConnLimiter
}

// NewMiddleWareHander 中间处理
func NewMiddleWareHander(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHander{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

// RegisterHandlers 注册处理
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/vidoes/:vid-id", streamHandler)

	router.POST("/vidoes/:vid-id", uploadHandler)

	router.GET("/testpage", testPageHandler)

	return router
}

// 拦截http服务
func (m middleWareHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}

	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func main() {
	r := RegisterHandlers()
	m := NewMiddleWareHander(r, 20)
	http.ListenAndServe(":8001", m)

}
