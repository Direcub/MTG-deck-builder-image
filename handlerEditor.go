package main

import (
	"io"
	"net/http"
)

func (pass *Passthroughs) handlerEditor(w http.ResponseWriter, r *http.Request) {
	deck := r.URL.Query().Get("deck")

	if deck == "" {
		http.Error(w, "Deck not specified", http.StatusBadRequest)
		return
	}

	pass.working_Deck.file = deck

	f, err := staticFiles.Open("html/editor.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer f.Close()
	if _, err := io.Copy(w, f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
