## Description
This is an ecommerce api that allows users to register, login, view and order products.

## How to run
- Clone repo
- Install dependencies with `go mod download`
- run `migrate-up` from Makefile to apply migrations
- `make run` to run start application

## Roadmap
- [x] connect to database
- [x] register user
- [x] test user registration
- [x] database migrations 
- [x] login user
- [x] jwt authentication
- [x] handle checkout
- [x] create order
- [x] check if items available in the store
- [x] calculate total price
- [x] create product
- [x] fetch all products


## Api routes

> POST /register - register user

> POST /login - login user

> POST /products - create product

> GET /products - fetch all products

> POST /cart/checkout - handle checkout
