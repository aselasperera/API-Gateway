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
api-gateway-go-fiber/
│
├── cmd/
│   └── main.go           # Entry point of the application
│
├── config/
│   └── config.go         # Configuration settings
│
├── internal/
│   ├── handler/
│   │   ├── users.go      # User routes handlers
│   │   └── products.go   # Product routes handlers
│   │
│   ├── routes/
│   │   └── routes.go     # API routes definitions
│   │
│   └── service/
│       ├── users.go      # User service integration
│       └── products.go   # Product service integration
│
├── go.mod                # Go module file
├── go.sum                # Go dependencies
├── Dockerfile            # Dockerfile for containerization
└── README.md             # Project README file
