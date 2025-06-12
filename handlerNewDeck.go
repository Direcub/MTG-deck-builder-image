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
	format := r.Form.Get("format")
	switch format {
	case "Commander":
		pass.format = format
	case "Standard":
		pass.format = format
	default:
		http.Error(w, "please select a format", http.StatusForbidden)
		return
	}
	log.Print(format)
	log.Print(deckName)

	err = os.WriteFile(fmt.Sprintf("decks/%s.txt", deckName), []byte(""), 0644)
	if err != nil {
		http.Error(w, "could not create deck", http.StatusInternalServerError)
		return
	}

	pass.working_Deck = Deck{name: deckName, file: fmt.Sprintf("./decks/%s.txt", deckName)}
	pass.format = format

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Deck created: " + deckName))
}
