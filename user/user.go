package user

type User struct {
	Username      string   `json:"username"`
	Email         string   `json:"email"`
	LastLoginTime string   `json:"lastLoginTime"`
	Decks         []string `json:"decks"`
	Cards         []string `json:"cards"`
	Sets          []string `json:"sets"`
}
