# Session Cookie Authentication
Login API that manages authentication, cookies and id Token from Firebase

## Running our application

To run this application, build and run the Go binary:

```sh
go run .
```
Now, using any HTTP client with support for cookies (e.g. Postman or Insomnia) make a login request with the appropriate credentials:

```
POST http://localhost:8080/login
```

```JSON
{
    "email":"gusirosx@email.com",
    "password":"123456"
}
```
Inspect the clients cookies to see the value of the `session` cookie:

You can now try hitting the welcome route from the same client to get the welcome message:
```
GET http://localhost:8000/home
```
Hit the refresh route, and then inspect the clients cookies to see the new value of the `session`:

```
POST http://localhost:8080/refresh
```

Finally, call the logout route to clear session data:

```
GET http://localhost:8080/logout
```

Calling the welcome and refresh routes after this will result in a `401` error.
