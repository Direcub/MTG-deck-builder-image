package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/BlueMonday/go-scryfall"
)

func (pass *Passthroughs) handlerSearchCards(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "missing search query", http.StatusInternalServerError)
		return
	}

	hasCommander := r.URL.Query().Get("commander")
	if hasCommander == "false" {
		query += " is:commander"
	}

	results, err := searchCard(pass.Client, query, scryfall.SearchCardsOptions{
		Unique: scryfall.UniqueModeCards,
	})
	if err != nil {
		http.Error(w, "failed to search card", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func searchCard(client *scryfall.Client, query string, opts scryfall.SearchCardsOptions) ([]scryfall.Card, error) {
	ctx := context.Background()
	results, err := client.SearchCards(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	return results.Cards, nil
}
