package model

import (
	"strings"
)

type Account struct {
	AccountId string            `json:"accountId,omitempty"`
	Email     string            `json:"email,omitempty"`
	Password  string            `json:"password,omitempty"`
	Token     string            `json:"token,omitempty"`
	Errors    map[string]string `json:"errors,omitempty"`
}

func (account *Account) Validate() bool {
	account.Errors = make(map[string]string)
	if strings.TrimSpace(account.Email) == "" {
		account.Errors["Email"] = "Please enter a valid email"
	}
	if strings.TrimSpace(account.Password) == "" {
		account.Errors["Password"] = "Please enter a password"
	}
	return len(account.Errors) == 0
}
