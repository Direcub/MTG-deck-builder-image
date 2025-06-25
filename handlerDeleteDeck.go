package main

import (
	"net/http"
	"os"
)

func (pass *Passthroughs) handlerDeleteDeck(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing deck name", http.StatusBadRequest)
		return
	}
	filename := "decks/" + name
	if err := os.Remove(filename); err != nil {
		http.Error(w, "Failed to delete deck", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deck deleted"))
}
