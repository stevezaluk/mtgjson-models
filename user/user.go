package user

type User struct {
	Username     string   `json:"username"`
	Email        string   `json:"email"`
	CredentialId string   `json:"credentialId"`
	Decks        []string `json:"decks"`
	Cards        []string `json:"cards"`
	Sets         []string `json:"sets"`
}
