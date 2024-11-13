package user

type User struct {
	Username              string   `json:"username"`
	Email                 string   `json:"email"`
	CredentialId          string   `json:"credentialId"`
	LastLoginTime         string   `json:"lastLoginTime"`
	LastPasswordResetTime string   `json:"lastPasswordResetTime"`
	LastKeyRotationTime   string   `json:"lastKeyRotationTime"`
	Decks                 []string `json:"decks"`
	Cards                 []string `json:"cards"`
	Sets                  []string `json:"sets"`
}
