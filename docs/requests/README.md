# HTTP Requests

This directory contains API requests in .http format (one endpoint per file), intended for use with REST Client in VS Code.

## Prerequisites

- API running at http://localhost:8080
- VS Code extension: REST Client (Humao)

## Recommended Test Flow

1. Run 01-root.http to confirm the server is online.
2. Run 02-login.http to create the session (session cookie).
3. Run 03-home.http to access a protected route.
4. Run 04-refresh.http to renew the session cookie.
5. Run 05-logout.http to end the session.

## Endpoints

- GET /  
  File: 01-root.http  
  Expected response: 200 with {"success":"Up and running..."}.

- POST /login  
  File: 02-login.http  
  JSON body: email and password.  
  Expected response: 200 and creation of the session cookie.

- GET /home (protected)  
  File: 03-home.http  
  Requires a valid session cookie.  
  Expected response: 200 with a welcome message.

- GET /refresh (protected)  
  File: 04-refresh.http  
  Requires a valid session cookie.  
  Expected response: 200 with {"success":"Cookie Refreshed"}.

- GET /logout (protected)  
  File: 05-logout.http  
  Requires a valid session cookie.  
  Expected effect: expires/clears the session cookie.

## Note

Protected routes (/home, /refresh, /logout) depend on the session cookie. If your HTTP client does not preserve cookies between requests, run POST /login before each protected call.
