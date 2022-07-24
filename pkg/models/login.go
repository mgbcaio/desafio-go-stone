package models

import "hash"

// Login represents the login object for a user.
type Login struct {
	Cpf    string    `json:"cpf"`
	Secret hash.Hash `json:"secret"`
}
