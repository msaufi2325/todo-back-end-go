## Backend Code

The backend for this Todo App is built using Go and PostgreSQL. It provides a RESTful API for managing todos, user authentication, and more.

### Features
- User registration and authentication
- JWT-based authentication
- CRUD operations for todos
- Secure password storage with bcrypt
- Environment variable management with `godotenv`

### Getting Started

#### Prerequisites
- Go 1.22.4 or later
- PostgreSQL

#### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/msaufi2325/todo-back-end-go.git
    cd todo-back-end-go
    ```

2. Copy the example environment file and set your environment variables:
    ```sh
    cp .env.example .env
    ```

3. Update the `.env` file with your database credentials and other configuration settings.

4. Install dependencies:
    ```sh
    go mod tidy
    ```

5. Build the project:
    ```sh
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o todo-backend ./cmd/api
    ```
    or on local server
        ```sh
    go run ./cmd/api
    ```

#### Running the Application

1. Start the PostgreSQL server and create a database for the application.

2. Run the application:
    ```sh
    ./todo-backend
    ```

3. The server will start on port `8081` by default. You can access the API at `http://localhost:8081`.

### API Endpoints

- `POST /register` - Register a new user
- `POST /login` - Authenticate a user and return a JWT
- `POST /refresh-token` - Refresh the JWT
- `POST /logout` - Log out the user
- `GET /todos` - Get all todos
- `POST /todos` - Create a new todo
- `PUT /todos/{id}` - Update a todo
- `DELETE /todos/{id}` - Delete a todo

### Project Structure

- `cmd/api` - Main application entry point
- `internal/repository` - Database repository interfaces
- `internal/repository/dbrepo` - PostgreSQL implementation of the repository interfaces
- `internal/models` - Data models
- `internal/handlers` - HTTP handlers

### Contributing

Contributions are welcome! Please open an issue or submit a pull request.

### License

This project is licensed under the MIT License.