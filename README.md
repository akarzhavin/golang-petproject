# Pet Project: Authentication and Blog Services

## Project Overview

This project is designed to deepen my knowledge in Go by building a set of microservices. The main focus is on creating an authentication service that handles user authentication and authorization using JWT tokens. Additionally, a blog service provides simple article management functionalities. The project includes the following components:

- **Backend Services:**
    - `authentication-service`: Handles user authentication and token management.
    - `blog-service`: Manages and serves blog articles.
- **Frontend Service:**
    - A React-based frontend to interact with the backend services.
- **Database:**
    - PostgreSQL for storing user and blog data.
- **Docker Compose:**
    - To orchestrate the services and database.

## Goals

- Deepen knowledge in Go to the level of writing a service.
- Implement an authentication service with:
    - Router
    - Middleware
    - Database integration
    - Docker-compose setup
    - User authentication by credentials and token-based authentication using JWT.
- Separate logic into distinct layers:
    - HTTP (transport) layer
    - Service logic layer
    - Database layer (optional)

## Services Overview

### Authentication Service

The authentication service is responsible for user login and token management. It provides endpoints for obtaining access and refresh tokens, as well as refreshing tokens.

#### Endpoints

- `/api/auth/token/`
    - Method: POST
    - Description: Takes an email and password, and returns access and refresh tokens on successful authentication.
    - Request Body:
      ```json
      {
        "email": "user@example.com",
        "password": "password"
      }
      ```
    - Response:
      ```json
      {
        "error": false,
        "message": "Logged in successfully",
        "data": {
          "access":"ACCESS_TOKEN",
          "refresh":"REFRESH_TOKEN"
        }
      }
      ```

- `/api/auth/token/refresh/`
    - Method: POST
    - Description: Takes a refresh token and returns a new pair of access and refresh tokens.
    - Request Body:
      ```json
      {
        "refresh_token": "REFRESH_TOKEN"
      }
      ```
    - Response:
      ```json
      {
        "error": false,
        "message": "Logged in successfully",
        "data": {
          "access":"ACCESS_TOKEN",
          "refresh":"REFRESH_TOKEN"
        }
      }
      ```

### Blog Service

The blog service provides endpoints for retrieving a list of articles and detailed article content. Access to these endpoints is restricted to authenticated users.

#### Endpoints

- `/api/article/list`
    - Method: GET
    - Description: Returns a list of article previews.
    - Headers:
    ```
    Authorization: Bearer ACCESS_TOKEN
    ```
    - Response:
      ```json
      {
        "error": false,
        "message": "",
        "data": [
          {
            "id": 1,
            "title": "Article 1",
            "image": "https://some.image.url/image.jpeg",
            "text": "Article content",
            "editable": true,
            "created_at": "2024 July 03",
            "updated_at": "2024-07-03 13:00:24"
          },
          ...
        ]
      }
      ```

- `/api/article/{id}`
    - Method: GET
    - Description: Returns the full text of a specific article.
    - Headers:
    ```
    Authorization: Bearer ACCESS_TOKEN
    ```
    - Path Parameter:
        - `id` (integer): The ID of the article.
    - Response:
      ```json
      {
        "error": false,
        "message": "",
        "data": {
            "id": 1,
            "title": "Article 1",
            "image": "https://some.image.url/image.jpeg",
            "text": "Article content",
            "editable": true,
            "created_at": "2024 July 03",
            "updated_at": "2024-07-03 13:00:24"
        }
      }
      ```

## How to Run

1. Installation steps:
   ```sh
   git clone https://github.com/akarzhavin/golang-petproject.git
   cd golang-petproject/project
   make up
   ```
2. Access the frontend service at http://localhost:3000.
3. Sing in with the following credentials:
   ```
   email: admin@example.com
   password: verysecret
   ```
