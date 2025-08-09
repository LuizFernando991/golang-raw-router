# Golang API - Pure HTTP Router

This project is an **API in Go** built **without any HTTP framework** (such as Gin or Echo).  
The routing is manually implemented using only the standard `net/http` library and regular expressions to map the routes.

---

## 🚀 Features

- **Custom router** implementation with support for dynamic parameters.
- **Route grouping** (`RouteGroup`) with middleware support.
- Controller structure to separate responsibilities.
- Example route `POST /user/create-user` that reads the request body and returns a mock created user.

---

## 📂 Project Structure

```
.
├── api
│   ├── controllers
│   │   └── user.go
│   ├── main.go
│   ├── middlewares
│   │   ├── json.go
│   │   └── logger.go
│   └── router
│       ├── controllers_factory.go
│       ├── route_context.go
│       ├── route_init.go
│       ├── router.go
│       └── routes.go
├── infra
│    ├── config
│    │   ├── env.go
│    │   └── logger.go
│    └── database
│        └── conn.go
├── go.mod
├── go.sum
├── .env
├── .env.example
├── .gitignore
├── Makefile
├── README.md
```

---

## 📌 Request Example

To create a test user, send a `POST` request to `/user/create-user`:

### **Request**

```bash
curl -X POST http://localhost:3000/user/create-user \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao@example.com"
  }'
```

### **Response**

```bash
{
  "id": 1,
  "name": "João Silva",
  "email": "joao@example.com"
}
```

## 🛠 How to Run
1. Clone the repository

```bash
git clone https://github.com/LuizFernando991/golang-raw-router
cd golang-raw-router
```

2. Install dependencies
```bash
go mod tidy
```
3. Run the server
```bash
make run
```
or
```bash
go run ./api/main.go
```
The server will be available at: http://localhost:3000


## 🧠 Notes about the router
- Routing is handled using regular expressions to capture dynamic parameters (/user/{id}).

- Middlewares are applied manually in reverse registration order.

- No external HTTP dependencies are used, only the standard library.
