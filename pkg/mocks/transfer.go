package mocks

import (
	"time"

	"github.com/mgbcaio/desafio-go-stone/pkg/models"
)

var Transfers = []models.Transfer{
	{
		Id:                   1,
		AccountOriginId:      1,
		AccountDestinationId: 2,
		Amount:               430.65,
		CreatedAt:            time.Now(),
	},
	{
		Id:                   2,
		AccountOriginId:      1,
		AccountDestinationId: 2,
		Amount:               150.78,
		CreatedAt:            time.Now(),
	},
}
