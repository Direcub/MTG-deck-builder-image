package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type CockatriceZone struct {
	XMLName xml.Name         `xml:"zone"`
	Name    string           `xml:"name,attr"`
	Cards   []CockatriceCard `xml:"card"`
}
type CockatriceDeck struct {
	XMLName  xml.Name         `xml:"cockatrice_deck"`
	Version  string           `xml:"version,attr"`
	DeckName string           `xml:"deckname"`
	Zones    []CockatriceZone `xml:"zone"`
}
type CockatriceCard struct {
	Name   string `xml:"name,attr"`
	Number int    `xml:"number,attr"`
}

func (p *Passthroughs) handlerExportDeck(w http.ResponseWriter, r *http.Request) {
	deckName := r.URL.Query().Get("name")
	format := r.URL.Query().Get("format")
	if deckName == "" || format == "" {
		http.Error(w, "Missing deck name or format", http.StatusBadRequest)
		return
	}

	deckPath := filepath.Join("./decks/", deckName)
	if !strings.HasSuffix(deckPath, ".txt") {
		deckPath += ".txt"
	}

	file, err := os.Open(deckPath)
	if err != nil {
		http.Error(w, "Deck not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read deck file", http.StatusInternalServerError)
		return
	}

	lines := strings.Split(string(data), "\n")

	var (
		mainCards []string
		commander string
	)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(strings.ToLower(line), "# format:") {
			continue
		} else if strings.HasPrefix(strings.ToLower(line), "# commander") {
			commander = strings.TrimSpace(strings.TrimPrefix(line, "# commander"))
		} else if line != "" && !strings.HasPrefix(line, "#") {
			mainCards = append(mainCards, line)
		}
	}

	baseName := strings.TrimSuffix(deckName, ".txt")

	switch strings.ToLower(format) {
	case "cockatrice":
		w.Header().Set("Content-Type", "application/xml")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.cod\"", baseName))
		writeCockatriceDeck(w, baseName, mainCards, commander)
	case "mtgo":
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.txt\"", baseName))
		writeMTGODeck(w, mainCards, commander)
	default:
		http.Error(w, "Unknown export format", http.StatusBadRequest)
	}
}

func writeCockatriceDeck(w http.ResponseWriter, deckName string, cards []string, commander string) {
	log.Printf("Writing Cockatrice deck: %s with %d cards", deckName, len(cards))
	mainMap := make(map[string]int)
	sideMap := make(map[string]int)
	commander = strings.TrimSpace(commander)
	for _, card := range cards {
		card = strings.TrimSpace(card)
		if commander != "" && strings.EqualFold(card, commander) {
			// Do NOT add to mainMap
			// Only count for sideMap below
			continue
		}
		mainMap[card]++
	}
	if commander != "" {
		sideMap[commander] = 1
	}
	var mainList, sideList []CockatriceCard
	for name, count := range mainMap {
		mainList = append(mainList, CockatriceCard{Name: name, Number: count})
	}
	for name, count := range sideMap {
		sideList = append(sideList, CockatriceCard{Name: name, Number: count})
	}
	deck := CockatriceDeck{
		Version:  "1",
		DeckName: deckName,
		Zones: []CockatriceZone{
			{Name: "main", Cards: mainList},
			{Name: "side", Cards: sideList},
		},
	}
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	if err := enc.Encode(deck); err != nil {
		log.Printf("Error encoding Cockatrice deck: %v", err)
	}
}

func writeMTGODeck(w http.ResponseWriter, cards []string, commander string) {
	cardMap := make(map[string]int)
	for _, card := range cards {
		card = strings.TrimSpace(card)
		cardMap[card]++
	}
	for name, count := range cardMap {
		fmt.Fprintf(w, "%d %s\n", count, name)
	}
	if commander != "" {
		fmt.Fprintf(w, "COMMANDER: %s\n", commander)
	}
}
