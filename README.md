# AWS Integration App

A full-stack application to integrate with AWS accounts, built with Go for the backend and React for the frontend.

## Features

- User registration and authentication
- AWS account integration
- Secure storage of user data
- REST API for communication between frontend and backend

## Technology Stack

### Backend
- Go
- PostgreSQL
- JWT for authentication
- AWS SDK for Go
- Gorilla Mux for routing

### Frontend
- React
- React Router
- JavaScript/ES6+
- CSS3

## Project Structure

The project is organized into two main parts:

1. **Backend**: Go API server
2. **Frontend**: React application

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Git
- Node.js and npm (for local frontend development)
- Go (for local backend development)

### Setup and Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourname/aws-integration-app.git
   cd aws-integration-app
   ```

2. Start the services using Docker Compose:
   ```
   docker-compose up
   ```

3. Access the application:
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080

### Running Locally (without Docker)

#### Backend

1. Install Go dependencies:
   ```
   cd backend
   go mod download
   ```

2. Set up environment variables (create a .env file based on the example in the repo)

3. Run the server:
   ```
   go run cmd/server/main.go
   ```

#### Frontend

1. Install dependencies:
   ```
   cd frontend
   npm install
   ```

2. Start the development server:
   ```
   npm start
   ```

## API Endpoints

### Authentication

- `POST /api/register` - Register a new user
- `POST /api/login` - Log in a user

### AWS Integration

- `POST /api/verify-aws` - Verify AWS credentials

## Security Considerations

- User passwords are hashed using bcrypt
- JWT tokens are used for authentication
- AWS credentials are validated but not stored in the database
- HTTPS is recommended for production use

## Development Guidelines

### Code Style

- Backend: Follow Go's standard formatting rules
- Frontend: Use ESLint and Prettier

### Testing

- Backend: Use the standard Go testing package
- Frontend: Use Jest and React Testing Library

## Deployment

For production deployment:

1. Update the JWT secret in the environment variables
2. Use HTTPS for all communication
3. Set up proper database backup procedures
4. Consider using a container orchestration system like Kubernetes

## License

[MIT License](LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.