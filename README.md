# Human Resource Management System (HRMS) Golang & MySQL

A scalable, modular **Human Resource Management System** built with **Golang** and **MySQL**, featuring robust student attendance tracking, clean API design, and automated background reporting.

## Key Features

- **RESTful API Architecture** using `gin-gonic`
- **Attendance Management** for student records
- **Background Job Scheduling** for weekly and monthly report generation (`robfig/cron/v3`)
- **Clean, Modular Codebase** designed for clarity and maintainability
- **Database Migrations** with `goose`
- **Fully Containerized Setup** using Docker & Docker Compose

## Technical Architecture

### **Tech Stack**
- **Backend:** Go 1.25 (Gin Framework)
- **Database:** MySQL (Dockerized)
- **Migrations:** Goose
- **Containerization:** Docker & Docker Compose

## Setup Instructions

### **Prerequisites**
- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)

### **Installation**

```bash
# Clone the repository
git clone https://github.com/Cheemx/go-hrms.git
cd go-hrms

# Install goose for migrations
go install github.com/pressly/goose/v3/cmd/goose@latest 

# Start MySQL and infrastructure services
docker compose up -d

# Create environment file
cp .env.sample .env
# Ensure port and DB_URL values match your MySQL container (default 3307)

# Run migrations
make migrationUp

# Start the backend API
go run .
```

The API will be available at: http://localhost:8080
## API Endpoints

### **Create Student**
```
POST /api/students
```
**Request**
```json
{
    "email": "cheems@yahoo.com",
    "name": "Cheems",
    "department": "backend"
}
```

**Response**
```json
{
    "id": "2b022397-c5c8-11f0-b5c3-2edc308ce685",
    "name": "Cheems",
    "department": "backend",
    "email": "cheems@yahoo.com",
    "created_at": "2025-11-20T04:19:55Z"
}
```

### **List Students**
```
GET /api/students
```

**Response**

```json
[
    {
        "id": "2b022397-c5c8-11f0-b5c3-2edc308ce685",
        "name": "Cheems",
        "email": "cheems@yahoo.com",
        "department": "frontend",
        "created_at": "2025-11-20T04:19:55Z"
    }
]
```

### **Get Student by ID**

```
GET /api/students/:id
```

**Response**

```json
{
    "id": "2b022397-c5c8-11f0-b5c3-2edc308ce685",
    "name": "Cheems",
    "email": "cheems@yahoo.com",
    "department": "backend",
    "created_at": "2025-11-20T04:19:55Z"
}
```

### **Update Student**

```
PUT /api/students/:id
```

**Request**

```json
{
    "name": "Cheems",
    "department":"frontend"
}
```

**Response**

```json
{
    "id": "2b022397-c5c8-11f0-b5c3-2edc308ce685",
    "name": "Cheems",
    "email": "cheems@yahoo.com",
    "department": "frontend",
    "created_at": "2025-11-20T04:19:55Z"
}
```

### **Delete Student**

```
DELETE /api/students/:id
```

**Response**

```json
{
    "message": "Deletion Successful"
}
```

### **Mark Attendance**

```
POST /api/attendance/mark
```

**Request**

```json
{
    "student_id":"2b022397-c5c8-11f0-b5c3-2edc308ce685",
    "status":"PRESENT"
}
```

**Response**

```json
{
    "message": "attendance marked successfully for id: 2b022397-c5c8-11f0-b5c3-2edc308ce685"
}
```

### **Get Attendance Summary**

```
GET /api/attendance/:student_id
```

**Response**

```json
[
    {
        "name": "Cheems",
        "email": "cheems@yahoo.com",
        "status": "PRESENT",
        "date": {
            "Time": "2025-11-20T00:00:00Z",
            "Valid": true
        }
    }
]
```

## Validation & Idempotency

* **Email Validation:** using Go’s built-in `net/mail` package
* **Duplicate Emails:** prevented via `UNIQUE` DB constraint
* **Attendance for a Day:** enforced via `(student_id, date)` unique constraint
* **Idempotent Behavior:**

  * Creating a student with an existing email returns `400 Bad Request`
  * Attempting duplicate attendance for the same day fails safely
  * Deleting a student multiple times does not cause errors

## Background Cron Jobs

Automated attendance summaries using `robfig/cron/v3`:

* **Weekly Reports:** Generated every week
* **Monthly Reports:** Generated at the start of each month

These background tasks can later be extended to send emails or export reports.

## Testing

Due to time constraints, testing was performed manually using real API calls against the running service.
Unit and integration tests can be added for production hardening.

## Postman Collection Link
[Postman Collection](https://warped-meadow-913182.postman.co/workspace/New-Team-Workspace~850b93a7-4078-4f7e-bcb5-331e137d6e73/collection/32759292-25fc068a-2fad-4aba-9c61-20e6a98a1b79?action=share&creator=32759292)

## Conclusion

*If you've read it till this end, consider giving a star!*

*Built with ❤️ using Go, designed for scale and performance by Cheems!*