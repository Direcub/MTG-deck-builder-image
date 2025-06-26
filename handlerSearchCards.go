package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/BlueMonday/go-scryfall"
)

func (pass *Passthroughs) handlerSearchCards(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "missing search query", http.StatusInternalServerError)
		return
	}

	hasCommander := r.URL.Query().Get("hasCommander")
	if hasCommander == "false" {
		query += " is:commander"
	}

	format := r.URL.Query().Get("format")
	if format == "Commander" || format == "Standard" {
		query += " f:" + format
	} else {
		log.Printf("Unsupported format: %s", format)
	}

	results, err := searchCard(pass.Client, query)
	if err != nil {
		http.Error(w, "failed to search card", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func searchCard(client *scryfall.Client, query string) ([]scryfall.Card, error) {
	ctx := context.Background()
	results, err := client.SearchCards(ctx, query, scryfall.SearchCardsOptions{
		Unique: scryfall.UniqueModeCards,
	})
	if err != nil {
		return nil, err
	}
	return results.Cards, nil
}
