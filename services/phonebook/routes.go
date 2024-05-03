package phonebook

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/shonginger/GoPhonebook/Lehem/types"
	"github.com/shonginger/GoPhonebook/Lehem/utils"
)

type Handler struct {
	store types.ContactStore
}

func NewHandler(store types.ContactStore) *Handler {
	return &Handler{store: store}
}

func (handler *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	}).Methods("GET")

	router.HandleFunc("/add", handler.HandleAddContact)
}

func (handler *Handler) HandleAddContact(writer http.ResponseWriter, request *http.Request) {
	// get payload
	var requestData types.ContactDTO
	if err := utils.ParseJSONFromHttpRequest(request, requestData); err != nil {
		utils.WriteError(writer, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(requestData); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(writer, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// check if it already exists
	_, err := handler.store.GetContactByPhoneNumber(requestData.Phone)

	// contact exists
	if err == nil {
		utils.WriteError(writer, http.StatusBadRequest, fmt.Errorf("user with phone %s already exists", requestData.Phone))
		return
	}

	// if it doesn't we create the new contact
	if err = handler.store.AddContact(requestData); err != nil {
		utils.WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSONResponse(writer, http.StatusCreated, nil)

}
