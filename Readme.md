# Order API

Order API is a simple RESTful API for managing orders and items. It is built using Golang, Gin framework, and Gorm for database interactions.

## Table of Contents

- [Order API](#order-api)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
    - [Create Database](#create-database)
    - [Install Dependencies](#install-dependencies)
    - [Run the Project](#run-the-project)
  - [API Documentation](#api-documentation)
  - [Author](#author)

## Installation

### Create Database

1. Install a PostgreSQL database.

2. Create a new database with the name `order_api`.

### Install Dependencies

Make sure you have [Go](https://golang.org/dl/) installed on your machine.

1. Clone the repository

```bash
git  clone  https://github.com/arifpatanduk/hacktiv8-assignment-2.git
```

2. Change directory to the project folder

```bash
cd  hacktiv8-assignment-2
```

3. Install project dependencies

```bash
go  mod  tidy
```

### Run the Project

```bash
go  run  main.go
```

The API will be accessible at http://localhost:8080.

## API Documentation

1. Swagger
   Explore the API using the Swagger documentation available at `http://localhost:8080/swagger/index.html`

2. Postman
   Import the Postman collection available at https://documenter.getpostman.com/view/11542567/2s9YyzcdBT

## Author

[Arif Patanduk](https://github.com/arifpatanduk)
