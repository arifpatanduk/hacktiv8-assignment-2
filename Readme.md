# Order API

Order API is a simple RESTful API for managing orders and items. It is built using Golang, Gin framework, and Gorm for database interactions.

## Table of Contents

- [Order API](#order-api)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
    - [Create Database](#create-database)
    - [Install Dependencies](#install-dependencies)
    - [Run the Project](#run-the-project)
  - [Swagger API Documentation](#swagger-api-documentation)
  - [Author](#author)

## Installation

### Create Database

1. Install a PostgreSQL database.

2. Create a new database with the name `order_api`.

### Install Dependencies

Make sure you have [Go](https://golang.org/dl/) installed on your machine.

```bash

# Clone the repository

git  clone  https://github.com/arifpatanduk/hacktiv8-assignment-2.git



# Change directory to the project folder

cd  hacktiv8-assignment-2



# Install project dependencies

go  mod  tidy

```

### Run the Project

```bash

# Run the project

go  run  main.go

```

The API will be accessible at http://localhost:8080.

## Swagger API Documentation

Explore the API using the Swagger documentation available at http://localhost:8080/swagger/index.html.

## Author

[Arif Patanduk](https://github.com/arifpatanduk)
