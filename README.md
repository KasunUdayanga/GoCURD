
# Book Management API

This project implements a RESTful API to manage books, allowing users to perform CRUD operations. It also includes a search functionality with performance optimization using Go's concurrency features (goroutines and channels). Additionally, the project is designed with the option of containerization using Docker and deployment on a Kubernetes cluster.

## Table of Contents

1. [Project Overview](#project-overview)
2. [Technologies Used](#technologies-used)
3. [API Endpoints](#api-endpoints)
4. [Concurrency Optimization](#concurrency-optimization)
5. [Docker Setup](#docker-setup)
6. [Unit Tests](#unit-tests)
7. [Kubernetes Deployment](#kubernetes-deployment)
8. [How to Run the Project Locally](#how-to-run-the-project-locally)
9. [How to Run the Project in Docker](#how-to-run-the-project-in-docker)
10. [How to Run the Project on Kubernetes](#how-to-run-the-project-on-kubernetes)

---

## Project Overview

This project provides an API to manage a collection of books. The API allows users to create, retrieve, update, and delete books in the system. The book data is stored in memory (a map) for simplicity, and the persistence is simulated using a text file. The project includes:

- A REST API with CRUD operations.
- A keyword search functionality based on the book title and description.
- Performance optimization using concurrency (goroutines and channels).
- Docker containerization for deployment.
- Pagination for retrieving books.
- Unit tests for the API.

---

## Technologies Used

- **GoLang**: The backend programming language.
- **Gorilla Mux**: HTTP router and URL matcher for Go.
- **JSON**: For data exchange between the API and clients.
- **Docker**: For containerizing the application.
- **Kubernetes**: For deploying the application on a Kubernetes cluster.

---

## API Endpoints

The following API endpoints are available:

### 1. **GET /books**
   - **Description**: Retrieves a list of all books in the system.
   - **Response**: A JSON array of books.

### 2. **POST /books**
   - **Description**: Creates a new book in the system.
   - **Request Body**: A JSON object containing the book details.
   - **Response**: A JSON object containing the created book.

### 3. **GET /books/{id}**
   - **Description**: Retrieves a specific book by its ID.
   - **Response**: A JSON object of the book.

### 4. **PUT /books/{id}**
   - **Description**: Updates a book by its ID.
   - **Request Body**: A JSON object containing the updated book details.
   - **Response**: A JSON object of the updated book.

### 5. **DELETE /books/{id}**
   - **Description**: Deletes a book by its ID.
   - **Response**: HTTP status code 204 (No Content).

### 6. **GET /books/search?q=<keyword>**
   - **Description**: Searches for books by keyword in the title and description fields.
   - **Response**: A JSON array of books matching the search keyword.

---

## Concurrency Optimization

To optimize the search process for large datasets, the search functionality uses **goroutines** and **channels** to parallelize the search. The search process is divided into smaller tasks, each handled by a separate goroutine. Results are then collected using channels and merged into a single response.

---

## Docker Setup

To containerize the application, a **Dockerfile** has been included. To build and run the application in Docker:

1. **Build the Docker Image**:
   ```bash
   docker build -t book-api .
   ```

2. **Run the Docker Container**:
   ```bash
   docker run -p 8000:8000 book-api
   ```

This will expose the API on `http://localhost:8000`.

---

## Unit Tests

Unit tests for the API are included in the `tests` directory. The tests use Go's built-in `testing` package to verify the correctness of the API's functionality. An example of a unit test for the `GET /books/{id}` endpoint:

```go
func TestGetBookByID(t *testing.T) {
    req, err := http.NewRequest("GET", "/books/1", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(getBookByID)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Expected status 200, got %v", status)
    }
}
```

Run tests with the following command:

```bash
go test ./tests
```

---


## Kubernetes Deployment

To deploy the application on Kubernetes:

1. Install Minikube or Kind to run Kubernetes locally.
2. Apply the Kubernetes manifest files for deployment and service.

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

Ensure the application is exposed correctly, and check the status using:

```bash
kubectl get pods
```

---

## How to Run the Project Locally

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-repo/book-api.git
   ```

2. **Navigate to the project directory**:
   ```bash
   cd book-api
   ```

3. **Install Go dependencies**:
   ```bash
   go mod tidy
   ```

4. **Run the application**:
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:8000`.

---

## How to Run the Project in Docker

1. **Build the Docker image**:
   ```bash
   docker build -t book-api .
   ```

2. **Run the Docker container**:
   ```bash
   docker run -p 8000:8000 book-api
   ```

The application will be available at `http://localhost:8000`.

---

## How to Run the Project on Kubernetes

1. **Build the Docker image**:
   ```bash
   docker build -t book-api .
   ```

2. **Push the image to a Docker registry** (e.g., Docker Hub):
   ```bash
   docker tag book-api your-dockerhub-username/book-api
   docker push your-dockerhub-username/book-api
   ```

3. **Deploy on Kubernetes** using Minikube or Kind, and apply the deployment and service YAML files.

---

## Conclusion

This project demonstrates how to build a simple but functional book management system with CRUD operations, search functionality, performance optimization, and deployment capabilities. It also includes bonus tasks like Docker containerization, unit testing, and Kubernetes deployment, offering flexibility and scalability for real-world applications.
