# Car Rental System

Full-stack car rental management system built with Go, PostgreSQL, React, TypeScript, Docker and JWT authentication.

## Features

### Admin CRM

* Authentication (Register / Login)
* Dashboard with statistics
* Cars management (CRUD)
* Customers management (CRUD)
* Bookings management
* Booking status updates
* Validation for phone and IIN
* JWT protected API

### Customer Website

* Browse available cars
* Submit booking requests
* Create customer profile automatically
* Booking requests are sent directly to CRM

## Tech Stack

### Backend

* Go
* Gin
* PostgreSQL
* JWT
* Docker

### Frontend

* React
* TypeScript
* Vite
* Ant Design
* Axios

### Infrastructure

* Docker
* Docker Compose

## Project Structure

```text
car-rental-system/
├── backend/
├── frontend-admin/
├── frontend-client/
├── docker-compose.yml
└── README.md
```

## Running with Docker

### Build and start

```bash
docker compose up --build
```

### Stop

```bash
docker compose down
```

## Services

| Service          | URL                   |
| ---------------- | --------------------- |
| Backend API      | http://localhost:8080 |
| Admin Panel      | http://localhost:5173 |
| Customer Website | http://localhost:5174 |

## API Examples

### Cars

```http
GET /api/cars
POST /api/cars
PUT /api/cars/:id
DELETE /api/cars/:id
```

### Customers

```http
GET /api/customers
POST /api/customers
```

### Bookings

```http
GET /api/bookings
POST /api/bookings
PATCH /api/bookings/:id/status
```

## Booking Workflow

Customer submits booking request

↓

Customer record is created

↓

Booking created with status "new"

↓

Admin reviews booking

↓

Status changes to:

* approved
* completed
* cancelled

## Future Improvements

* Role-based access control
* Email notifications
* Payment integration
* Booking calendar
* Automated migrations
* Deployment with Railway

## Author

Mukhammed Kaldybay
