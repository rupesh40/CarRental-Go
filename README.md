# Car Rental Go

## Overview

This is a Go-based Car Rental system that allows users to sign up, log in, manage cars, book and return rentals. The application supports both local and Docker-based environments.

---

## Requirements

Before running the project, ensure you have the following installed:
- Go (1.19.1 version or above)
- Docker (if using Docker environment)
- PostgreSQL (for local development)

---

## Setup

### 1. Running Locally (Without Docker)

1. Clone the repository:
    ```bash
    git clone <repository-url>
    cd <repository-directory>
    ```

2. Install Go dependencies:
    ```bash
    go mod tidy
    ```

3. Set up PostgreSQL locally:
    - Create a PostgreSQL user.
    - Create a database.
    - Grant the database access to the created user.

4. Update the `.env` file with the following credentials:
    ```env
    DB_USERNAME=<your-db-username>
    DB_PASSWORD=<your-db-password>
    DB_NAME=<your-db-name>
    DB_HOST=localhost
    ```

5. Run the application:
    ```bash
    go run cmd/main.go
    ```

### 2. Running with Docker

1. To start the application with Docker, run the following command:
    ```bash
    docker-compose up
    ```

2. To stop the application and containers, run:
    ```bash
    docker-compose down
    ```

---

## API Usage

### 1. **Sign Up a New User**
- **Endpoint**: `POST http://localhost:3000/user/signup`
- **Request Body**:
    ```json
    {
        "first_name": "userF",
        "last_name": "userL",
        "email": "user@gmail.com",
        "phone": "1234567890",
        "password": "password",
        "user_type": "ADMIN"
    }
    ```
    You can change the `user_type` to `USER`.

---

### 2. **Login User**
- **Endpoint**: `POST http://localhost:3000/user/login`
- **Request Body**:
    ```json
    {
        "email": "user@gmail.com",
        "password": "password"
    }
    ```
- **Response**: The response will contain a `token` field (JWT token).
  
    **Example Response**:
    ```json
    {
        "token": "your-jwt-token"
    }
    ```
    Copy the JWT token from the response to use in subsequent requests.

---

### 3. **Add a New Car**
- **Endpoint**: `POST http://localhost:3000/cars/AddCar`
- **Request Body**:
    ```json
    {
        "carModel": "fc-001",
        "date_of_manufacture": "2021-04-01",
        "last_service_date": "2022-02-06",
        "last_used_date": "2023-01-04",
        "status": "Rented"
    }
    ```
    **Important**: Don’t forget to add the JWT token in the request header.
    
    - Header Key: `token`
    - Header Value: `your-jwt-token`

---

### 4. **Get All Cars**
- **Endpoint**: `GET http://localhost:3000/cars`
- Alternatively, you can use a GraphQL query:
    ```graphql
    query GetAllCars {
        cars {
            id
            car_model
            date_of_manufacture
            last_service_date
            last_used_date
            status
        }
    }
    ```

    **Important**: Don’t forget to add the JWT token in the request header.

---

### 5. **Book a Car**
- **Endpoint**: `GET http://localhost:3000/cars/bookCar/:carID`
- Replace `:carID` with the actual car ID you want to book. Example:
    ```bash
    http://localhost:3000/cars/bookCar/4ef1af7e-3076-4c29-9358-4365cd42f0d8
    ```

    **Important**: Don’t forget to add the JWT token in the request header.

---

### 6. **Return a Car**
- **Endpoint**: `GET http://localhost:3000/cars/returnCar/:carID`
- Replace `:carID` with the actual car ID you want to return. Example:
    ```bash
    http://localhost:3000/cars/returnCar/4ef1af7e-3076-4c29-9358-4365cd42f0d8
    ```

    **Important**: Don’t forget to add the JWT token in the request header.

---

### 7. **Logout User**
- **Endpoint**: `GET http://localhost:3000/users/logout`

    **Important**: Don’t forget to add the JWT token in the request header.

---

## Testing API Endpoints

You can use **Postman** or any other API testing tool to interact with the API. For each request, ensure you add the JWT token obtained during login in the `Authorization` header.

---

## Notes

- Replace all `<placeholders>` in the provided code with your actual values.
- The JWT token must be added to the header for all requests requiring authentication.
- Ensure your PostgreSQL service is running if using a local environment.

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
