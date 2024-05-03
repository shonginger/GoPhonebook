package types

import "time"

type ContactStore interface {
	GetContactByPhoneNumber(phone string) (*Contact, error)
	DeleteContactByPhoneNumber(phone string) (*Contact, error)
	AddContact(contactData ContactDTO) error
	UpdateContact(contactData ContactDTO) (*Contact, error)
}

type ContactDTO struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone" validate:"required,min=7,max=20"`
	Address   string `json:"address"`
}

type Contact struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}
