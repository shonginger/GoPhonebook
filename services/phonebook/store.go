package phonebook

import (
	"database/sql"
	"fmt"

	"github.com/shonginger/GoPhonebook/Lehem/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetContactByPhoneNumber(phone string) (*types.Contact, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE phone = ?", phone)
	if err != nil {
		return nil, err
	}

	contact := new(types.Contact)
	for rows.Next() {
		contact, err = scnaRowIntoUser(rows)
		if err != nil {
			return nil, err
		}

		if contact.ID == 0 {
			return nil, fmt.Errorf("user not found")
		}
	}

	return contact, nil
}

func (s *Store) DeleteContactByPhoneNumber(phone string) (*types.Contact, error) {
	return nil, nil
}

func (s *Store) AddContact(contactData types.ContactDTO) error {
	return nil
}

func (s *Store) UpdateContact(contactData types.ContactDTO) (*types.Contact, error) {
	return nil, nil
}

func scnaRowIntoUser(rows *sql.Rows) (*types.Contact, error) {
	contact := new(types.Contact)

	err := rows.Scan(
		&contact.ID,
		&contact.FirstName,
		&contact.LastName,
		&contact.Phone,
		&contact.Address,
		&contact.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return contact, nil
}
