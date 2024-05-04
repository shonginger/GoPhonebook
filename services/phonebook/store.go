package phonebook

import (
	"database/sql"
	"fmt"

	"github.com/shonginger/GoPhonebook/Lehem/logger"
	"github.com/shonginger/GoPhonebook/Lehem/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetContactByPhoneNumber(phone string) (*types.Contact, error) {
	rows, err := s.db.Query("SELECT * FROM contacts WHERE phone = ?", phone)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	contact := new(types.Contact)
	for rows.Next() {
		contact, err = ScanRowIntoUser(rows)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		if contact.ID == 0 {
			return nil, fmt.Errorf("user not found")
		}
	}

	return contact, nil
}

func (s *Store) GetContactById(id int) (*types.Contact, error) {
	rows, err := s.db.Query("SELECT * FROM contacts WHERE ID = ?", id)
	if err != nil {
		return nil, err
	}

	contact := new(types.Contact)

	if !rows.Next() {
		return nil, fmt.Errorf("contact not found")
	} else {
		contact, err = ScanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	return contact, nil
}

func (s *Store) GetContactsPage(page int) ([]types.Contact, error) {
	pageSize := 10
	offset := (page - 1) * pageSize

	rows, err := s.db.Query("SELECT * FROM contacts LIMIT ? OFFSET ?", pageSize, offset)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var contacts []types.Contact
	for rows.Next() {
		var contact types.Contact
		if err := rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Address, &contact.CreatedAt); err != nil {
			// Handle error
			logger.Error(err)
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (s *Store) DeleteContactById(id int) error {
	_, err := s.db.Exec("DELETE FROM contacts WHERE ID = ?", id)

	// if no err we return nil anyway
	return err
}

func (s *Store) AddContact(contactData types.ContactDTO) error {
	_, err := s.db.Exec("INSERT INTO contacts (firstName, lastName, phone, address) VALUES (?,?,?,?)",
		contactData.FirstName, contactData.LastName.String, contactData.Phone, contactData.Address.String)

	// if no err we return nil anyway
	return err
}

func (s *Store) UpdateContact(id int, contactData types.UpdateContactDTO) error {
	// attempt to get the existing contact
	contact, err := s.GetContactById(id)

	if err != nil {
		return err
	}

	//  update the fields that were filled in the request
	if contactData.Phone.Valid {
		contact.Phone = contactData.Phone.String
	}

	if contactData.FirstName.Valid {
		contact.FirstName = contactData.FirstName.String
	}

	if contactData.LastName.Valid {
		contact.LastName = contactData.LastName.String
	}

	if contactData.Address.Valid {
		contact.Address = contactData.Address.String
	}

	_, err = s.db.Exec("UPDATE contacts SET firstName = ?, lastName = ?, phone = ?, address = ? WHERE ID = ?",
		contact.FirstName, contact.LastName, contact.Phone, contact.Address, id)

	// if no err we return nil anyway
	return err
}

func ScanRowIntoUser(rows *sql.Rows) (*types.Contact, error) {
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
