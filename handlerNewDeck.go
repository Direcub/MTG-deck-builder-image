package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func handlercreatedeck(w http.ResponseWriter, r *http.Request) {
	deckName := chi.URLParam(r, "deckname")

	err := os.WriteFile(fmt.Sprintf("decks/&s.txt", deckName), []byte(""), 0644)
	if err != nil {
		http.Error(w, "could bnote create deck", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Deck createed: " + deckName))
}
