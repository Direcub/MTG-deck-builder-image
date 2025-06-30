# Welcome

This app is the culmination of my study on [Boot.dev](https://boot.dev). For my final project, I decided to merge two passions: **Magic: The Gathering** and **coding**. I wrote this to demonstrate proficiency in setting up backend servers and working with frontend data transfer.

# Front-end Feature Set

The webapp supports creation and deletion of decks, exporting to either Cockatrice (.cod) or MTGO (.txt) formats, as well as a robust deck editor with full rules enforcement.

# Back-end Endpoints

- **GET /** — Landing page; supports calls to various other endpoints.
- **POST /newdeck/** — Creates a new deck within either the server's filesystem or the database (not yet implemented).
- **GET /listdecks/** — Lists all decks in the filesystem (or database). Mostly for server functions; called automatically on the landing page.
- **GET /editor** — Changes the page to the editor and initializes it.
- **GET /searchcards** — Searches Scryfall for cards. Normally has built-in filtering through query concatenation.
- **GET /deckinfo** — Used by the server to pass information about the current deck.
- **POST /savedeck** — Saves the deck and any changes made to it in the filesystem.
- **DELETE /deletedeck** — Permanently removes the specified deck from the filesystem.
- **GET /exportdeck** — Formats and starts download of the current deck.

---
