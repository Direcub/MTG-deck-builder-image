package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/BlueMonday/go-scryfall"
)

func (pass *Passthroughs) handlerSaveDeck(w http.ResponseWriter, r *http.Request) {
	type saveDeckRequest struct {
		Name      string          `json:"name"`
		Commander *scryfall.Card  `json:"commander"`
		Cards     []scryfall.Card `json:"cards"`
		Format    string          `json:"format"`
	}

	var req saveDeckRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON"))
		return
	}
	if req.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing deck name"))
		return
	}
	filename := "decks/" + req.Name + ".txt"

	var firstLine string
	if f, err := os.Open(filename); err == nil {
		defer f.Close()
		buf := make([]byte, 4096)
		n, _ := f.Read(buf)
		lines := string(buf[:n])
		if idx := indexOfNewline(lines); idx >= 0 {
			firstLine = lines[:idx+1]
		} else {
			firstLine = lines
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to create deck file"))
		return
	}
	defer f.Close()

	// Write the format tag as first line (preserve if present, else write new)
	if firstLine != "" {
		f.WriteString(firstLine)
	} else if req.Format != "" {
		f.WriteString("# format:" + req.Format + "\n")
	}
	// Write commander if present, using correct tag
	if req.Commander != nil && req.Commander.Name != "" {
		f.WriteString("# commander " + req.Commander.Name + "\n")
	}
	// Write each card name on its own line
	for _, card := range req.Cards {
		if card.Name != "" {
			f.WriteString(card.Name + "\n")
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deck saved"))
}

// indexOfNewline returns the index of the first newline character, or -1 if not found
func indexOfNewline(s string) int {
	for i, c := range s {
		if c == '\n' || c == '\r' {
			return i
		}
	}
	return -1
}
