# Hello HTMX

A simple Todo application built with Go and [HTMX](https://htmx.org/), demonstrating how to create dynamic web applications with minimal JavaScript using HTMX's powerful HTML extensions.

Additional frontend interactivity powered by [alpine.js](https://alpinejs.dev/).

## Project Structure

```
├── cmd/
│   └── server/       # Main application entry point
├── data/             # Data models
├── handlers/         # HTTP request handlers
├── htmx/             # HTML templates
├── middlewares/      # HTTP middleware functions
├── routers/          # URL routing configuration
├── services/         # Business logic layer
└── static/           # Static assets
```

## Prerequisites

- Go 1.25.0 or higher
- Make (for running Makefile commands)

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/EdgeJay/hello-htmx.git
   cd hello-htmx
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   make dev
   ```

   Or using Go directly:
   ```bash
   go run cmd/server/main.go
   ```

## Features

- Create, Read, Update, and Delete (CRUD) operations for todos
- Real-time UI updates using HTMX
- Server-side rendering with Go templates
- Clean architecture with separation of concerns

## Project Components

- `cmd/server/main.go`: Application entry point and server initialization
- `data/`: Contains data models and structures
- `handlers/`: HTTP request handlers for all CRUD operations
- `htmx/`: HTML templates with HTMX attributes
- `middlewares/`: Custom middleware functions
- `routers/`: URL routing and endpoint definitions
- `services/`: Business logic implementation
- `static/`: Static assets like favicon
