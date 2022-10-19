# FACE IT Coding Challenge

This is a project to handle CRUD for users.

## Structure
The Domain-Driven Design is used to implement this project.
The project uses `MYSQL` as database because the scale of data is small.
Also `gin-gonic` is used to serve a http-server.

- __Config Package:__ This package contains the configurations for the database and the service.
- __Domain Package:__ This package contains the entities, repositories, and logic of the service.
- __Infrastructure Package:__ In this package the connection to database is initialized.
The `users` table is created by the following SQL script:
```sql
CREATE TABLE IF NOT EXISTS users (
    id INT(32) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(32) NOT NULL,
    last_name VARCHAR(32) NOT NULL,
    nick_name VARCHAR(32) NOT NULL,
    password VARCHAR(32) NOT NULL,
    email VARCHAR(32) NOT NULL,
    country VARCHAR(10) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

```
- __Mocks Package:__ This package mocks the behaviour of the interfaces for unit testing.
- __.golangci.yml:__ This file is the configuration for golangci lint.

## How to Run

First we have to get the database up and running. So we use the docker-compose file provided.
```bash
make docker.up
```

This will create a local mysql database on port `3306`.
You don't need to worry about the tables. After running the program the migrations are done automatically while initializing the database connection.

Then we get to run the main program.
```bash
go run main.go
```

This will start a http server on port `:8080`

## Test Coverage
By running the following command you can run all the tests in the project:
```bash
make test
```

Unit testing is used here. I could achieve the following coverages:
- `repository`: 75%.
- `service`: 73.6%.
- `utils`: 92.5%.

## How to Use

There are five APIs in total which are listed below:
- `GET /health`: This API checks the healthiness of the database by checking the ping.
- `POST /v1/users/create`: This API gets the user information and inserts the user in the database.
  - I assumed that the email and nickname must be unique. As a result, if the email or nickname already exist in the database, the API returns an error.
- `POST /v1/users/update`: This API gets the information we want to change for a user, and updates the user in the database.
  - If the user ID passed through the API does not exist in the database, the API returns an error.
  - In addition, if the user ID exists in the database, and we want to update it, the provided information is compared to the user information in the database. 
If there are no changes then the API returns an error.
- `DELETE /v1/users/remove`: This API gets an ID and removes the user with the given ID.
  - If there are no records deleted from the database, for instance, if the provided user ID does not exist in the database, the API returns an error.
- `POST /v1/users/get`: This API returns the users based on the criteria passed as URL Parameters to it. It also handles pagination by the `page` and `page_size` fields passed in the body of the request.
  - This API can handle `country` and `nickname` filters. For instance if the `country=UK` is given, only users who live in the `UK` are returned,
Or by providing `nickname=mehran`, The API will return all the users whose nickname contains `mehran`. Of course you can mix these two criteria.

The more detailed API information with examples can be found in the postman collection [here](https://www.getpostman.com/collections/5348cab405154fa13fc4)

## What to Add in the Future

- This project uses `MySQL` database. As we know this database is not suitable for large scales of data. In the future if there are more users, we can move to `MongoDB`.
- The criteria in the `get users` API can be extended to include other fields as well. Right now, for simplicity I just handled `country` and `nickname`.
