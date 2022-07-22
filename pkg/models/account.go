package models

import (
	"hash"
	"time"
)

type Account struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Secret    hash.Hash `json:"secret"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}
