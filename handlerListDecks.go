package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func (pass *Passthroughs) handlerListdecks(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat("./decks"); os.IsNotExist(err) {
		err := os.Mkdir("./decks", 0755)
		if err != nil {
			http.Error(w, "Unable to create deck directory", http.StatusInternalServerError)
			return
		}
	}

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
