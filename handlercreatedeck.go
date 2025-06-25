package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func (pass *Passthroughs) handlercreatedeck(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "could not create deck", http.StatusInternalServerError)
		return
	}

	deckName := r.Form.Get("deckname")
	format := r.Form.Get("format")

	if deckName == "" || format == "" {
		http.Error(w, "Missing deck name or format", http.StatusBadRequest)
		return
	}

	os.MkdirAll("decks", os.ModePerm)

	filePath := filepath.Join("decks", deckName+".txt")

	f, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to create deck file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	header := fmt.Sprintf("# format: %s\n\n", format)
	_, err = f.WriteString(header)
	if err != nil {
		http.Error(w, "Unable to write to deck file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Deck '%s' created successfully!", deckName)
	w.Write([]byte("Deck created: " + deckName))
}
