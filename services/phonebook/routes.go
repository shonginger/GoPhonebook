package phonebook

import (
	"fmt"
	"net/http"
	"strconv"

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
	router.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		utils.WriteJSONResponse(writer, http.StatusCreated, "pong")
	}).Methods("GET")
	router.HandleFunc("/add", handler.HandleAddContact).Methods("POST")
	router.HandleFunc("/update", handler.HandleUpdateContact).Methods("PUT")
	router.HandleFunc("/delete", handler.HandleDeleteContact).Methods("DELETE")
	router.HandleFunc("/search", handler.HandleSearchContact).Methods("GET")
	router.HandleFunc("/getContacts", handler.HandleGetContacts).Methods("GET")

}

func (handler *Handler) HandleAddContact(writer http.ResponseWriter, request *http.Request) {
	var requestData types.ContactDTO
	if err := ParseAndValidateRequest(request, &requestData); err != nil {
		utils.WriteError(writer, http.StatusBadRequest, err)
		return
	}

	// check if it already exists
	_, err := handler.store.GetContactByPhoneNumber(requestData.Phone)

	// contact exists
	if err != nil {
		utils.WriteError(writer, http.StatusBadRequest, fmt.Errorf("user with phone %s already exists", requestData.Phone))
		return
	}

	// if it doesn't we create the new contact
	if err = handler.store.AddContact(requestData); err != nil {
		utils.WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSONResponse(writer, http.StatusCreated, "user successfuly created")
}

func (handler *Handler) HandleDeleteContact(writer http.ResponseWriter, request *http.Request) {
	id, err := ParseAndValidateParamFromRequestURL(request, "id")
	if err != nil {
		utils.WriteError(writer, http.StatusBadRequest, err)
		return
	}

	if err := handler.store.DeleteContactById(id); err != nil {
		utils.WriteError(writer, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSONResponse(writer, http.StatusCreated, "user successfuly deleted")
}

func (handler *Handler) HandleSearchContact(writer http.ResponseWriter, request *http.Request) {
	id, err := ParseAndValidateParamFromRequestURL(request, "id")
	if err != nil {
		utils.WriteError(writer, http.StatusBadRequest, err)
		return
	}

	contact, err := handler.store.GetContactById(id)
	if err != nil {
		utils.WriteError(writer, http.StatusBadRequest, fmt.Errorf("failed to get contact"))
		return
	}

	utils.WriteJSONResponse(writer, http.StatusCreated, *contact)
}

func (handler *Handler) HandleGetContacts(writer http.ResponseWriter, request *http.Request) {
	page, err := ParseAndValidateParamFromRequestURL(request, "page")
	if err != nil {
		utils.WriteError(writer, http.StatusBadRequest, err)
		return
	}

	contacts, err := handler.store.GetContactsPage(page)
	if err != nil {
		utils.WriteError(writer, http.StatusBadRequest, fmt.Errorf("failed to get contacts"))
		return
	}

	if len(contacts) == 0 {
		utils.WriteError(writer, http.StatusBadRequest, fmt.Errorf("page does not exist"))
		return
	}

	utils.WriteJSONResponse(writer, http.StatusCreated, contacts)
}

func (handler *Handler) HandleUpdateContact(writer http.ResponseWriter, request *http.Request) {
	var requestData types.UpdateContactDTO
	if err := ParseAndValidateRequest(request, &requestData); err != nil {
		utils.WriteError(writer, http.StatusBadRequest, err)
		return
	}

	id, err := ParseAndValidateParamFromRequestURL(request, "id")
	if err != nil {
		utils.WriteError(writer, http.StatusBadRequest, err)
		return
	}

	if err := handler.store.UpdateContact(id, requestData); err != nil {
		utils.WriteError(writer, http.StatusBadRequest, fmt.Errorf("could not update contact"))
		return
	}

	utils.WriteJSONResponse(writer, http.StatusCreated, "user successfuly updated")
}

func ParseAndValidateRequest(request *http.Request, requestData any) error {
	// get payload
	if err := utils.ParseJSONFromHttpRequest(request, requestData); err != nil {
		return err
	}

	// validate payload
	if err := utils.Validate.Struct(requestData); err != nil {
		errors := err.(validator.ValidationErrors)
		return errors
	}

	return nil
}

func ParseAndValidateParamFromRequestURL(request *http.Request, param string) (int, error) {
	// Parse URL parameters
	query := request.URL.Query()

	// Get specific parameter values
	id, err := strconv.Atoi(query.Get(param))
	if err != nil {
		// Handle conversion error
		return -1, fmt.Errorf("invalid %s value", param)
	}

	return id, nil
}
