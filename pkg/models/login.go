package models

import "hash"

type Login struct {
	Cpf    string    `json:"cpf"`
	Secret hash.Hash `json:"secret"`
}
