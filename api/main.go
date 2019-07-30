package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter" //绑定uri
)

type middleWareHandler struct {
	r *httprouter.Router
}

// RegisterHandlers 注册处理方法
// *httprouter.Router
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser) //创建用户

	router.POST("/user/:user_name", Login) //用户登录

	router.DELETE("/user/:user_name", DeleteUser) //用户注销

	return router
}

// newMiddleWareHandler 新的处理方法
func newMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

// 拦截 httprouter.Router 包的 ServeHTTP 方法
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	ValidateUserSession(r)

	m.r.ServeHTTP(w, r)
}

func main() {
	fmt.Print("start\n")
	r := RegisterHandlers()
	handler := newMiddleWareHandler(r)
	http.ListenAndServe(":8000", handler)

	fmt.Print("end\n")
}
