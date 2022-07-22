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

var Account = []models.Account{
	{
		Id:        1,
		Name:      "Account number 1",
		Cpf:       "07120867300",
		Secret:    hashing("test-secret-123"),
		Balance:   3400.5,
		CreatedAt: time.Now(),
	},
}
