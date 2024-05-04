package types

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

type NullableString sql.NullString

type ContactStore interface {
	GetContactByPhoneNumber(phone string) (*Contact, error)
	GetContactById(id int) (*Contact, error)
	DeleteContactById(id int) error
	AddContact(contactData ContactDTO) error
	UpdateContact(id int, updateContactData UpdateContactDTO) error
	GetContactsPage(pageNum int) ([]Contact, error)
}

type ContactDTO struct {
	FirstName string         `json:"firstName" validate:"required"`
	LastName  NullableString `json:"lastName"`
	Phone     string         `json:"phone" validate:"required,min=7,max=20"`
	Address   NullableString `json:"address"`
}

type UpdateContactDTO struct {
	FirstName NullableString `json:"firstName"`
	LastName  NullableString `json:"lastName"`
	Phone     NullableString `json:"phone"`
	Address   NullableString `json:"address"`
}

func (dto *UpdateContactDTO) Validate() error {
	if dto.Phone.Valid && (len(dto.Phone.String) < 7 || len(dto.Phone.String) > 20) {
		return errors.New("phone length must be between 7 and 20 characters")
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (ns *NullableString) UnmarshalJSON(data []byte) error {
	// Check if the data is null
	if string(data) == "null" {
		ns.Valid = false
		return nil
	}

	// Unmarshal the data into a string
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	// Assign the value to the sql.NullString
	ns.String = value
	ns.Valid = true
	return nil
}

type Contact struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}
