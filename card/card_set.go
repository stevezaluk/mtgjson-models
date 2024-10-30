package card

import "github.com/stevezaluk/mtgjson-models/card/meta"

/*
Card - A model representing an MTGJSON Card

Ommiting card descriptions for brevity.
See: https://mtgjson.com/data-models/card/card-set/
*/
type Card struct {
	AsciiName               string                `json:"asciiName"`
	AttractionLights        []string              `json:"attractionLights"`
	ColorIdentity           []string              `json:"colorIdentity"`
	ColorIndicator          []string              `json:"colorIndicator"`
	Colors                  []string              `json:"colors"`
	ConvertedManaCost       int                   `json:"convertedManaCost"`
	Defense                 string                `json:"defense"`
	EDHRecRank              int                   `json:"edhrecRank"`
	EDHRecSaltiness         float64               `json:"edhrecSaltiness"`
	FaceConvertedManaCost   int                   `json:"faceConvertedManaCost"`
	FaceManaValue           int                   `json:"faceManaValue"`
	FaceName                string                `json:"faceName"`
	FirstPrinting           string                `json:"firstPrinting"`
	ForeignData             []card.ForeignData    `json:"foreignData"`
	Hand                    string                `json:"hand"`
	HasAlternativeDeckLimit bool                  `json:"hasAlternativeDeckLimit"`
	Identifiers             card.CardIdentifiers  `json:"identifiers"`
	IsFunny                 bool                  `json:"isFunny"`
	IsReserved              bool                  `json:"isReserved"`
	Keywords                []string              `json:"keywords"`
	Layout                  string                `json:"layout"`
	LeadershipSkills        card.LeadershipSkills `json:"leadershipSkills"`
	Legalities              card.CardLegalities   `json:"legalities"`
	Life                    string                `json:"life"`
	Loyalty                 string                `json:"loyalty"`
	ManaCost                string                `json:"manaCost"`
	ManaValue               int                   `json:"manaValue"`
	Name                    string                `json:"name"`
	Power                   string                `json:"power"`
	Printings               []string              `json:"printings"`
	PurchaseUrls            card.PurchaseUrls     `json:"purchaseUrls"`
	RelatedCards            card.RelatedCards     `json:"relatedCards"`
	Rulings                 []card.CardRulings    `json:"rulings"`
	Side                    string                `json:"side"`
	Subsets                 []string              `json:"subsets"`
	Subtypes                []string              `json:"subtypes"`
	Supertypes              []string              `json:"supertypes"`
	Text                    string                `json:"text"`
	Toughness               string                `json:"toughness"`
	Type                    string                `json:"type"`
	Types                   []string              `json:"types"`
	Artist                  string                `json:"artist"`
	ArtistIds               []string              `json:"artistIds"`
	Availability            []string              `json:"availability"`
	BoosterTypes            []string              `json:"boosterTypes"`
	BorderColor             string                `json:"borderColor"`
	CardParts               []string              `json:"cardParts"`
	DuelDeck                string                `json:"duelDeck"`
	FaceFlavorName          string                `json:"faceFlavorName"`
	Finishes                []string              `json:"finishes"`
	FlavorName              string                `json:"flavorName"`
	FrameEffects            []string              `json:"frameEffects"`
	FrameVersion            string                `json:"frameVersion"`
	HasContentWarning       bool                  `json:"hasContentWarning"`
	HasFoil                 bool                  `json:"hasFoil"`
	HasNonFoil              bool                  `json:"hasNonFoil"`
	IsAlternative           bool                  `json:"isAlternative"`
	IsFullArt               bool                  `json:"isFullArt"`
	IsOnlineOnly            bool                  `json:"isOnlineOnly"`
	IsOversized             bool                  `json:"isOversized"`
	IsPromo                 bool                  `json:"isPromo"`
	IsRebalanced            bool                  `json:"isRebalanced"`
	IsReprint               bool                  `json:"isReprint"`
	IsStarter               bool                  `json:"isStarter"`
	IsStorySpotlight        bool                  `json:"isStorySpotlight"`
	IsTextless              bool                  `json:"isTextless"`
	IsTimeshifted           bool                  `json:"isTimeshifted"`
	Language                string                `json:"language"`
	Number                  string                `json:"number"`
	OriginalPrintings       []string              `json:"originalPrintings"`
	OriginalReleaseDate     string                `json:"originalReleaseDate"`
	OriginalText            string                `json:"originalText"`
	OriginalType            string                `json:"originalType"`
	OtherFaceIds            []string              `json:"otherFaceIds"`
	PromoTypes              []string              `json:"promoTypes"`
	Rarity                  string                `json:"rarity"`
	RebalancedPrintings     []string              `json:"rebalancedPrintings"`
	SecurityStamp           string                `json:"securityStamp"`
	SetCode                 string                `json:"setCode"`
	Signature               string                `json:"signature"`
	SourceProducts          card.SourceProducts   `json:"sourceProducts"`
	UUID                    string                `json:"uuid"`
	Variations              []string              `json:"variations"`
	Watermark               []string              `json:"watermark"`
}
