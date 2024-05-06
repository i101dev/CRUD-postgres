# RSS Aggregator API

This repository contains the source code for a RESTful API built with the [chi](https://github.com/go-chi/chi) Go package. The API includes user management functionalities and a simple CORS configuration.

## Project Structure

-   **Routes**: Defined in the `routes` package. Handles HTTP routing and CORS configuration.
-   **Controllers**: Contains logic for handling user-related operations.
-   **Handlers**: Defines additional HTTP handlers for various endpoints.

## Routing Configuration

-   The router uses the `chi` package to manage HTTP routes.
-   CORS (Cross-Origin Resource Sharing) settings are applied to allow broad HTTP access.
-   The main routes are nested under the `/v1` prefix.

### CORS Settings

-   **Allowed Origins**: All HTTP and HTTPS origins are allowed.
-   **Allowed Methods**: `GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`.
-   **Allowed Headers**: All headers are allowed.
-   **Exposed Headers**: The `Link` header is exposed.
-   **Max Age**: 300 seconds.

### Available Routes

#### Version 1 (v1)

These routes are mounted under the `/v1` prefix:

-   **GET `/v1/test`**: Test endpoint. Returns a simple success response.
-   **GET `/v1/error`**: Error endpoint. Returns an intentional error response.
-   **POST `/v1/users`**: Create a new user.
-   **GET `/v1/users`**: Get a list of all users.
-   **GET `/v1/users/{id}`**: Get details for a specific user by their ID.
-   **PUT `/v1/users/{id}`**: Update a user's information by their ID.
-   **DELETE `/v1/users/{id}`**: Delete a user by their ID.

## How to Use

1. Clone this repository.
2. Ensure you have Go installed on your system.
3. Run the application to start the server:
    ```bash
    go run main.go
    ```
