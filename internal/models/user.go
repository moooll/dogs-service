package models

import "github.com/google/uuid"

// User describes user of the service
type User struct {
	// ID describes user's id
	ID uuid.UUID
	// Username describes user's username
	Username string `json:"username" govalid:"req|min:1|max:100|regex:^[a-zA-Z0-9]+$"`
	// Password describes user's password
	Password string `json:"password" govalid:"req|min:1|max:100"`
}
