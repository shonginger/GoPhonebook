# GoPhonebook
backend phonebook api project in go


# Starting the server

1. Go into the root directory and run the following:
    a. $ docker-compose build --no-cache. this might take a few minutes
    b. $ docker compose up -d

This will build and create the docker container for the application and the database
The server has a short delay on startup to ensure the DB is up and running.

# APIS
BASE ROUTE
http://localhost:8080/api/v1/<api_method>

### GET METHODS:
    ping - test method server responds with pong:
        - parameters: NONE
        - example: curl -X GET http://localhost:8080/api/v1/ping
    search - search contact by id:
      - parameter: ID - the number representing the contact id
      - example: curl -X GET http://localhost:8080/api/v1/search?id=2
    getContacts - get page consisting of up to 10 contacts:
     - parameter: page - the number representing the page to get
     - example: curl -X GET   http://localhost:8080/api/v1/search?page=2

### POST METHODS:
    add - create a new contact:
        - body - a json cotaining the following: 
            - firstName: required - a string represnting the first name of the contact
            - lastName: a string representing the last name of the contact
            - address: a string representing the address of the contact
            - phone: required, must be between 7 and 20 characeters - a string representing the phone number of the contact
        example:
         curl -X POST   http://localhost:8080/api/v1/add \ 
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
        curl -X PUT   http://localhost:8080/api/v1/update?id=2   -H \
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
        - example: curl -X GET http://localhost:8080/api/v1/delete?id=2 
  


# NOTES:
I didn't really have time to invest in unit tests as I was learning Docker and Golang the go, so I left it pretty barren. how ever here are scenrios I would test

FOR APIS:
 - DELETE
 1. delete ID that doesn't exist.
 2. delete ID that exists and then deleting it again
 - ADD
 1. add user with all fields
 2. add user with no required fields
 3. attempt to add user with required fields missing
 4. attempt to add user with invalid values such as length constraints
 5. adding the same user twice
 - UPDATE
 1. update each field indivually 
 2. update attempts that are invalid. (see ADD)
 - GET
 1. Get contact that exists
 2. Get contact that does not exist
 - GetContacts
 1. Get existing page
 2. Get invalid page (such as -10)
 3. Get page that does not exist
 4. Adding more than the page size then getting both pages, afterwards delete one user and attempt to get the 2nd page.

 # SCALE
 I'm not sure what you meant by that requirement but the code I wrote does not support scale maybe it needs a rate limiter :D 

