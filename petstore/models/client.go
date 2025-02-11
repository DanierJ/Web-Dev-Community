package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Client is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Client struct {
    ID uuid.UUID `json:"id" db:"id"`
    Name string `json:"name" db:"name"`
    LastName string `json:"last_name" db:"last_name"`
    Email string `json:"email" db:"email"`
    Phone string `json:"phone" db:"phone"`
    Address string `json:"address" db:"address"`
    Gender string `json:"gender" db:"gender"`
    Age int `json:"age" db:"age"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (c Client) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Clients is not required by pop and may be deleted
type Clients []Client

// String is not required by pop and may be deleted
func (c Clients) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Client) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Name, Name: "Name"},
		&validators.StringIsPresent{Field: c.LastName, Name: "LastName"},
		&validators.StringIsPresent{Field: c.Email, Name: "Email"},
		&validators.StringIsPresent{Field: c.Phone, Name: "Phone"},
		&validators.StringIsPresent{Field: c.Address, Name: "Address"},
		&validators.StringIsPresent{Field: c.Gender, Name: "Gender"},
		&validators.IntIsPresent{Field: c.Age, Name: "Age"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Client) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Client) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
