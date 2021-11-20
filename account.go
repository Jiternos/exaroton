package exaroton

import "encoding/json"

type Account struct {
	// The user's username.
	Name string `json:"name"`

	// The user's email.
	Email string `json:"email"`

	// The user's credits.
	Credits int64 `json:"credits"`

	// Whether the user is verified.
	Verified bool `json:"verified"`
}

// Account returns the account information as struct
func (s *Session) Account() (account *Account, err error) {
	body, err := s.Request("GET", EndpointAccount, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &account)
	return
}
