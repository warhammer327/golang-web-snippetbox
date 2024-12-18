# Snippetbox
This project demonstrates how to configure a Go web application, handle errors gracefully, use a database for dynamic responses, and implement middleware for request processing.

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
- **Database Driven Response**: Uses a database (MySQL recommended) to retrieve dynamic data.
- **Dynamic HTML Rendering**: Serve HTML pages with dynamic data using Go's html/template package.
- **Middleware**: Implements custom middleware for logging, session handling, and more.

### Acknowledgements
Thanks to Alex Edwards for his excellent book Let's Go, which provided the foundational principles and best practices followed in this project.
