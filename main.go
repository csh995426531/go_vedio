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

	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	fmt.Println("start")
	r := RegisterHandlers()
	http.ListenAndServe(":8082", r)
	fmt.Println("end")
}
