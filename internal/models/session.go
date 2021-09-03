package models

import (
	"time"

	"github.com/google/uuid"
)

// Session describes resfresh session
type Session struct {
	// ID describes session id
	ID uuid.UUID `json:"id"`
	// USerID describes id of the user logged in
	UserID uuid.UUID `json:"user_id"`
	// Fingerprint is the unique uuid describing the client app
	Fingerprint uuid.UUID `json:"fingerprint"`
	// RefreshToken is the token for refreshing
	RefreshToken string `json:"refresh_token"`
	// ExpiresAt is the time when the token becomes invalid
	ExpiresAt time.Time `json:"expires_at"`
}
