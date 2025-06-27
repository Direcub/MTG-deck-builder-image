package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/BlueMonday/go-scryfall"
)

func (pass *Passthroughs) handlerEditor(w http.ResponseWriter, r *http.Request) {
	deck := r.URL.Query().Get("deck")

	if deck == "" {
		http.Error(w, "Deck not specified", http.StatusBadRequest)
		return
	}

	pass.working_Deck.File = "./decks/" + deck
	log.Printf("Working deck file: %s", pass.working_Deck.File)
	log.Printf("Loading deck: %s", pass.working_Deck.File)

	d, err := os.Open(pass.working_Deck.File)
	log.Printf("%v", err)
	if err != nil {
		http.Error(w, "cannot open deck file", http.StatusInternalServerError)
		return
	}

	defer d.Close()

	scanner := bufio.NewScanner(d)
	format := ""
	commander := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "# format:") {
			format = strings.TrimSpace(strings.TrimPrefix(line, "# format:"))
		}
		if strings.HasPrefix(line, "# commander:") {
			commander = strings.TrimSpace(strings.TrimPrefix(line, "# commander:"))
		}
	}

	pass.working_Deck.Format = format
	if commander != "" {
		pass.working_Deck.Commander, err = pass.Client.GetCardByName(context.Background(), commander, false, scryfall.GetCardByNameOptions{})
		if err != nil {
			http.Error(w, "cannot load commander card", http.StatusInternalServerError)
			return
		}
	}

	// Set the deck name (without .txt) for clarity
	pass.working_Deck.Name = strings.TrimSuffix(deck, ".txt")
	log.Printf("handlerEditor: working_Deck set: Name=%q, Format=%q, File=%q", pass.working_Deck.Name, pass.working_Deck.Format, pass.working_Deck.File)

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
