package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func (pass *Passthroughs) handlercreatedeck(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "could not create deck", http.StatusInternalServerError)
		return
	}

	deckName := r.Form.Get("deckname")
	log.Printf(deckName)

	err = os.WriteFile(fmt.Sprintf("decks/%s.txt", deckName), []byte(""), 0644)
	if err != nil {
		http.Error(w, "could not create deck", http.StatusInternalServerError)
		return
	}

	// add header reading to check format

	pass.working_Deck = Deck{name: deckName, file: fmt.Sprintf("./decks/%s.txt", deckName)}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Deck created: " + deckName))
}
