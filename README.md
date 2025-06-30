Your README is clear, concise, and covers the essentials! For a hiring manager, a little extra polish and formatting will make it even more professional and inviting. Here are some suggestions:

---

# MTG Deck Builder

This app is the culmination of my study on [Boot.dev](https://boot.dev). For my final project, I decided to merge two passions: **Magic: The Gathering** and **coding**. I wrote this to demonstrate proficiency in setting up backend servers and working with frontend data transfer.

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

## Future Features

- Add EDHREC recommendations
- Add more filtering options
