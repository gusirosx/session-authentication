# Session Authentication
Session Authentication API that manages authentication, cookies and JWT Token's

## Features
- Registered users can authenticate with their email and password.
- Once a user successfully authenticates, a JWT token is created with a session cookie.
- When a user makes a request for a secured route e.g. /home, the session cookie is received by the server and verified. If the cookie value is valid, the user is granted access to the secured route.
- Sessions and cookies have a maximum life before they expire, once expired, the user will need to login again.
- Sessions and cookies are automatically renewed when the user makes a request to any secured routes. This will reset the max life of the session and cookie.
- Session entries along with the session cookie is destroyed upon logging out.

## How to run this application

To run this application, build and run the Go binary:

```sh
go run .
```
Now, using any HTTP client with support for cookies make a login request with the appropriate credentials:
### Request

```
GET http://localhost:8080/
```


### Response
```JSON
{
    "success": "Up and running..."
}
```


```
POST http://localhost:8080/login
```

```JSON
{
    "email":"gusirosx@email.com",
    "password":"123456"
}
```
Inspect the clients cookies to see the value of the `session` cookie, and see the output:

```JSON
{
    "success": "Helo Gustavo"
}
```

You can now try hitting the welcome route from the same client to get the welcome message:
```
GET http://localhost:8000/home
```
```JSON
{
    "response": "Hello gusirox"
}
```

Hit the refresh route, and then inspect the clients cookies to see the new value of the `session`:

```
POST http://localhost:8080/refresh
```

```JSON
{
    "success": "Cookie Refreshed"
}

```

Finally, call the logout route to clear session data:

```
GET http://localhost:8080/logout
```

Calling the welcome and refresh routes after this will result in a `401` error.

```JSON
{
    "error": "No session cookie header provided"
}

```
## License
MIT License

Copyright (c) 2022 Gustavo Rodrigues


Please feel free to create a new issue if you come across one or want a new feature to be added. I am looking for contributors, feel free to send pull requests.





##TODO List##
- [X] Role-based authorization including group permissions
- [X] Improve the API to include a consistent error payload
- [ ] Activity Logs
- [ ] Encrypted passworda
- [ ] Improve documentation
- [ ] Add TLS support for the end point's
- [ ] Add support for databases


