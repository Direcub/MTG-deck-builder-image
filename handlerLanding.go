package main

import (
	"io"
	"net/http"
)

func (pass *Passthroughs) handlerLanding(w http.ResponseWriter, r *http.Request) {
	f, err := staticFiles.Open("html/interface.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	if _, err := io.Copy(w, f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
