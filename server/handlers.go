package main

import (
	"net/http"

	"github.com/samueltej/todo"
)

type getAllHandler struct {
	dataFile string
}

func (h getAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var list todo.List
	if err := list.Get(h.dataFile); err != nil {
		errorReply(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	resp := todoResponse{Results: list}
	jsonReply(w, r, http.StatusOK, &resp)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorReply(w, r, http.StatusNotFound, "404 page not found")
		return
	}

	textReply(w, r, http.StatusOK, "Hello World!!")
}
