# Library Management System

## Overview
This console-based application demonstrates fundamental Go constructs including structs, interfaces, methods, slices, and maps. It models a small library where books can be registered, borrowed, and returned by members.

## Project Structure
- `main.go` – program entry point; wires controllers and services.
- `controllers/library_controller.go` – handles interactive console commands and delegates to the service layer.
- `models/book.go` – defines the `Book` entity.
- `models/member.go` – defines the `Member` entity.
- `services/library_service.go` – implements the `LibraryManager` interface and encapsulates business logic.
- `go.mod` – Go module definition.

## Features
- Add and remove books from the catalogue.
- Register members (optional based on service implementation).
- Borrow and return books with validation.
- Display available books.
- Display books borrowed by a specific member.

## Running the Application
```bash
cd library_management
go run ./...
```

Follow the on-screen menu to interact with the system.

