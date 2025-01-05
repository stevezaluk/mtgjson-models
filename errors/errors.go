package errors

import "errors"

/*
Generic Issues - Holds all errors that are generic and dont relate to a specific card object
*/
var ErrInvalidObjectStructure = errors.New("request: Failed to bind object to request. Object structure may be incorrect")

/*
Card Errors - Holds all errors that could arise from fetching or inserting cards
*/
var ErrNoCard = errors.New("card: failed to find card with specified uuid")
var ErrNoCards = errors.New("card: No card found on index operation")
var ErrInvalidUUID = errors.New("card: Operation failed. Invalid v5 uuid")
var ErrCardMissingId = errors.New("card: Card is missing a name and/or a mtgjsonV4Id")
var ErrCardAlreadyExist = errors.New("card: Operation failed. Card already exists")
var ErrCardDeleteFailed = errors.New("card: Operation failed. Failed to remove card")
var ErrCardUpdateFailed = errors.New("card: Operation failed. Failed to update card with new metadata")
var ErrInvalidCards = errors.New("card: Could not update deck or set. Some cards are invalid or do not exist")

/*
Deck Errors - Holds all errors that could arise from fetching or inserting decks
*/
var ErrNoDeck = errors.New("deck: Failed to find deck with the specified code")
var ErrNoDecks = errors.New("deck: No decks found on index operation")
var ErrDeckUpdateFailed = errors.New("deck: Operation failed. Failed to update deck")
var ErrDeckMissingId = errors.New("deck: Operation Failed.. Deck is missing a name and/or a deck code")
var ErrDeckAlreadyExists = errors.New("deck: Operation Failed. Deck already exists")
var ErrDeckDeleteFailed = errors.New("deck: Operation failed. Issue while deleting deck")
var ErrBoardNotExist = errors.New("deck: Operation failed. Requested board does not exist")
var ErrDeckMissingContentIds = errors.New("deck: The contentId's field of the deck object is nil")
var ErrDeckNoCards = errors.New("deck: Could not update deck. No cards were passed to update the deck with")

/*
Set Errors - Holds all errors that could arise from fetching or creating sets
*/
var ErrNoSet = errors.New("set: failed to find set with the specified code")
var ErrNoSets = errors.New("set: No sets found on index operation")
var ErrSetUpdateFailed = errors.New("set: Operation failed. Failed to update set with new card(s)")
var ErrSetDeleteFailed = errors.New("set: Operation failed. Failed to remove cards from set")
var ErrSetMissingId = errors.New("set: Operation failed. Set is missing either a name or a set code")
var ErrSetAlreadyExists = errors.New("set: Operation failed. A set already exists under this set code")
var ErrSetMissingContentIds = errors.New("set: The contentId's field of the set object is nil")
var ErrSetNoCards = errors.New("set: Could not update set. No cards were passed to update the set with")

/*
User Errors - Holds all errors that could arise from fetching or creating new users
*/
var ErrNoUser = errors.New("user: Failed to find user with the specified username")
var ErrUserAlreadyExist = errors.New("user: Operation failed. A user already exists with this username")
var ErrUserDeleteFailed = errors.New("user: Failed to delete user")
var ErrUserUpdateFailed = errors.New("user: Failed to update user")
var ErrInvalidEmail = errors.New("user: Operation failed. Email is invalid")
var ErrInvalidPasswordLength = errors.New("user: Failed to create user. Password must be at least 12 characters long")
var ErrUserMissingId = errors.New("user: Operation failed. User model is missing Id")
var ErrFailedToRegisterUser = errors.New("user: Failed to register user with Auth0")
var ErrFailedToLoginUser = errors.New("user: Failed to login user with Auth0")

/*
Token Errors - Holds all errors that could arise from fetching or creating tokens
*/
var ErrTokenInvalid = errors.New("user: Access Token is not valid")
var ErrFailedToCreateToken = errors.New("token: Failed to create user access token")
var ErrInvalidPermissions = errors.New("user: You do not have the permissions to access this resource")

/*
Meta Errors - Holds all errors that could arise relating the mtgjsonApiMeta object
*/
var ErrMissingMetaApi = errors.New("meta: Object is missing the mtgjsonMetaApi field. This is usually due to creating the object manually")
var ErrMetaApiMustBeNull = errors.New("meta: The mtgjsonMetaApi field must be null. This will be created automatically")

/*
Credential Errors - Holds all errors that could arise from fetching or creating new user credentials
*/
var ErrNoCred = errors.New("userCredentials: Failed to find user credentials with the specified credentialId")
var ErrFailedToGenerateKey = errors.New("userCredentials: Failed to generate AES key for user credentials")
