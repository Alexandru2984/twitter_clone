# Twitter Clone 🐦

A full-stack Twitter clone built with **Golang**, **React**, and **MongoDB**.

## Features

- 🔐 User authentication (register/login)
- ✍️ Create, read, and delete tweets
- ❤️ Like/unlike tweets
- 👤 User profiles
- 👥 Follow/unfollow users
- 📱 Responsive design

## Tech Stack

### Backend
- **Go (Golang)** - Backend API
- **Gin** - Web framework
- **MongoDB** - Database
- **JWT** - Authentication

### Frontend
- **React** - UI framework
- **Vite** - Build tool
- **React Router** - Routing
- **Axios** - HTTP client

## Project Structure

```
twitter_clone/
├── backend/
│   ├── config/          # Database configuration
│   ├── models/          # Data models
│   ├── routes/          # API routes
│   ├── middleware/      # Auth and CORS middleware
│   ├── main.go          # Entry point
│   └── Dockerfile
├── frontend/
│   ├── src/
│   │   ├── components/  # React components
│   │   ├── pages/       # Page components
│   │   ├── services/    # API services
│   │   ├── context/     # React context
│   │   └── App.jsx      # Main app component
│   └── Dockerfile
└── docker-compose.yml
```

## Prerequisites

- Go 1.21+ (for local development)
- Node.js 20+ (for local development)
- MongoDB (for local development)
- Docker & Docker Compose (for containerized deployment)

## Quick Start with Docker

The easiest way to run the application is using Docker Compose:

```bash
# Clone the repository
git clone https://github.com/Alexandru2984/twitter_clone.git
cd twitter_clone

# Start all services
docker-compose up -d

# View logs
docker-compose logs -f
```

Access the application:
- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **MongoDB**: localhost:27017

To stop the services:
```bash
docker-compose down
```

## Local Development Setup

### Backend Setup

```bash
# Navigate to backend directory
cd backend

# Install dependencies
go mod download

# Set environment variables (optional)
export MONGODB_URI="mongodb://localhost:27017"
export DB_NAME="twitter_clone"
export JWT_SECRET="your-secret-key"
export PORT="8080"

# Run the server
go run main.go
```

The backend API will be available at http://localhost:8080

### Frontend Setup

```bash
# Navigate to frontend directory
cd frontend

# Install dependencies
npm install

# Create environment file
cp .env.example .env

# Start development server
npm run dev
```

The frontend will be available at http://localhost:3000

### MongoDB Setup

If you don't have MongoDB installed locally, you can run it with Docker:

```bash
docker run -d -p 27017:27017 --name mongodb mongo:7.0
```

## API Endpoints

### Authentication
- `POST /api/register` - Register a new user
- `POST /api/login` - Login user
- `GET /api/profile` - Get current user profile (protected)

### Tweets
- `GET /api/tweets` - Get all tweets
- `GET /api/tweets/user/:username` - Get tweets by user
- `POST /api/tweets` - Create a tweet (protected)
- `POST /api/tweets/:id/like` - Like/unlike a tweet (protected)
- `DELETE /api/tweets/:id` - Delete a tweet (protected)

### Users
- `GET /api/users/:username` - Get user profile
- `POST /api/users/:username/follow` - Follow/unfollow user (protected)

## Environment Variables

### Backend (.env or environment)
```
MONGODB_URI=mongodb://localhost:27017
DB_NAME=twitter_clone
JWT_SECRET=your-secret-key
PORT=8080
```

### Frontend (.env)
```
VITE_API_URL=http://localhost:8080/api
```

## Building for Production

### Backend
```bash
cd backend
go build -o server
./server
```

### Frontend
```bash
cd frontend
npm run build
# Serve the dist/ folder with a web server
```

## Usage

1. **Register**: Create a new account at `/register`
2. **Login**: Sign in with your credentials at `/login`
3. **Tweet**: Post tweets from the main feed
4. **Interact**: Like tweets and view other users' profiles
5. **Follow**: Follow users to see their tweets
6. **Profile**: View your profile and tweets

## Development Notes

- The backend uses JWT tokens for authentication
- Passwords are hashed using bcrypt
- The frontend stores the JWT token in localStorage
- CORS is enabled for cross-origin requests
- MongoDB collections are automatically created on first use

## License

MIT License

## Contributing

Feel free to submit issues and pull requests!

## Author

Alexandru2984
