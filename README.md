# amartha

## How to run migration
`go run cmd/migration/main.go`

## How to run service
`go run cmd/main.go`

## How to test make payment failed after customer is delinquent
### instead of waiting for two weeks for delinquent customer, we can manipulate data on database
`
- import the postman collection into your postman
- create loan (the created at will be today)
- update `loan.created_at` to 2 weeks ago
- hit API `delinquent payment` to check the it's true or not
- if it's true, try to `make payment` with 1 week amount the expectation is the payment `failed`
`