package models

import "time"

type Transfer struct {
	Id                   int64     `json:"id"`
	AccountOriginId      int64     `json:"account_origin_id"`
	AccountDestinationId int64     `json:"account_destination_id"`
	Amount               int64     `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}
