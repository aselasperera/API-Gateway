# API Gateway with Go Fiber

This project implements an API Gateway using Go Fiber to route requests to different backend services. It includes routing for user and product services.

## Table of Contents

- [Introduction](#introduction)
- [Architecture](#architecture)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This API Gateway serves as a single entry point for various backend services. It forwards incoming API requests to the appropriate backend service based on the request path.

## Architecture

- **API Gateway**: Handles incoming requests and forwards them to the appropriate backend service.
- **User Service**: Manages user-related operations.
- **Product Service**: Manages product-related operations.

## Project Structure
<img width="448" alt="Screenshot 2024-06-04 at 14 39 46" src="https://github.com/aselasperera/API-Gateway/assets/136217850/527ad295-b34f-455e-ad1f-4d5f38b145e3">


## Prerequisites

- Go 1.16 or later
- Docker (optional, for running services in containers)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/aselasperera/API-Gateway
    cd api-gateway-go-fiber
    ```

2. Install Go Fiber and other dependencies:
    ```sh
    go mod tidy
    ```

3. Set up your backend services (ensure they are running and accessible).

## Usage

1. Update the URLs of your backend services in the `internal/service` files if necessary:
    ```go
    // User Service URL
    usersServiceURL := "http://users-service:3002/api/v1/users"
    
    // Product Service URL
    productsServiceURL := "http://products-service:3003/api/v1/products"
    ```

2. Run the API Gateway:
    ```sh
    go run main.go
    ```

3. The API Gateway will be running on port `3001` by default.

## Testing

Use Postman to test the API Gateway. Here are some example requests:

1. Get all users:
    ```
    GET http://localhost:3001/api/v1/users
    ```

2. Get all products:
    ```
    GET http://localhost:3001/api/v1/products
    ```

### Example Postman Collection

You can import the following collection into Postman for easy testing:

            ```json
            {
                "info": {
                    "_postman_id": "your-collection-id",
                    "name": "API Gateway Tests",
                    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
                },
                "item": [
                    {
                        "name": "Get Users",
                        "request": {
                            "method": "GET",
                            "header": [],
                            "url": {
                                "raw": "http://localhost:3001/api/v1/users",
                                "protocol": "http",
                                "host": [
                                    "localhost"
                                ],
                                "port": "3001",
                                "path": [
                                    "api",
                                    "v1",
                                    "users"
                                ]
                            }
                        }
                    },
                    {
                        "name": "Get Products",
                        "request": {
                            "method": "GET",
                            "header": [],
                            "url": {
                                "raw": "http://localhost:3001/api/v1/products",
                                "protocol": "http",
                                "host": [
                                    "localhost"
                                ],
                                "port": "3001",
                                "path": [
                                    "api",
                                    "v1",
                                    "products"
                                ]
                            }
                        }
                    }
                ]
            }

## Contributing
Contributions are welcome! Please submit a pull request or open an issue to discuss your ideas.
