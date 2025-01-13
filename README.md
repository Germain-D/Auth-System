# Auth System Template

A standardized authentication system template with Golang backend, Nuxt 3 frontend, and PostgreSQL database. This template provides a modern, responsive authentication solution that works seamlessly across desktop and mobile devices.

## Features

- Full-stack authentication system
- Golang backend with Fiber framework
- Nuxt 3 frontend with responsive design
- PostgreSQL database
- JWT authentication
- Multiple authentication providers:
  - Google
  - LinkedIn 
  - Facebook
  - GitHub
  - Email/Password

## Project Structure

```bash
Auth-System/
├── backend/
│   ├── controllers/
│   │   ├── auth.go
│   │   └── oauth.go
│   ├── database/
│   │   ├── connection.go
│   │   └── migrations/
│   ├── models/
│   │   └── user.go
│   ├── routes/
│   │   └── routes.go
│   ├── middleware/
│   │   └── jwt.go
│   ├── config/
│   │   └── config.go
│   └── main.go
│
└── frontend/
    ├── components/
    │   ├── Auth/
    │   │   ├── LoginForm.vue
    │   │   └── RegisterForm.vue
    │   └── Layout/
    ├── pages/
    │   ├── auth/
    │   │   ├── login.vue
    │   │   └── register.vue
    │   └── index.vue
    ├── layouts/
    │   └── default.vue
    ├── middleware/
    │   └── auth.ts
    ├── stores/
    │   └── auth.ts
    ├── public/
    │   └── assets/
    ├── nuxt.config.ts
    └── package.json
```

## Quick Start

### Prerequisites

- Go 1.20+
- Node.js 16+
- PostgreSQL
- Docker (optional)

### Backend Setup

```bash
cd Auth-System-Back
go mod download
go run main.go
```

### Frontend  Setup

```bash
cd Auth-System-Front
npm install
npm run dev
```

### Database Setup

1. Create a PostgreSQL database
2. Configure database connection in .env file
3. Run migrations

### Configuration

Create a .env file in both frontend and backend directories:
# Backend .env

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=auth_system
JWT_SECRET=your_jwt_secret
```

# Frontend .env

```bash
NUXT_PUBLIC_API_BASE=http://localhost:8000
```

### Configuration

Create New Project
```bash
auth-system create \
  --google \
  --linkedin \
  --facebook \
  --github \
  --email \
  --name my-auth-project
  ```

### Contributing
1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

### License
MIT License

Author
Moi



This README provides:
1. Overview of the project
2. Key features
3. Project structure
4. Setup instructions
5. Configuration details
6. Usage example
7. Contributing guidelines

You can customize it further based on specific requirements or additional features you may add to the template.
This README provides:
1. Overview of the project
2. Key features
3. Project structure
4. Setup instructions
5. Configuration details
6. Usage example
7. Contributing guidelines

You can customize it further based on specific requirements or additional features you may add to the template.