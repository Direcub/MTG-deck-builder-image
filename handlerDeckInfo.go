package main

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/BlueMonday/go-scryfall"
)

func (pass *Passthroughs) handlerDeckInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if pass.working_Deck.File == "" {
		http.Error(w, "No deck loaded", http.StatusBadRequest)
		return
	}
	if pass.Client == nil {
		http.Error(w, "Scryfall client not initialized", http.StatusInternalServerError)
		return
	}
	file, err := os.Open(pass.working_Deck.File)
	if err != nil {
		http.Error(w, "Failed to open deck file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var format string
	var commanderName string
	var cardNames []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "# format:") {
			format = strings.TrimSpace(strings.TrimPrefix(line, "# format:"))
		} else if strings.HasPrefix(line, "# commander") {
			commanderName = strings.TrimSpace(strings.TrimPrefix(line, "# commander"))
		} else if line != "" && !strings.HasPrefix(line, "#") {
			cardNames = append(cardNames, line)
		}
	}
	if err := scanner.Err(); err != nil {
		http.Error(w, "Failed to read deck file", http.StatusInternalServerError)
		return
	}

	var commanderCard *scryfall.Card
	if commanderName != "" {
		card, err := pass.Client.GetCardByName(r.Context(), commanderName, false, scryfall.GetCardByNameOptions{})
		if err == nil {
			commanderCard = &card
		}
	}
	var deckCards []scryfall.Card
	for _, name := range cardNames {
		card, err := pass.Client.GetCardByName(r.Context(), name, false, scryfall.GetCardByNameOptions{})
		if err == nil {
			deckCards = append(deckCards, card)
		}
	}
	resp := struct {
		Format    string          `json:"format"`
		Commander *scryfall.Card  `json:"Commander"`
		Cards     []scryfall.Card `json:"cards"`
		Name      string          `json:"Name"`
	}{
		Format:    format,
		Commander: commanderCard,
		Cards:     deckCards,
		Name:      pass.working_Deck.Name,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode deck info", http.StatusInternalServerError)
		return
	}
}
