# MTG Deck Builder

This app is the culmination of my study on [Boot.dev](https://boot.dev). It supports building, exporting, and storage of Magic: The Gathering (MTG) decks

For my final project, I decided to merge two passions: **Magic: The Gathering** and **coding**. I wanted to ensure that i had access to a deck builder that both **had access to all cards**, and **was low maintainence.** 

I source card information from scryfall itself so that theres never a card database to update on the user end, and made it a dockerizable app that loads onto the browser to help with compatibility.

## Front-end Feature Set

- Create and delete decks
- Export decks to Cockatrice (.cod) or MTGO (.txt) formats
- Robust deck editor with full rules enforcement

## Back-end Endpoints

| Method | Endpoint         | Description                                                                 |
|--------|------------------|-----------------------------------------------------------------------------|
| GET    | `/`              | Landing page; supports calls to various other endpoints                      |
| POST   | `/newdeck/`      | Creates a new deck within the server's filesystem or the database (TBD)      |
| GET    | `/listdecks/`    | Lists all decks in the filesystem (or database); called automatically on load|
| GET    | `/editor`        | Changes the page to the editor and initializes it                            |
| GET    | `/searchcards`   | Searches Scryfall for cards, with built-in filtering                         |
| GET    | `/deckinfo`      | Used by the server to pass information about the current deck                |
| POST   | `/savedeck`      | Saves the deck and any changes made to it in the filesystem                  |
| DELETE | `/deletedeck`    | Permanently removes the specified deck from the filesystem                   |
| GET    | `/exportdeck`    | Formats and starts download of the current deck                              |

![landing screen](images/landing.png)

![editor screen](images/editor.png)

## Tech Stack
I built this app using   
[Chi](https://github.com/go-chi/chi): For routing the server.     
[GO](https://go.dev/): Language the server is written in.    
[Go-scryfall](https://github.com/BlueMonday/go-scryfall?tab=readme-ov-file): API client for accessing Scryfall.        
JavaScript: Front-end logic and interactivity  



## Future Features

- Add EDHREC recommendations
- Add more filtering options

## Setup
1. Clone the repo
2. Run `go mod init`
3. Build and run the app with `go build` and `./MTG-deck-builder-image`
4. The app will open on http://localhost:8080/ by default, port can be set in the .env file.   

## Docker
If you'd prefer to use a docker image to handle dependancies and such, perform steps 1 and 2 of the setup and run

`docker build -t mtg-deck-builder .`

and 

`docker run -p 8080:8080 --env-file .env mtg-deck-builder`

the app will be available at http://localhost:8080

