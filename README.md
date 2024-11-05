# Snippetbox
This is a Go project following the best practices from the book Let's Go by Alex Edward. This project demonstrates how to configure a Go web application, handle errors gracefully, use a database for dynamic responses, and implement middleware for request processing.

```
/root
|-- /cmd
|   |-- /web            # Entry point (main.go), handlers, helpers, routes
|-- /pkg                
|   |--models           # Database models
|-- /ui                 # HTML files
|-- /static             # CSS, JS files
|-- go.mod              # Go module file
|-- README.md           # Project documentation
```
### Requirements
- Go 1.22.7
- Database (MySQL recommended)
- Any HTTP client (Postman, curl, or browser)

### Features
- **Project Configuration**: Configures the application using environment variables and a configuration struct.
- **Error Handling**: Uses a central error handler for graceful and consistent error management.
- **Database Driven Response**: Uses a database (e.g., PostgreSQL, MySQL) to retrieve dynamic data.
- **Dynamic HTML Rendering**: Serves HTML pages with dynamic data using Go's html/template package.
- **Middleware**: Implements custom middleware for logging, session handling, and more.
