package mocks

import (
	"hash"
	"hash/fnv"
	"time"

	"github.com/mgbcaio/desafio-go-stone/pkg/models"
)

func hashing(s string) hash.Hash {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h
}

var Accounts = []models.Account{
	{
		Id:        1,
		Name:      "Account number 1",
		Cpf:       "00099988877",
		Secret:    hashing("test-secret-123"),
		Balance:   3400.5,
		CreatedAt: time.Now(),
	},
	{
		Id:        2,
		Name:      "Account number 2",
		Cpf:       "11122233344",
		Secret:    hashing("test-secret-456"),
		Balance:   143.5,
		CreatedAt: time.Now(),
	},
}
