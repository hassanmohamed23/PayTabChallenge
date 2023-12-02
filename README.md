# PayTabChallenge
## Getting started

List Accounts

Description
This endpoint retrieves all accounts available in the system.

Request: localhost:8080/accounts
Method: GET
Response
Status Code: 200 OK
Body: JSON array containing account information

Successful Response Body Example
[
    {
        "id": "3d253e29-8785-464f-8fa0-9e4b57699db9",
        "name": "Trupe",
        "balance": 87.11
    },
    {
        "id": "17f904c1-806f-4252-9103-74e7a5d3e340",
        "name": "Fivespan",
        "balance": 996.15
    }
]


Money Transfer

Description
This endpoint money make transfer between accounts.

Request
Method: POST
Body: JSON payload containing transfer details
Request Body Example

{
  "fromAccount": "accountId1",
  "toAccount": "accountId2",
  "amount": 100.00
}

Responses
Status Code: 200 OK 

Successful Response Body Example
{
  "code": 200,
  "message": "Transfer successful"
}

Status Code: 400 Bad Request 

Error Response Body Example
{
  "code": 400,
  "message": "Invalid request payload"
}

{
  "code": 400,
  "message": "insufficient funds"
}
