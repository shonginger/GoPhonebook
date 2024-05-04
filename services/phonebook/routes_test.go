package phonebook

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/shonginger/GoPhonebook/Lehem/types"
)

type mockContactStore struct{}

func (m *mockContactStore) GetContactById(id int) (*types.Contact, error) {
	return nil, nil
}

func (m *mockContactStore) GetContactByPhoneNumber(phone string) (*types.Contact, error) {
	return nil, nil
}

func (m *mockContactStore) DeleteContactById(id int) error {
	return nil
}

func (m *mockContactStore) AddContact(contactData types.ContactDTO) error {
	return nil
}
func (m *mockContactStore) UpdateContact(id int, updateContactData types.UpdateContactDTO) error {
	return nil
}

func (m *mockContactStore) GetContactsPage(pageNum int) ([]types.Contact, error) {
	return nil, nil
}

func TestPhonebookServicerHandlers(t *testing.T) {
	contactStore := &mockContactStore{}
	handler := NewHandler(contactStore)

	t.Run("Should fail if user request is invalid", func(t *testing.T) {
		payload := types.ContactDTO{
			FirstName: "",
			LastName:  types.NullableString{"shady", true},
			Phone:     "",
			Address:   types.NullableString{"hollywood", true},
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/add", handler.HandleAddContact)
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusBadRequest {
			t.Errorf("expected status code was %d, actual was %d", http.StatusBadRequest, recorder.Code)
		}
	})

	t.Run("Should fail if user request body is empty", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodPost, "/add", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/add", handler.HandleAddContact)
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusBadRequest {
			t.Errorf("expected status code was %d, actual was %d", http.StatusBadRequest, recorder.Code)
		}
	})
}
