// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: card/card_set.proto

package card

import (
	meta "github.com/stevezaluk/mtgjson-models/meta"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CardSet - Represents a card within a Magic: The Gathering set
// See here for more info: https://mtgjson.com/data-models/card/card-set/
type CardSet struct {
	state                   protoimpl.MessageState `protogen:"open.v1"`
	AsciiName               string                 `protobuf:"bytes,1,opt,name=asciiName,proto3" json:"asciiName,omitempty" bson:"asciiName"`                               // @gotags: bson:"asciiName"
	AttractionLights        []string               `protobuf:"bytes,2,rep,name=attractionLights,proto3" json:"attractionLights,omitempty" bson:"attractionLights"`                 // @gotags: bson:"attractionLights"
	ColorIdentity           []string               `protobuf:"bytes,3,rep,name=colorIdentity,proto3" json:"colorIdentity,omitempty" bson:"colorIdentity"`                       // @gotags: bson:"colorIdentity"
	ColorIndicator          []string               `protobuf:"bytes,4,rep,name=colorIndicator,proto3" json:"colorIndicator,omitempty" bson:"colorIndicator"`                     // @gotags: bson:"colorIndicator"
	Colors                  []string               `protobuf:"bytes,5,rep,name=colors,proto3" json:"colors,omitempty" bson:"colors"`                                     // @gotags: bson:"colors"
	ConvertedManaCost       int64                  `protobuf:"varint,6,opt,name=convertedManaCost,proto3" json:"convertedManaCost,omitempty" bson:"convertedManaCost"`              // @gotags: bson:"convertedManaCost"
	Defense                 string                 `protobuf:"bytes,7,opt,name=defense,proto3" json:"defense,omitempty" bson:"defense"`                                   // @gotags: bson:"defense"
	EdhrecRank              int64                  `protobuf:"varint,8,opt,name=edhrecRank,proto3" json:"edhrecRank,omitempty" bson:"edhrecRank"`                            // @gotags: bson:"edhrecRank"
	EdhrecSaltiness         float64                `protobuf:"fixed64,9,opt,name=edhrecSaltiness,proto3" json:"edhrecSaltiness,omitempty" bson:"edhrecSaltiness"`                 // @gotags: bson:"edhrecSaltiness"
	FaceConvertedManaCost   int64                  `protobuf:"varint,10,opt,name=faceConvertedManaCost,proto3" json:"faceConvertedManaCost,omitempty" bson:"faceConvertedManaCost"`     // @gotags: bson:"faceConvertedManaCost"
	FaceManaValue           int64                  `protobuf:"varint,11,opt,name=faceManaValue,proto3" json:"faceManaValue,omitempty" bson:"faceManaValue"`                     // @gotags: bson:"faceManaValue"
	FaceName                string                 `protobuf:"bytes,12,opt,name=faceName,proto3" json:"faceName,omitempty" bson:"faceName"`                                // @gotags: bson:"faceName"
	FirstPrinting           string                 `protobuf:"bytes,13,opt,name=firstPrinting,proto3" json:"firstPrinting,omitempty" bson:"firstPrinting"`                      // @gotags: bson:"firstPrinting"
	ForeignData             []*meta.ForeignData    `protobuf:"bytes,14,rep,name=foreignData,proto3" json:"foreignData,omitempty" bson:"foreignData"`                          // @gotags: bson:"foreignData"
	Hand                    string                 `protobuf:"bytes,15,opt,name=hand,proto3" json:"hand,omitempty" bson:"hand"`                                        // @gotags: bson:"hand"
	HasAlternativeDeckLimit bool                   `protobuf:"varint,16,opt,name=hasAlternativeDeckLimit,proto3" json:"hasAlternativeDeckLimit,omitempty" bson:"hasAlternativeDeckLimit"` // @gotags: bson:"hasAlternativeDeckLimit"
	Identifiers             *meta.CardIdentifiers  `protobuf:"bytes,17,opt,name=identifiers,proto3" json:"identifiers,omitempty" bson:"identifiers"`                          // @gotags: bson:"identifiers"
	IsFunny                 bool                   `protobuf:"varint,18,opt,name=isFunny,proto3" json:"isFunny,omitempty" bson:"isFunny"`                                 // @gotags: bson:"isFunny"
	IsReserved              bool                   `protobuf:"varint,19,opt,name=isReserved,proto3" json:"isReserved,omitempty" bson:"isReserved"`                           // @gotags: bson:"isReserved"
	Keywords                []string               `protobuf:"bytes,20,rep,name=keywords,proto3" json:"keywords,omitempty" bson:"keywords"`                                // @gotags: bson:"keywords"
	Layout                  string                 `protobuf:"bytes,21,opt,name=layout,proto3" json:"layout,omitempty" bson:"layout"`                                    // @gotags: bson:"layout"
	LeadershipSkills        *meta.LeadershipSkills `protobuf:"bytes,22,opt,name=leadershipSkills,proto3" json:"leadershipSkills,omitempty" bson:"leadershipSkills"`                // @gotags: bson:"leadershipSkills"
	Legalities              *meta.CardLegalities   `protobuf:"bytes,23,opt,name=legalities,proto3" json:"legalities,omitempty" bson:"legalities"`                            // @gotags: bson:"legalities"
	Life                    string                 `protobuf:"bytes,24,opt,name=life,proto3" json:"life,omitempty" bson:"life"`                                        // @gotags: bson:"life"
	Loyalty                 string                 `protobuf:"bytes,25,opt,name=loyalty,proto3" json:"loyalty,omitempty" bson:"loyalty"`                                  // @gotags: bson:"loyalty"
	ManaCost                string                 `protobuf:"bytes,26,opt,name=manaCost,proto3" json:"manaCost,omitempty" bson:"manaCost"`                                // @gotags: bson:"manaCost"
	ManaValue               int64                  `protobuf:"varint,27,opt,name=manaValue,proto3" json:"manaValue,omitempty" bson:"manaValue"`                             // @gotags: bson:"manaValue"
	Name                    string                 `protobuf:"bytes,28,opt,name=name,proto3" json:"name,omitempty" bson:"name"`                                        // @gotags: bson:"name"
	Power                   string                 `protobuf:"bytes,29,opt,name=power,proto3" json:"power,omitempty" bson:"power"`                                      // @gotags: bson:"power"
	Printings               []string               `protobuf:"bytes,30,rep,name=printings,proto3" json:"printings,omitempty" bson:"printings"`                              // @gotags: bson:"printings"
	PurchaseUrls            *meta.PurchaseUrls     `protobuf:"bytes,31,opt,name=purchaseUrls,proto3" json:"purchaseUrls,omitempty" bson:"purchaseUrls"`                        // @gotags: bson:"purchaseUrls"
	RelatedCards            *meta.RelatedCards     `protobuf:"bytes,32,opt,name=relatedCards,proto3" json:"relatedCards,omitempty" bson:"relatedCards"`                        // @gotags: bson:"relatedCards"
	Rulings                 []*meta.CardRulings    `protobuf:"bytes,33,rep,name=rulings,proto3" json:"rulings,omitempty" bson:"rulings"`                                  // @gotags: bson:"rulings"
	Side                    string                 `protobuf:"bytes,34,opt,name=side,proto3" json:"side,omitempty" bson:"side"`                                        // @gotags: bson:"side"
	Subsets                 []string               `protobuf:"bytes,35,rep,name=subsets,proto3" json:"subsets,omitempty" bson:"subsets"`                                  // @gotags: bson:"subsets"
	Subtypes                []string               `protobuf:"bytes,36,rep,name=subtypes,proto3" json:"subtypes,omitempty" bson:"subtypes"`                                // @gotags: bson:"subtypes"
	Supertypes              []string               `protobuf:"bytes,37,rep,name=supertypes,proto3" json:"supertypes,omitempty" bson:"supertypes"`                            // @gotags: bson:"supertypes"
	Text                    string                 `protobuf:"bytes,38,opt,name=text,proto3" json:"text,omitempty" bson:"text"`                                        // @gotags: bson:"text"
	Toughness               string                 `protobuf:"bytes,39,opt,name=toughness,proto3" json:"toughness,omitempty" bson:"toughness"`                              // @gotags: bson:"toughness"
	Type                    string                 `protobuf:"bytes,40,opt,name=type,proto3" json:"type,omitempty" bson:"type"`                                        // @gotags: bson:"type"
	Types                   []string               `protobuf:"bytes,41,rep,name=types,proto3" json:"types,omitempty" bson:"types"`                                      // @gotags: bson:"types"
	Artist                  string                 `protobuf:"bytes,42,opt,name=artist,proto3" json:"artist,omitempty" bson:"artist"`                                    // @gotags: bson:"artist"
	ArtistIds               []string               `protobuf:"bytes,43,rep,name=artistIds,proto3" json:"artistIds,omitempty" bson:"artistIds"`                              // @gotags: bson:"artistIds"
	Availability            []string               `protobuf:"bytes,44,rep,name=availability,proto3" json:"availability,omitempty" bson:"availability"`                        // @gotags: bson:"availability"
	BoosterTypes            []string               `protobuf:"bytes,45,rep,name=boosterTypes,proto3" json:"boosterTypes,omitempty" bson:"boosterTypes"`                        // @gotags: bson:"boosterTypes"
	BorderColor             string                 `protobuf:"bytes,46,opt,name=borderColor,proto3" json:"borderColor,omitempty" bson:"borderColor"`                          // @gotags: bson:"borderColor"
	CardParts               []string               `protobuf:"bytes,47,rep,name=cardParts,proto3" json:"cardParts,omitempty" bson:"cardParts"`                              // @gotags: bson:"cardParts"
	DuelDeck                string                 `protobuf:"bytes,48,opt,name=duelDeck,proto3" json:"duelDeck,omitempty" bson:"duelDeck"`                                // @gotags: bson:"duelDeck"
	FaceFlavorName          string                 `protobuf:"bytes,49,opt,name=faceFlavorName,proto3" json:"faceFlavorName,omitempty" bson:"faceFlavorName"`                    // @gotags: bson:"faceFlavorName"
	Finishes                []string               `protobuf:"bytes,50,rep,name=finishes,proto3" json:"finishes,omitempty" bson:"finishes"`                                // @gotags: bson:"finishes"
	FlavorName              string                 `protobuf:"bytes,51,opt,name=flavorName,proto3" json:"flavorName,omitempty" bson:"flavorName"`                            // @gotags: bson:"flavorName"
	FrameEffects            []string               `protobuf:"bytes,52,rep,name=frameEffects,proto3" json:"frameEffects,omitempty" bson:"frameEffects"`                        // @gotags: bson:"frameEffects"
	FrameVersion            string                 `protobuf:"bytes,53,opt,name=frameVersion,proto3" json:"frameVersion,omitempty" bson:"frameVersion"`                        // @gotags: bson:"frameVersion"
	HasContentWarning       bool                   `protobuf:"varint,54,opt,name=hasContentWarning,proto3" json:"hasContentWarning,omitempty" bson:"hasContentWarning"`             // @gotags: bson:"hasContentWarning"
	HasFoil                 bool                   `protobuf:"varint,55,opt,name=hasFoil,proto3" json:"hasFoil,omitempty" bson:"hasFoil"`                                 // @gotags: bson:"hasFoil"
	HasNonFoil              bool                   `protobuf:"varint,56,opt,name=hasNonFoil,proto3" json:"hasNonFoil,omitempty" bson:"hasNonFoil"`                           // @gotags: bson:"hasNonFoil"
	IsAlternative           bool                   `protobuf:"varint,57,opt,name=isAlternative,proto3" json:"isAlternative,omitempty" bson:"isAlternative"`                     // @gotags: bson:"isAlternative"
	IsFullArt               bool                   `protobuf:"varint,58,opt,name=isFullArt,proto3" json:"isFullArt,omitempty" bson:"isFullArt"`                             // @gotags: bson:"isFullArt"
	IsOnlineOnly            bool                   `protobuf:"varint,59,opt,name=isOnlineOnly,proto3" json:"isOnlineOnly,omitempty" bson:"isOnlineOnly"`                       // @gotags: bson:"isOnlineOnly"
	IsOversized             bool                   `protobuf:"varint,60,opt,name=isOversized,proto3" json:"isOversized,omitempty" bson:"isOversized"`                         // @gotags: bson:"isOversized"
	IsPromo                 bool                   `protobuf:"varint,61,opt,name=isPromo,proto3" json:"isPromo,omitempty" bson:"isPromo"`                                 // @gotags: bson:"isPromo"
	IsRebalanced            bool                   `protobuf:"varint,62,opt,name=isRebalanced,proto3" json:"isRebalanced,omitempty" bson:"isRebalanced"`                       // @gotags: bson:"isRebalanced"
	IsReprint               bool                   `protobuf:"varint,63,opt,name=isReprint,proto3" json:"isReprint,omitempty" bson:"isReprint"`                             // @gotags: bson:"isReprint"
	IsStarter               bool                   `protobuf:"varint,64,opt,name=isStarter,proto3" json:"isStarter,omitempty" bson:"isStarter"`                             // @gotags: bson:"isStarter"
	IsStorySpotlight        bool                   `protobuf:"varint,65,opt,name=isStorySpotlight,proto3" json:"isStorySpotlight,omitempty" bson:"isStorySpotlight"`               // @gotags: bson:"isStorySpotlight"
	IsTextless              bool                   `protobuf:"varint,66,opt,name=isTextless,proto3" json:"isTextless,omitempty" bson:"isTextless"`                           // @gotags: bson:"isTextless"
	IsTimeshifted           bool                   `protobuf:"varint,67,opt,name=isTimeshifted,proto3" json:"isTimeshifted,omitempty" bson:"isTimeshifted"`                     // @gotags: bson:"isTimeshifted"
	Language                string                 `protobuf:"bytes,68,opt,name=language,proto3" json:"language,omitempty" bson:"language"`                                // @gotags: bson:"language"
	Number                  string                 `protobuf:"bytes,69,opt,name=number,proto3" json:"number,omitempty" bson:"number"`                                    // @gotags: bson:"number"
	OriginalPrintings       []string               `protobuf:"bytes,70,rep,name=originalPrintings,proto3" json:"originalPrintings,omitempty" bson:"originalPrintings"`              // @gotags: bson:"originalPrintings"
	OriginalReleaseDate     string                 `protobuf:"bytes,71,opt,name=originalReleaseDate,proto3" json:"originalReleaseDate,omitempty" bson:"originalReleaseDate"`          // @gotags: bson:"originalReleaseDate"
	OriginalText            string                 `protobuf:"bytes,72,opt,name=originalText,proto3" json:"originalText,omitempty" bson:"originalText"`                        // @gotags: bson:"originalText"
	OriginalType            string                 `protobuf:"bytes,73,opt,name=originalType,proto3" json:"originalType,omitempty" bson:"originalType"`                        // @gotags: bson:"originalType"
	OtherFaceIds            []string               `protobuf:"bytes,74,rep,name=otherFaceIds,proto3" json:"otherFaceIds,omitempty" bson:"otherFaceIds"`                        // @gotags: bson:"otherFaceIds"
	PromoTypes              []string               `protobuf:"bytes,75,rep,name=promoTypes,proto3" json:"promoTypes,omitempty" bson:"promoTypes"`                            // @gotags: bson:"promoTypes"
	Rarity                  string                 `protobuf:"bytes,76,opt,name=rarity,proto3" json:"rarity,omitempty" bson:"rarity"`                                    // @gotags: bson:"rarity"
	RebalancedPrintings     []string               `protobuf:"bytes,77,rep,name=rebalancedPrintings,proto3" json:"rebalancedPrintings,omitempty" bson:"rebalancedPrintings"`          // @gotags: bson:"rebalancedPrintings"
	SecurityStamp           string                 `protobuf:"bytes,78,opt,name=securityStamp,proto3" json:"securityStamp,omitempty" bson:"securityStamp"`                      // @gotags: bson:"securityStamp"
	SetCode                 string                 `protobuf:"bytes,79,opt,name=setCode,proto3" json:"setCode,omitempty" bson:"setCode"`                                  // @gotags: bson:"setCode"
	Signature               string                 `protobuf:"bytes,80,opt,name=signature,proto3" json:"signature,omitempty" bson:"signature"`                              // @gotags: bson:"signature"
	SourceProducts          *meta.SourceProducts   `protobuf:"bytes,81,opt,name=sourceProducts,proto3" json:"sourceProducts,omitempty" bson:"sourceProducts"`                    // @gotags: bson:"sourceProducts"
	Uuid                    string                 `protobuf:"bytes,82,opt,name=uuid,proto3" json:"uuid,omitempty" bson:"uuid"`                                        // @gotags: bson:"uuid"
	Variations              []string               `protobuf:"bytes,83,rep,name=variations,proto3" json:"variations,omitempty" bson:"variations"`                            // @gotags: bson:"variations"
	Watermark               []string               `protobuf:"bytes,84,rep,name=watermark,proto3" json:"watermark,omitempty" bson:"watermark"`                              // @gotags: bson:"watermark"
	MtgjsonApiMeta          *meta.MTGJSONAPIMeta   `protobuf:"bytes,85,opt,name=mtgjsonApiMeta,proto3" json:"mtgjsonApiMeta,omitempty" bson:"mtgjsonApiMeta"`                    // @gotags: bson:"mtgjsonApiMeta"
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *CardSet) Reset() {
	*x = CardSet{}
	mi := &file_card_card_set_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CardSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardSet) ProtoMessage() {}

func (x *CardSet) ProtoReflect() protoreflect.Message {
	mi := &file_card_card_set_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardSet.ProtoReflect.Descriptor instead.
func (*CardSet) Descriptor() ([]byte, []int) {
	return file_card_card_set_proto_rawDescGZIP(), []int{0}
}

func (x *CardSet) GetAsciiName() string {
	if x != nil {
		return x.AsciiName
	}
	return ""
}

func (x *CardSet) GetAttractionLights() []string {
	if x != nil {
		return x.AttractionLights
	}
	return nil
}

func (x *CardSet) GetColorIdentity() []string {
	if x != nil {
		return x.ColorIdentity
	}
	return nil
}

func (x *CardSet) GetColorIndicator() []string {
	if x != nil {
		return x.ColorIndicator
	}
	return nil
}

func (x *CardSet) GetColors() []string {
	if x != nil {
		return x.Colors
	}
	return nil
}

func (x *CardSet) GetConvertedManaCost() int64 {
	if x != nil {
		return x.ConvertedManaCost
	}
	return 0
}

func (x *CardSet) GetDefense() string {
	if x != nil {
		return x.Defense
	}
	return ""
}

func (x *CardSet) GetEdhrecRank() int64 {
	if x != nil {
		return x.EdhrecRank
	}
	return 0
}

func (x *CardSet) GetEdhrecSaltiness() float64 {
	if x != nil {
		return x.EdhrecSaltiness
	}
	return 0
}

func (x *CardSet) GetFaceConvertedManaCost() int64 {
	if x != nil {
		return x.FaceConvertedManaCost
	}
	return 0
}

func (x *CardSet) GetFaceManaValue() int64 {
	if x != nil {
		return x.FaceManaValue
	}
	return 0
}

func (x *CardSet) GetFaceName() string {
	if x != nil {
		return x.FaceName
	}
	return ""
}

func (x *CardSet) GetFirstPrinting() string {
	if x != nil {
		return x.FirstPrinting
	}
	return ""
}

func (x *CardSet) GetForeignData() []*meta.ForeignData {
	if x != nil {
		return x.ForeignData
	}
	return nil
}

func (x *CardSet) GetHand() string {
	if x != nil {
		return x.Hand
	}
	return ""
}

func (x *CardSet) GetHasAlternativeDeckLimit() bool {
	if x != nil {
		return x.HasAlternativeDeckLimit
	}
	return false
}

func (x *CardSet) GetIdentifiers() *meta.CardIdentifiers {
	if x != nil {
		return x.Identifiers
	}
	return nil
}

func (x *CardSet) GetIsFunny() bool {
	if x != nil {
		return x.IsFunny
	}
	return false
}

func (x *CardSet) GetIsReserved() bool {
	if x != nil {
		return x.IsReserved
	}
	return false
}

func (x *CardSet) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *CardSet) GetLayout() string {
	if x != nil {
		return x.Layout
	}
	return ""
}

func (x *CardSet) GetLeadershipSkills() *meta.LeadershipSkills {
	if x != nil {
		return x.LeadershipSkills
	}
	return nil
}

func (x *CardSet) GetLegalities() *meta.CardLegalities {
	if x != nil {
		return x.Legalities
	}
	return nil
}

func (x *CardSet) GetLife() string {
	if x != nil {
		return x.Life
	}
	return ""
}

func (x *CardSet) GetLoyalty() string {
	if x != nil {
		return x.Loyalty
	}
	return ""
}

func (x *CardSet) GetManaCost() string {
	if x != nil {
		return x.ManaCost
	}
	return ""
}

func (x *CardSet) GetManaValue() int64 {
	if x != nil {
		return x.ManaValue
	}
	return 0
}

func (x *CardSet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CardSet) GetPower() string {
	if x != nil {
		return x.Power
	}
	return ""
}

func (x *CardSet) GetPrintings() []string {
	if x != nil {
		return x.Printings
	}
	return nil
}

func (x *CardSet) GetPurchaseUrls() *meta.PurchaseUrls {
	if x != nil {
		return x.PurchaseUrls
	}
	return nil
}

func (x *CardSet) GetRelatedCards() *meta.RelatedCards {
	if x != nil {
		return x.RelatedCards
	}
	return nil
}

func (x *CardSet) GetRulings() []*meta.CardRulings {
	if x != nil {
		return x.Rulings
	}
	return nil
}

func (x *CardSet) GetSide() string {
	if x != nil {
		return x.Side
	}
	return ""
}

func (x *CardSet) GetSubsets() []string {
	if x != nil {
		return x.Subsets
	}
	return nil
}

func (x *CardSet) GetSubtypes() []string {
	if x != nil {
		return x.Subtypes
	}
	return nil
}

func (x *CardSet) GetSupertypes() []string {
	if x != nil {
		return x.Supertypes
	}
	return nil
}

func (x *CardSet) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CardSet) GetToughness() string {
	if x != nil {
		return x.Toughness
	}
	return ""
}

func (x *CardSet) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CardSet) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *CardSet) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *CardSet) GetArtistIds() []string {
	if x != nil {
		return x.ArtistIds
	}
	return nil
}

func (x *CardSet) GetAvailability() []string {
	if x != nil {
		return x.Availability
	}
	return nil
}

func (x *CardSet) GetBoosterTypes() []string {
	if x != nil {
		return x.BoosterTypes
	}
	return nil
}

func (x *CardSet) GetBorderColor() string {
	if x != nil {
		return x.BorderColor
	}
	return ""
}

func (x *CardSet) GetCardParts() []string {
	if x != nil {
		return x.CardParts
	}
	return nil
}

func (x *CardSet) GetDuelDeck() string {
	if x != nil {
		return x.DuelDeck
	}
	return ""
}

func (x *CardSet) GetFaceFlavorName() string {
	if x != nil {
		return x.FaceFlavorName
	}
	return ""
}

func (x *CardSet) GetFinishes() []string {
	if x != nil {
		return x.Finishes
	}
	return nil
}

func (x *CardSet) GetFlavorName() string {
	if x != nil {
		return x.FlavorName
	}
	return ""
}

func (x *CardSet) GetFrameEffects() []string {
	if x != nil {
		return x.FrameEffects
	}
	return nil
}

func (x *CardSet) GetFrameVersion() string {
	if x != nil {
		return x.FrameVersion
	}
	return ""
}

func (x *CardSet) GetHasContentWarning() bool {
	if x != nil {
		return x.HasContentWarning
	}
	return false
}

func (x *CardSet) GetHasFoil() bool {
	if x != nil {
		return x.HasFoil
	}
	return false
}

func (x *CardSet) GetHasNonFoil() bool {
	if x != nil {
		return x.HasNonFoil
	}
	return false
}

func (x *CardSet) GetIsAlternative() bool {
	if x != nil {
		return x.IsAlternative
	}
	return false
}

func (x *CardSet) GetIsFullArt() bool {
	if x != nil {
		return x.IsFullArt
	}
	return false
}

func (x *CardSet) GetIsOnlineOnly() bool {
	if x != nil {
		return x.IsOnlineOnly
	}
	return false
}

func (x *CardSet) GetIsOversized() bool {
	if x != nil {
		return x.IsOversized
	}
	return false
}

func (x *CardSet) GetIsPromo() bool {
	if x != nil {
		return x.IsPromo
	}
	return false
}

func (x *CardSet) GetIsRebalanced() bool {
	if x != nil {
		return x.IsRebalanced
	}
	return false
}

func (x *CardSet) GetIsReprint() bool {
	if x != nil {
		return x.IsReprint
	}
	return false
}

func (x *CardSet) GetIsStarter() bool {
	if x != nil {
		return x.IsStarter
	}
	return false
}

func (x *CardSet) GetIsStorySpotlight() bool {
	if x != nil {
		return x.IsStorySpotlight
	}
	return false
}

func (x *CardSet) GetIsTextless() bool {
	if x != nil {
		return x.IsTextless
	}
	return false
}

func (x *CardSet) GetIsTimeshifted() bool {
	if x != nil {
		return x.IsTimeshifted
	}
	return false
}

func (x *CardSet) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *CardSet) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *CardSet) GetOriginalPrintings() []string {
	if x != nil {
		return x.OriginalPrintings
	}
	return nil
}

func (x *CardSet) GetOriginalReleaseDate() string {
	if x != nil {
		return x.OriginalReleaseDate
	}
	return ""
}

func (x *CardSet) GetOriginalText() string {
	if x != nil {
		return x.OriginalText
	}
	return ""
}

func (x *CardSet) GetOriginalType() string {
	if x != nil {
		return x.OriginalType
	}
	return ""
}

func (x *CardSet) GetOtherFaceIds() []string {
	if x != nil {
		return x.OtherFaceIds
	}
	return nil
}

func (x *CardSet) GetPromoTypes() []string {
	if x != nil {
		return x.PromoTypes
	}
	return nil
}

func (x *CardSet) GetRarity() string {
	if x != nil {
		return x.Rarity
	}
	return ""
}

func (x *CardSet) GetRebalancedPrintings() []string {
	if x != nil {
		return x.RebalancedPrintings
	}
	return nil
}

func (x *CardSet) GetSecurityStamp() string {
	if x != nil {
		return x.SecurityStamp
	}
	return ""
}

func (x *CardSet) GetSetCode() string {
	if x != nil {
		return x.SetCode
	}
	return ""
}

func (x *CardSet) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *CardSet) GetSourceProducts() *meta.SourceProducts {
	if x != nil {
		return x.SourceProducts
	}
	return nil
}

func (x *CardSet) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CardSet) GetVariations() []string {
	if x != nil {
		return x.Variations
	}
	return nil
}

func (x *CardSet) GetWatermark() []string {
	if x != nil {
		return x.Watermark
	}
	return nil
}

func (x *CardSet) GetMtgjsonApiMeta() *meta.MTGJSONAPIMeta {
	if x != nil {
		return x.MtgjsonApiMeta
	}
	return nil
}

var File_card_card_set_proto protoreflect.FileDescriptor

const file_card_card_set_proto_rawDesc = "" +
	"\n" +
	"\x13card/card_set.proto\x12\x04card\x1a\x16meta/identifiers.proto\x1a\x15meta/legalities.proto\x1a\x18meta/related_cards.proto\x1a\x18meta/purchase_urls.proto\x1a\x17meta/card_rulings.proto\x1a\x15meta/leadership.proto\x1a\x1ameta/source_products.proto\x1a\x17meta/foreign_data.proto\x1a\x16meta/mtgjson_api.proto\"\x9a\x17\n" +
	"\aCardSet\x12\x1c\n" +
	"\tasciiName\x18\x01 \x01(\tR\tasciiName\x12*\n" +
	"\x10attractionLights\x18\x02 \x03(\tR\x10attractionLights\x12$\n" +
	"\rcolorIdentity\x18\x03 \x03(\tR\rcolorIdentity\x12&\n" +
	"\x0ecolorIndicator\x18\x04 \x03(\tR\x0ecolorIndicator\x12\x16\n" +
	"\x06colors\x18\x05 \x03(\tR\x06colors\x12,\n" +
	"\x11convertedManaCost\x18\x06 \x01(\x03R\x11convertedManaCost\x12\x18\n" +
	"\adefense\x18\a \x01(\tR\adefense\x12\x1e\n" +
	"\n" +
	"edhrecRank\x18\b \x01(\x03R\n" +
	"edhrecRank\x12(\n" +
	"\x0fedhrecSaltiness\x18\t \x01(\x01R\x0fedhrecSaltiness\x124\n" +
	"\x15faceConvertedManaCost\x18\n" +
	" \x01(\x03R\x15faceConvertedManaCost\x12$\n" +
	"\rfaceManaValue\x18\v \x01(\x03R\rfaceManaValue\x12\x1a\n" +
	"\bfaceName\x18\f \x01(\tR\bfaceName\x12$\n" +
	"\rfirstPrinting\x18\r \x01(\tR\rfirstPrinting\x123\n" +
	"\vforeignData\x18\x0e \x03(\v2\x11.meta.ForeignDataR\vforeignData\x12\x12\n" +
	"\x04hand\x18\x0f \x01(\tR\x04hand\x128\n" +
	"\x17hasAlternativeDeckLimit\x18\x10 \x01(\bR\x17hasAlternativeDeckLimit\x127\n" +
	"\videntifiers\x18\x11 \x01(\v2\x15.meta.CardIdentifiersR\videntifiers\x12\x18\n" +
	"\aisFunny\x18\x12 \x01(\bR\aisFunny\x12\x1e\n" +
	"\n" +
	"isReserved\x18\x13 \x01(\bR\n" +
	"isReserved\x12\x1a\n" +
	"\bkeywords\x18\x14 \x03(\tR\bkeywords\x12\x16\n" +
	"\x06layout\x18\x15 \x01(\tR\x06layout\x12B\n" +
	"\x10leadershipSkills\x18\x16 \x01(\v2\x16.meta.LeadershipSkillsR\x10leadershipSkills\x124\n" +
	"\n" +
	"legalities\x18\x17 \x01(\v2\x14.meta.CardLegalitiesR\n" +
	"legalities\x12\x12\n" +
	"\x04life\x18\x18 \x01(\tR\x04life\x12\x18\n" +
	"\aloyalty\x18\x19 \x01(\tR\aloyalty\x12\x1a\n" +
	"\bmanaCost\x18\x1a \x01(\tR\bmanaCost\x12\x1c\n" +
	"\tmanaValue\x18\x1b \x01(\x03R\tmanaValue\x12\x12\n" +
	"\x04name\x18\x1c \x01(\tR\x04name\x12\x14\n" +
	"\x05power\x18\x1d \x01(\tR\x05power\x12\x1c\n" +
	"\tprintings\x18\x1e \x03(\tR\tprintings\x126\n" +
	"\fpurchaseUrls\x18\x1f \x01(\v2\x12.meta.PurchaseUrlsR\fpurchaseUrls\x126\n" +
	"\frelatedCards\x18  \x01(\v2\x12.meta.RelatedCardsR\frelatedCards\x12+\n" +
	"\arulings\x18! \x03(\v2\x11.meta.CardRulingsR\arulings\x12\x12\n" +
	"\x04side\x18\" \x01(\tR\x04side\x12\x18\n" +
	"\asubsets\x18# \x03(\tR\asubsets\x12\x1a\n" +
	"\bsubtypes\x18$ \x03(\tR\bsubtypes\x12\x1e\n" +
	"\n" +
	"supertypes\x18% \x03(\tR\n" +
	"supertypes\x12\x12\n" +
	"\x04text\x18& \x01(\tR\x04text\x12\x1c\n" +
	"\ttoughness\x18' \x01(\tR\ttoughness\x12\x12\n" +
	"\x04type\x18( \x01(\tR\x04type\x12\x14\n" +
	"\x05types\x18) \x03(\tR\x05types\x12\x16\n" +
	"\x06artist\x18* \x01(\tR\x06artist\x12\x1c\n" +
	"\tartistIds\x18+ \x03(\tR\tartistIds\x12\"\n" +
	"\favailability\x18, \x03(\tR\favailability\x12\"\n" +
	"\fboosterTypes\x18- \x03(\tR\fboosterTypes\x12 \n" +
	"\vborderColor\x18. \x01(\tR\vborderColor\x12\x1c\n" +
	"\tcardParts\x18/ \x03(\tR\tcardParts\x12\x1a\n" +
	"\bduelDeck\x180 \x01(\tR\bduelDeck\x12&\n" +
	"\x0efaceFlavorName\x181 \x01(\tR\x0efaceFlavorName\x12\x1a\n" +
	"\bfinishes\x182 \x03(\tR\bfinishes\x12\x1e\n" +
	"\n" +
	"flavorName\x183 \x01(\tR\n" +
	"flavorName\x12\"\n" +
	"\fframeEffects\x184 \x03(\tR\fframeEffects\x12\"\n" +
	"\fframeVersion\x185 \x01(\tR\fframeVersion\x12,\n" +
	"\x11hasContentWarning\x186 \x01(\bR\x11hasContentWarning\x12\x18\n" +
	"\ahasFoil\x187 \x01(\bR\ahasFoil\x12\x1e\n" +
	"\n" +
	"hasNonFoil\x188 \x01(\bR\n" +
	"hasNonFoil\x12$\n" +
	"\risAlternative\x189 \x01(\bR\risAlternative\x12\x1c\n" +
	"\tisFullArt\x18: \x01(\bR\tisFullArt\x12\"\n" +
	"\fisOnlineOnly\x18; \x01(\bR\fisOnlineOnly\x12 \n" +
	"\visOversized\x18< \x01(\bR\visOversized\x12\x18\n" +
	"\aisPromo\x18= \x01(\bR\aisPromo\x12\"\n" +
	"\fisRebalanced\x18> \x01(\bR\fisRebalanced\x12\x1c\n" +
	"\tisReprint\x18? \x01(\bR\tisReprint\x12\x1c\n" +
	"\tisStarter\x18@ \x01(\bR\tisStarter\x12*\n" +
	"\x10isStorySpotlight\x18A \x01(\bR\x10isStorySpotlight\x12\x1e\n" +
	"\n" +
	"isTextless\x18B \x01(\bR\n" +
	"isTextless\x12$\n" +
	"\risTimeshifted\x18C \x01(\bR\risTimeshifted\x12\x1a\n" +
	"\blanguage\x18D \x01(\tR\blanguage\x12\x16\n" +
	"\x06number\x18E \x01(\tR\x06number\x12,\n" +
	"\x11originalPrintings\x18F \x03(\tR\x11originalPrintings\x120\n" +
	"\x13originalReleaseDate\x18G \x01(\tR\x13originalReleaseDate\x12\"\n" +
	"\foriginalText\x18H \x01(\tR\foriginalText\x12\"\n" +
	"\foriginalType\x18I \x01(\tR\foriginalType\x12\"\n" +
	"\fotherFaceIds\x18J \x03(\tR\fotherFaceIds\x12\x1e\n" +
	"\n" +
	"promoTypes\x18K \x03(\tR\n" +
	"promoTypes\x12\x16\n" +
	"\x06rarity\x18L \x01(\tR\x06rarity\x120\n" +
	"\x13rebalancedPrintings\x18M \x03(\tR\x13rebalancedPrintings\x12$\n" +
	"\rsecurityStamp\x18N \x01(\tR\rsecurityStamp\x12\x18\n" +
	"\asetCode\x18O \x01(\tR\asetCode\x12\x1c\n" +
	"\tsignature\x18P \x01(\tR\tsignature\x12<\n" +
	"\x0esourceProducts\x18Q \x01(\v2\x14.meta.SourceProductsR\x0esourceProducts\x12\x12\n" +
	"\x04uuid\x18R \x01(\tR\x04uuid\x12\x1e\n" +
	"\n" +
	"variations\x18S \x03(\tR\n" +
	"variations\x12\x1c\n" +
	"\twatermark\x18T \x03(\tR\twatermark\x12<\n" +
	"\x0emtgjsonApiMeta\x18U \x01(\v2\x14.meta.MTGJSONAPIMetaR\x0emtgjsonApiMetaB+Z)github.com/stevezaluk/mtgjson-models/cardb\x06proto3"

var (
	file_card_card_set_proto_rawDescOnce sync.Once
	file_card_card_set_proto_rawDescData []byte
)

func file_card_card_set_proto_rawDescGZIP() []byte {
	file_card_card_set_proto_rawDescOnce.Do(func() {
		file_card_card_set_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_card_card_set_proto_rawDesc), len(file_card_card_set_proto_rawDesc)))
	})
	return file_card_card_set_proto_rawDescData
}

var file_card_card_set_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_card_card_set_proto_goTypes = []any{
	(*CardSet)(nil),               // 0: card.CardSet
	(*meta.ForeignData)(nil),      // 1: meta.ForeignData
	(*meta.CardIdentifiers)(nil),  // 2: meta.CardIdentifiers
	(*meta.LeadershipSkills)(nil), // 3: meta.LeadershipSkills
	(*meta.CardLegalities)(nil),   // 4: meta.CardLegalities
	(*meta.PurchaseUrls)(nil),     // 5: meta.PurchaseUrls
	(*meta.RelatedCards)(nil),     // 6: meta.RelatedCards
	(*meta.CardRulings)(nil),      // 7: meta.CardRulings
	(*meta.SourceProducts)(nil),   // 8: meta.SourceProducts
	(*meta.MTGJSONAPIMeta)(nil),   // 9: meta.MTGJSONAPIMeta
}
var file_card_card_set_proto_depIdxs = []int32{
	1, // 0: card.CardSet.foreignData:type_name -> meta.ForeignData
	2, // 1: card.CardSet.identifiers:type_name -> meta.CardIdentifiers
	3, // 2: card.CardSet.leadershipSkills:type_name -> meta.LeadershipSkills
	4, // 3: card.CardSet.legalities:type_name -> meta.CardLegalities
	5, // 4: card.CardSet.purchaseUrls:type_name -> meta.PurchaseUrls
	6, // 5: card.CardSet.relatedCards:type_name -> meta.RelatedCards
	7, // 6: card.CardSet.rulings:type_name -> meta.CardRulings
	8, // 7: card.CardSet.sourceProducts:type_name -> meta.SourceProducts
	9, // 8: card.CardSet.mtgjsonApiMeta:type_name -> meta.MTGJSONAPIMeta
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_card_card_set_proto_init() }
func file_card_card_set_proto_init() {
	if File_card_card_set_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_card_card_set_proto_rawDesc), len(file_card_card_set_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_card_card_set_proto_goTypes,
		DependencyIndexes: file_card_card_set_proto_depIdxs,
		MessageInfos:      file_card_card_set_proto_msgTypes,
	}.Build()
	File_card_card_set_proto = out.File
	file_card_card_set_proto_goTypes = nil
	file_card_card_set_proto_depIdxs = nil
}
