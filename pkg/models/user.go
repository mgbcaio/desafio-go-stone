package models

// User represents a user object
type User struct {
	Id        int64  `json:"id"`
	Cpf       string `json:"cpf"`
	Secret    string `json:"secret"`
	AccountID int64  `json:"account_id"`
}
