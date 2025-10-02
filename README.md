# Proyecto-API-GO


REST API for Product Management (Go)
Project Overview
RESTful API built with Go for managing product data with CRUD operations. This project demonstrates backend development skills using Go's standard library and the Gorilla Mux router for HTTP routing.
Features

Create new products (POST)
Retrieve all products (GET)
Delete products by ID (DELETE)
JSON data handling
In-memory data storage

Technologies Used

Go: Backend programming language
Gorilla Mux: HTTP router and URL matcher
JSON: Data interchange format

Installation
Prerequisites

Go 1.16 or higher installed
Git

Setup
bash# Clone the repository
git clone https://github.com/Jorge20SVJ/product-api-go.git
cd product-api-go

# Install Gorilla Mux dependency
go get -u github.com/gorilla/mux

# Run the application
go run main.go
The server will start on http://localhost:8000
API Endpoints
GET /
Returns welcome message
Response: "Bienvenidos a mi API REST"
GET /Datos
Retrieve all products
jsonResponse: [
  {
    "ID": 1,
    "Nombre": "juan",
    "Genero": "Masculino"
  }
]
POST /Datos
Create a new product
jsonRequest Body:
{
  "Nombre": "Martha",
  "Genero": "Femenino"
}

Response: 201 Created
{
  "ID": 2,
  "Nombre": "Martha",
  "Genero": "Femenino"
}
DELETE /Datos/{id}
Delete a product by ID
Example: DELETE /Datos/1
Response: "El registro 1 se ha eliminado correctamente de la API"
Project Structure
product-api-go/
├── README.md
└── main.go          (API implementation)
Data Model
gotype Base struct {
    ID     int    `json:"ID"`
    Nombre string `json:"Nombre"`
    Genero string `json:"Genero"`
}
Testing with Postman or cURL
Create a product
bashcurl -X POST http://localhost:8000/Datos \
  -H "Content-Type: application/json" \
  -d '{"Nombre":"Pedro","Genero":"Masculino"}'
Get all products
bashcurl http://localhost:8000/Datos
Delete a product
bashcurl -X DELETE http://localhost:8000/Datos/1
Current Limitations

Data is stored in memory (resets on server restart)
No database persistence
No authentication/authorization
Limited error handling

Future Improvements

Add UPDATE (PUT) endpoint for complete CRUD
Implement database persistence (PostgreSQL/MySQL)
Add input validation
Implement error handling middleware
Add authentication (JWT tokens)
Write unit tests
Add logging functionality
Implement pagination for GET requests

Technical Decisions
Why Gorilla Mux?

Powerful URL routing with variable extraction
Clean syntax for defining routes
RESTful design support with method-specific handlers
Well-maintained and widely used in Go community

Why In-Memory Storage?

Simplifies initial implementation and learning
Fast read/write operations for demonstration
Easy to test and understand data flow
Foundation for future database integration

Learning Outcomes
This project demonstrates:

RESTful API design principles
HTTP methods (GET, POST, DELETE)
JSON serialization/deserialization in Go
Router configuration with Gorilla Mux
Basic CRUD operations
Go project structure

Contact
Jorge Jair Sanchez Vazquez
Computer Engineering Graduate - IPN
Email: jairsanchez198@gmail.com
GitHub: https://github.com/Jorge20SVJ
License
This project was developed for educational and portfolio purposes.
