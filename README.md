# Session Authentication
Session Authentication API that manages authentication, cookies and JWT Token's

![walkthrough](https://user-images.githubusercontent.com/61150315/161883403-e5382edc-1078-4492-8997-e00b5f62f0ad.gif)

## Features
- Registered users can authenticate with their email and password.
- Once a user successfully authenticates, a JWT token is created with a session cookie.
- When a user makes a request for a secured route e.g. /home, the session cookie is received by the server and verified. If the cookie value is valid, the user is granted access to the secured route.
- Sessions and cookies have a maximum life before they expire, once expired, the user will need to login again.
- Session entries along with the session cookie is destroyed upon logging out.

## How to run this application

To run this application, build and run the Go binary:

```sh
go run .
```

Check if the application is running
###### Request
```
GET http://localhost:8080/
```
###### Response
```JSON
{
    "success": "Up and running..."
}
```

Now, using any HTTP client with support for cookies make a login request with the appropriate credentials:

```
POST http://localhost:8080/login
```
###### Request
```JSON
{
    "email":"gusirosx@email.com",
    "password":"123456"
}
```
Inspect the clients cookies to see the value of the session cookie, and see the output:
###### Response
```JSON
{
    "success": "Helo Gustavo"
}
```

You can now try hitting the home route from the same client to get the welcome message.
###### Request
```
GET http://localhost:8080/home
```
###### Response
```JSON
{
    "response": "Hello gusirox"
}
```

Hit the refresh route, and then inspect the cookie of your client and see the new value of the session cookie.

###### Request
```
GET http://localhost:8080/refresh
```
###### Response
```JSON
{
    "success": "Cookie Refreshed"
}
```

Finally, call the logout route to clear session data.

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

## List of improvements to be made
- [ ] Role-based authorization including group permissions
- [ ] Improve the API to include a consistent error payload
- [ ] Activity Logs
- [ ] Encrypted password
- [ ] Improved documentation
- [ ] Add TLS support for the end point's
- [ ] Add support for databases
- [ ] Sessions and cookies are automatically renewed when the user makes a request to any secured routes

Please feel free to create a new issue if you come across one or want a new feature to be added. 