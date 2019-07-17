package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RegisterHandlers 注册处理
// *httprouter.Router
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	router.DELETE("/user/:user_name", DeleteUser)

	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8082", r)
}
