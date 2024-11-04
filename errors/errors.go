package errors

import "errors"

/*
Card Errors - Holds all errors that could arise from fetching or inserting cards
*/
var ErrNoCard = errors.New("card: failed to find card with specified uuid")
var ErrNoCards = errors.New("card: No card found on index operation")
var ErrInvalidUUID = errors.New("card: invalid v5 uuid")
var ErrCardMissingId = errors.New("card: Card is missing a name and/or a mtgjsonV4Id")
var ErrCardAlreadyExist = errors.New("card: Could not complete operation. Card already exists")
var ErrCardDeleteFailed = errors.New("card: Failed to delete deck. Unknown server issue")
var ErrCardUpdateFailed = errors.New("card: Failed to update card with new metadata")

/*
Deck Errors - Holds all errors that could arise from fetching or inserting decks
*/
var ErrNoDeck = errors.New("deck: failed to find deck with the specified code")
var ErrNoDecks = errors.New("deck: No deck found on index operation")
var ErrDeckUpdateFailed = errors.New("deck: Failed to update deck with new card")
var ErrDeckMissingId = errors.New("deck: Failed to create deck. Deck is missing a name and/or a deck code")
var ErrDeckAlreadyExists = errors.New("deck: Failed to create deck. Deck already exists")
var ErrDeckDeleteFailed = errors.New("deck: Failed to delete deck. Unknown server issue")
var ErrBoardNotExist = errors.New("deck: Failed to add card to deck. Requested board does not exist")

/*
Set Errors - Holds all errors that could arise from fetching or creating sets
*/
var ErrNoSet = errors.New("set: failed to find set with the specified code")
