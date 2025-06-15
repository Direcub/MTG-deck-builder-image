package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func (pass *Passthroughs) handlerListdecks(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("./decks")
	if err != nil {
		http.Error(w, "Unable to read deck directory", http.StatusInternalServerError)
		return
	}

	deckNames := []string{}
	for _, file := range files {
		if !file.IsDir() {
			deckNames = append(deckNames, file.Name())
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deckNames)
}
