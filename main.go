package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"time"

	Scryfall "github.com/BlueMonday/go-scryfall"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type Deck struct {
	Name          string          `json:"name"`
	Commander     Scryfall.Card   `json:"commander"`
	Cards         []Scryfall.Card `json:"cards"`
	ColorIdentity string          `json:"colorIdentity"`
	File          string          `json:"file"`
	Format        string          `json:"format"`
}

type Passthroughs struct {
	working_Deck Deck
	Client       *Scryfall.Client
}

//go:embed html/*
var staticFiles embed.FS

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("warning: assuming default configuration. .env unreadable: %v", err)
	}

	port := os.Getenv("PORT")
	log.Printf("port is %v", port)
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	psth := &Passthroughs{}

	psth.Client, err = Scryfall.NewClient()
	if err != nil {
		log.Fatalf("Failed to create Scryfall client: %v", err)
	}

	router.Get("/", psth.handlerLanding)
	router.Post("/newdeck/", psth.handlercreatedeck)
	router.Get("/listdecks/", psth.handlerListdecks)
	router.Get("/editor", psth.handlerEditor)
	router.Get("/searchcards", psth.handlerSearchCards)
	router.Get("/deckinfo", psth.handlerDeckInfo)
	router.Post("/savedeck", psth.handlerSaveDeck)
	router.Delete("/deletedeck", psth.handlerDeleteDeck)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: time.Minute,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
