# GoPhonebook
backend phonebook api project in go


# Starting the server
...

# APIS
BASE ROUTE
api/v1/<api_method>

### GET METHODS:
    ping - test method server responds with pong:
        - parameters: NONE
        - example: curl -X GET http://localhost:8081/api/v1/ping
    search - search contact by id:
      - parameter: ID - the number representing the contact id
      - example: curl -X GET http://localhost:8081/api/v1/search?id=2
    getContacts - get page consisting of up to 10 contacts:
     - parameter: page - the number representing the page to get
     - example: curl -X GET   http://localhost:8081/api/v1/search?page=2

### POST METHODS:
    add - create a new contact:
        - body - a json cotaining the following: 
            - firstName: required - a string represnting the first name of the contact
            - lastName: a string representing the last name of the contact
            - address: a string representing the address of the contact
            - phone: required, must be between 7 and 20 characeters - a string representing the phone number of the contact
        example:
         curl -X POST   http://localhost:8081/api/v1/add \ 
         -H 'Content-Type: application/json' \
         -d '{
                "firstName": "John",
                "lastName": "Doe",
                "phone": "1234567890"
                "address": "bikini_bottom"
            }'

### PUT METHODS:
    update - update existing contact:
        - parameters: ID - the number representing the contact id
        - body - a json cotaining the following: 
            - firstName:a string represnting the first name of the contact
            - lastName: a string representing the last name of the contact
            - address: a string representing the address of the contact
            - phone: if filled must be between 7 and 20 characeters - a string representing the phone number of the contact
        example: 
        curl -X PUT   http://localhost:8081/api/v1/update?id=2   -H \
        'Content-Type: application/json' \
        -d 
        '{
            "firstName": "John",
            "lastName": "Cena",
            "phone": "5206969"
        }'   

### DELETE METHODS:
    delete - delete existing contact:
        - parameter: ID - the number representing the contact id
        - example: curl -X GET http://localhost:8081/api/v1/delete?id=2 
  


