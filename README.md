# recy
An app to help you recycle your trash, for those without a car 

## Explanation of the Structure
1. cmd/recy/main.go: This is the entry point of your application. It initializes and starts your web server.

2. internal/: This directory contains the core logic of your application.
    * handlers/: Contains HTTP handlers for your web routes.
    * models/: Defines the data models used in your application.
    * services/: Contains business logic and services.
    * storage/: Handles data storage, whether it's a database or in-memory storage.
3. pkg/: This directory contains reusable packages.
    * config/: Manages configuration settings for your application.
    * middleware/: Contains middleware for your HTTP server.
    * utils/: Contains utility functions that can be used across different parts of your application.
4. web/: This directory contains web-related files.
    * static/: Contains static files like CSS, JavaScript, and images.
    * templates/: Contains HTML templates for your web pages.

go.mod and go.sum: These files are used by Go modules to manage dependencies.

README.md: Contains documentation for your project.