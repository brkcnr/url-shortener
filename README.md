# URL Shortener

A simple and efficient URL shortening service built with Go. This application allows users to convert long URLs into shorter, more manageable links.

## Features

- Convert long URLs to short URLs
- Copy shortened URLs to clipboard with one click
- Clean and responsive user interface
- Fast and lightweight implementation

## Requirements

- Go 1.23 or higher
- SQLite3

## Installation

1. Clone the repository
```bash
git clone https://github.com/yourusername/url-shortener.git
```

2. Install dependencies
```bash
go mod tidy
```

3. Run the application
```bash
go run cmd/server/main.go

The server will start at `http://localhost:8080`
```

## Usage

1. Open your web browser and navigate to `http://localhost:8080`
2. Enter the long URL you want to shorten in the input field
3. Click "Shorten" button
4. Copy the shortened URL using the "Copy to clipboard" button

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to your branch (`git push`)
5. Open a Pull Request
