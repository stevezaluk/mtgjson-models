package user

type UserCredentials struct {
	CredentialId      string `json:"credentialId"`
	AESPrivateKey     string `json:"aesPrivateKey"`
	EncryptedPassword string `json:"encryptedPassword"`
}
