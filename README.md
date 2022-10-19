# FACE IT Coding Challenge

This is a project to handle CRUD for users.

## Structure
- The Domain-Driven Design is used to implement this project.
- The project uses `MYSQL` as a database because the scale of data is small.
- Also, `gin-gonic` is used to serve an HTTP server.
- `Redis` is used to notify other services of the changes made to the users.

### packages
- __Config Package:__ This package contains the database and service configurations.
- __Domain Package:__ This package contains the services' entities, repositories, and logic.
- __Infrastructure Package:__ In this package the database connection is initialized.
  The following SQL script creates the `users` table:
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

First, we have to get the database up and running. So we use the docker-compose file provided.
```shell
make docker.up
```

This will create a local MySQL database on port `3306`.
You don't need to worry about the tables. After running the program, the migrations are automatically done while initializing the database connection.

Then we get to run the main program.
```shell
make run
```

This will start a HTTP server on port `:8080`

## Test Coverage
By running the following command you can run all the tests in the project:
```shell
make test
```

Unit testing is used here. I could achieve the following coverages:
- `repository`: 74.4%.
- `service`: 73.6%.
- `utils`: 92.5%.

## How to Use

There are five APIs in total, which are listed below:
- `GET /health`: This API checks the healthiness of the database by checking the ping.
- `POST /v1/users/create`: This API gets the user information and inserts the user in the database.
  - I assumed that the email and nickname must be unique. As a result, the API returns an error if the email or nickname already exists in the database.
- `POST /v1/users/update`: This API gets the information we want to change for a user and updates the user in the database.
  - If the user ID passed through the API does not exist in the database, the API returns an error.
  - In addition, if the user ID exists in the database, and we want to update it, the provided information is compared to the user information in the database.
    If there are no changes, then the API returns an error.
  - When a user has changed, the user ID is pushed into a queue in Redis, so that other services can check the queue and be notified of the change.
- `DELETE /v1/users/remove`: This API gets an ID and removes the user with the given ID.
  - If no records are deleted from the database, for instance, if the provided user ID does not exist in the database, the API returns an error.
- `POST /v1/users/get`: This API returns the users based on the criteria passed as URL Parameters to it. It also handles pagination by the `page` and `page_size` fields passed in the request's body.
  - This API can handle `country` and `nickname` filters. For instance if the `country=UK` is given, only users who live in the `UK` are returned,
    Or by providing `nickname=mehran`, The API will return all the users whose nickname contains `mehran`. Of course, you can mix these two criteria.

The more detailed API information with examples can be found in the postman collection [here](https://www.getpostman.com/collections/5348cab405154fa13fc4)

## What to Add in the Future

- This project uses `MySQL` database. As we know, this database is not suitable for large scales of data. In the future, if there are more users, we can move to `MongoDB`.
- The criteria in the `get users` API can be extended to include other fields as well. Right now, I just handled `country` and `nickname` for simplicity.
