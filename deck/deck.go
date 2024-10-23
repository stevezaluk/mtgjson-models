package deck

const (
	MAINBOARD = "mainBoard"
	SIDEBOARD = "sideBoard"
	COMMANDER = "commanderBoard"
)

/*
Deck - Represents a MTGJSON deck

Code (string) - A 3 or 4 digit code as an identifier for the deck
Commander (slice[string]) - A list of UUID's that represents the commander for the deck
Mainboard (slice[string]) - A list of UUID's that represents the main board for the deck
Name (string) - The name of the deck
ReleaseDate (string) - The release date of the deck
Sideboard (slice[string]) - A list of UUID's that represents the side board for the deck
Type (string) - The deck type
*/
type Deck struct {
	Code        string   `json:"code"`
	Commander   []string `json:"commander"`
	MainBoard   []string `json:"mainBoard"`
	Name        string   `json:"name"`
	ReleaseDate string   `json:"releaseDate"`
	SideBoard   []string `json:"sideBoard"`
	Type        string   `json:"type"`
}
