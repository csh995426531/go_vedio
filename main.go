package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RegisterHandlers 注册处理
// *httprouter.Router
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser) //创建用户

	router.POST("/user/:user_name", Login) //用户登录

	router.DELETE("/user/:user_name", DeleteUser) //用户注销

	return router
}

func main() {
	fmt.Println("start")
	r := RegisterHandlers()
	http.ListenAndServe(":8088", r)
	fmt.Println("end")
}
	return new
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	ValidateUserSession(r)

	m.r.ServeHTTP(w, r)
}

func main() {

	r := RegisterHandlers()
	handler := newMiddleWareHandler(r)
	http.ListenAndServe(":8088", handler)

}
