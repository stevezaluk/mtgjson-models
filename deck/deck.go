package deck

import (
	"github.com/stevezaluk/mtgjson-models/card"
	"github.com/stevezaluk/mtgjson-models/errors"
	"slices"
)

const (
	MAINBOARD = "mainBoard"
	SIDEBOARD = "sideBoard"
	COMMANDER = "commanderBoard"
)

/*
Deck - Represents a MTGJSON deck

Name (string) - The name of the deck
Code (string) - A 3 or 4 digit code as an identifier for the deck
Type (string) - The deck type
ReleaseDate (string) - The release date of the deck
CommanderIds (slice[string]) - A list of UUID's that represents the commander for the deck
MainboardIds (slice[string]) - A list of UUID's that represents the main board for the deck
SideboardIds (slice[string]) - A list of UUID's that represents the side board for the deck
CommanderContents (slice[card.Card]) - A list of card models that represent the cards of the commander board
MainboardContents (slice[card.Card]) - A list of card models that represent the cards of the main board
SideBoardContents (slide[Card.Card]) - A list of card models that represent the side board
*/
type Deck struct {
	Name        string       `json:"name"`
	Code        string       `json:"code"`
	Type        string       `json:"type"`
	ReleaseDate string       `json:"releaseDate"`
	Commander   []string     `json:"commander"`
	Mainboard   []string     `json:"mainBoard"`
	Sideboard   []string     `json:"sideBoard"`
	Contents    DeckContents `json:"contents"`
}

type DeckContents struct {
	Commander []card.Card `json:"commander"`
	Mainboard []card.Card `json:"mainBoard"`
	Sideboard []card.Card `json:"sideBoard"`
}

/*
GetIdBoard - Returns a pointer to the slice that represents ids in the requested board

Parameters:
board (string) - The board you want a pointer too

Returns
*slice[string] - The board the caller requested
*/
func (d *Deck) GetIdBoard(board string) *[]string {
	if board == MAINBOARD {
		return &d.Mainboard
	} else if board == SIDEBOARD {
		return &d.Sideboard
	} else if board == COMMANDER {
		return &d.Commander
	}

	return nil
}

/*
GetIdBoard - Returns a pointer to the slice that represents ids in the requested board

Parameters:
board (string) - The board you want a pointer too

Returns
*slice[string] - The board the caller requested
*/
func (d *Deck) GetBoard(board string) *[]card.Card {
	if board == MAINBOARD {
		return &d.Contents.Mainboard
	} else if board == SIDEBOARD {
		return &d.Contents.Sideboard
	} else if board == COMMANDER {
		return &d.Contents.Commander
	}

	return nil
}

/*
CardExists - Ensure that a card exists on a specific board using a UUID

Parameters:
uuid (string) - The uuid to check
board (string) - The board you want to check. Can be either: mainBoard, sideBoard, commanderBoard
*/
func (d Deck) CardExists(uuid string, board string) (bool, error) {
	sourceBoard := d.GetIdBoard(board)
	if sourceBoard == nil {
		return false, errors.ErrBoardNotExist
	}
	var ret = false

	for _, val := range *sourceBoard {
		if val == uuid {
			ret = true
			break
		}
	}
	return ret, nil
}

/*
AllCards - Combine all boards into a list of UUID's

Parameters: None

Returns:
allCard ([]string) - A list of all UUID's in the deck
*/
func (d Deck) AllCardIds() []string {
	var allCards []string

	allCards = append(d.Mainboard, d.Sideboard...)
	allCards = append(allCards, d.Commander...)

	return allCards
}

/*
AddCards - Add a list of cards to a specific board within a deck.
Card validation is not performed here as it is performed before this operation

Parameters:
uuids (slice[string]) - A list of UUID's you want to add to your deck
board (string) - The board you want to add to

Returns
errors.ErrBoardNotExist - If the board does not exist
*/
func (d *Deck) AddCards(uuids []string, board string) error {
	sourceBoard := d.GetIdBoard(board)
	if sourceBoard == nil {
		return errors.ErrBoardNotExist
	}

	*sourceBoard = append(*sourceBoard, uuids...)

	return nil
}

/*
DeleteCards - Delete a list of cards to a specific board within a deck.
Card validation is not performed here as it is performed before this operation

Parameters:
uuids (slice[string]) - A list of UUID's you want to remove from your deck
board (string) - The board you want to add to

Returns
errors.ErrBoardNotExist - If the board does not exist
*/
func (d *Deck) DeleteCards(uuids []string, board string) error {
	sourceBoard := d.GetIdBoard(board)
	if sourceBoard == nil {
		return errors.ErrBoardNotExist
	}

	for _, uuid := range uuids {
		for i, val := range *sourceBoard {
			if uuid == val {
				*sourceBoard = slices.Delete(*sourceBoard, i, i+1)
			}
		}
	}

	return nil
}
