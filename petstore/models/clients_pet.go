package models

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

// ClientsPet is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type ClientsPet struct {
	ID        uuid.UUID `json:"id" db:"id"`
	ClientID  uuid.UUID `json:"client_id" db:"client_id"`
	PetID     uuid.UUID `json:"pet_id" db:"pet_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (c ClientsPet) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// ClientsPets is not required by pop and may be deleted
type ClientsPets []ClientsPet

// String is not required by pop and may be deleted
func (c ClientsPets) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}
