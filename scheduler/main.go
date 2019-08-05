package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/csh995426531/go_vedio/scheduler/taskrunner"
)

func registerHander() *httprouter.Router {
	router := httprouter.New()

	router.DELETE("/video-delete-record/:vid-id", vidDelRecHandler)

	return router
}

func main() {

	go taskrunner.Start()
	r := registerHander()
	http.ListenAndServe(":8002", r)
}
