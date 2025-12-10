package models

type Account struct {
	Email             string
	PasswordHash      string
	VerificationToken string
}
