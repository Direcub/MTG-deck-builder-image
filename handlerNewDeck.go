package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func (pass *Passthroughs) handlercreatedeck(w http.ResponseWriter, r *http.Request) {
	deckName := chi.URLParam(r, "deckname")

	err := os.WriteFile(fmt.Sprintf("decks/%s.txt", deckName), []byte(""), 0644)
	if err != nil {
		http.Error(w, "could bnote create deck", http.StatusInternalServerError)
		return
	}

	pass.working_Directory = "decks/%s.txt"

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Deck created: " + deckName))
}
