# Wex TAG and Gateways Product Brief - Transaction Converter

## Description
Transaction Converter is a RESTful API service built in GoLang that allows users to store purchase transactions and retrieve them in a specified country's currency.

## Features
- Store purchase transactions with a description, transaction date, and amount in USD.
- Retrieve stored transactions converted to supported currencies.

## Usage
1. Store Transaction:

Send a POST request to /transactions endpoint with the following JSON payload:


```json
{
    "description": "Grocery shopping",
    "transaction_date": "2024-03-05",
    "purchase_amount_usd": 45.75
}
```


2. Retrieve Transaction:

Send a GET request to /transactions/{id} endpoint with a query parameter currency specifying the target currency. Example: ?currency=Real. Please note that you need to retrieve a transaction that was previously stored.

Examples of Currencies you can use: 
| Currency |
|----------|
| Franc    |
| Euro     |
| Peso     |
| Pound    |
| Real     |
| Yen      |

### Examples
#### Store Transaction
```
curl -X POST 'http://localhost:8080/transactions' \
     -H 'Content-Type: application/json' \
     -d '{
            "description": "Grocery shopping",
            "transaction_date": "2024-03-05",
            "purchase_amount_usd": 45.75
         }'
```

Response:

```json
HTTP/1.1 201 Created
Content-Type: application/json

{
  "id": "b1c3d6fc-933c-4b7a-808b-dbbb3e171d09",
  "description": "Grocery shopping",
  "transaction_date": "2024-03-05T00:00:00Z",
  "purchase_amount_usd": 45.75
}
```
##### Retrieve Transaction

```
curl -X GET 'http://localhost:8080/transactions/b1c3d6fc-933c-4b7a-808b-dbbb3e171d09?currency=Real'

```

Response:

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
  "ID": "b1c3d6fc-933c-4b7a-808b-dbbb3e171d09",
  "Description": "Grocery shopping",
  "TransactionDate": "2024-03-05T00:00:00Z",
  "OriginalAmount": 45.75,
  "ExchangeRate": 4.852,
  "ConvertedAmount": 221.98
}
```

### Installation
Clone the repository:


```
git clone https://github.com/barbaramariani/transaction.git
```
Navigate to the project directory:
```cd transaction```

Build the project:
```go build```

Run the project:
```go run main.go```

##### Running Tests
Execute the following command to run the tests:
```go test ./...```


### Dockerization
To run the application using Docker:

Build the Docker image:
```docker build -t transaction-converter .```

Run the Docker container:
```docker run -p 8080:8080 transaction-converter```