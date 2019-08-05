package main

import (
	"io"
	"net/http"

	"github.com/csh995426531/go_vedio/scheduler/dbops"
	"github.com/julienschmidt/httprouter"
)

func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")

	if len(vid) == 0 {
		sendErrorResponse(w, http.StatusBadRequest, "Video id should not be empty")
		return
	}

	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Success")
	return
}
