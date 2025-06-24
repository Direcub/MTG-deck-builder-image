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

func (pass *Passthroughs) handlerCommanderSelect(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "missing search query", http.StatusInternalServerError)
		return
	}

	results, err := searchCard(pass.Client, query, scryfall.SearchCardsOptions{
		Unique: scryfall.UniqueModeCards,
	})
	if err != nil {
		http.Error(w, "failed to search card", http.StatusInternalServerError)
		return
	}

	commanders := make([]scryfall.Card, 0, len(results))

	for _, card := range results {
		if card.TypeLine != "" && (card.TypeLine == "Legendary Creature" || card.TypeLine == "Planeswalker") {
			commanders = append(commanders, card)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commanders)
}

func searchCard(client *scryfall.Client, query string, opts scryfall.SearchCardsOptions) ([]scryfall.Card, error) {
	ctx := context.Background()
	results, err := client.SearchCards(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	return results.Cards, nil
}
