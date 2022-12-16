# SnippetHub
Simple fullstack project to learn Go. The goal is to build a website like pastebin or such.

Feel free to fork or copy the code and use it as a base for your own projects.

![homepage](/images/homepage.png "Snippethub homepage")

## Features

- Frontend: dynamic HTML templates featuring Golang's template language
- Backend: MySQL database
- Uses HTTPS
- User authentication
- Session management
- CSRF protection
- Secure password hashing
- Unit and end-to-end tests

## API Routes

| Method | Pattern | Handler | Action |
| ----------- | ----------- | ----------- | ----------- |
| GET | / | home | Display the home page |
| GET | /snippet/view/:id | snippetView | Display a specific snippet |
| GET | /snippet/create | snippetCreate | Display a HTML form for creating a new snippet |
| POST | /snippet/create | snippetCreatePost | Create a new snippet |
| GET | /user/signup | userSignup | Display a HTML form for signing up a new user |
| POST | /user/signup | userSignupPost | Create a new user |
| GET | /user/login | userLogin | Display a HTML form for logging in a user |
| POST | /user/login | userLoginPost | Authenticate and login the user |
| POST | /user/logout | userLogoutPost | Logout the user |
| GET | /static/*filepath | http.FileServer | Serve a specific static file |

## Requirements
- go 1.19
- mysql 8.0.31

## Local Setup
### MySQL
1 - Create database: `CREATE DATABASE snippethub CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;`

2 - Create specific user (you can use `root`or any other):
```
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippethub.* TO 'web'@'localhost';
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';
```
> The password can be changed as long as the connection string inside `cmd/web/main.go` is updated (line 32)

3 - Create "snippets", "sessions", "users" tables:
```
CREATE TABLE snippets (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
title VARCHAR(100) NOT NULL,
content TEXT NOT NULL,
created DATETIME NOT NULL,
expires DATETIME NOT NULL
);
CREATE INDEX idx_snippets_created ON snippets(created);

CREATE TABLE sessions (
token CHAR(43) PRIMARY KEY,
data BLOB NOT NULL,
expiry TIMESTAMP(6) NOT NULL
);
CREATE INDEX sessions_expiry_idx ON sessions (expiry);

CREATE TABLE users (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
name VARCHAR(255) NOT NULL,
email VARCHAR(255) NOT NULL,
hashed_password CHAR(60) NOT NULL,
created DATETIME NOT NULL
);
ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email
```
> You can choose to populate the "snippets" or "users" table before running. If so you can adapt the code from `internal/models/testdata/setup.sql`
### TLS
- `mkdir tls`
- `cd tls`
- `go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost`

## Running
`go run ./cmd/web -addr=":4000"` or `make run`

## Testing

### Local Setup
A test database should be setup to perform the required tests correctly.

Please repeat steps 1 and 2 from the `MySQL` section inside `Local Setup` as such:

1 - Create database: `CREATE DATABASE test_snippethub CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;`

2 - Create specific user (you can use `root`or any other):
```
CREATE USER 'test_web'@'localhost';
GRANT SELECT, INSERT, UPDATE, DELETE ON test_snippethub.* TO 'test_web'@'localhost';
ALTER USER 'test_web'@'localhost' IDENTIFIED BY 'pass';
```
> The password can be changed as long as the connection string inside `internal/models/testutils_test.go` is updated (line 11)

### Running Tests
Please check the commands with a `test-` prefix inside `Makefile`.

### TODO
- Tests: improve coverage;
- Improve project setup and config
  - Add dockerfile + docker-compose to ease local setup
- Maybe migrate to Postgres before docker stuff
- TBD