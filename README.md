# SnippetHub
Simple fullstack project to learn Go. The goal is to build a website like pastebin or such.
![homepage](/images/homepage.png "Snippethub homepage")
## Routes
- TODO: auth routes

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

## Setup (todo)
- Create `"snippets" table`
- Create `"sessions" table`
- Create `"users" table`

### MySQL
### TLS
- `mkdir tls`
- `cd tls`
- `go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost`
