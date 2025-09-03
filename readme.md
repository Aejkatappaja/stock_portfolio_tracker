# Stock Portfolio Tracker

A comprehensive stock portfolio management application built as the final project for CS50x (Introduction to Computer Science) at Harvard University. This project demonstrates full-stack development skills using modern technologies and best practices.

## 🎯 Project Overview

The Stock Portfolio Tracker is a web application that allows users to manage their stock investments by tracking buy/sell transactions, monitoring portfolio performance, and maintaining a comprehensive record of their trading activities.

## ✨ Features

- **Transaction Management**: Add, edit, and delete stock transactions
- **Portfolio Tracking**: Monitor your stock positions and performance
- **Real-time Data**: View transaction history with detailed information
- **Responsive Design**: Modern, mobile-friendly interface
- **Data Persistence**: Secure storage of all transaction data

## 🏗️ Architecture

This project follows a modern microservices architecture with:

- **Backend**: Go (Golang) REST API with Gin framework
- **Frontend**: Vue.js 3 with TypeScript and Tailwind CSS
- **Database**: PostgreSQL for data persistence
- **Containerization**: Docker and Docker Compose for easy deployment
- **Reverse Proxy**: Traefik for load balancing and routing

## 🛠️ Technology Stack

### Backend
- **Language**: Go 1.23.0
- **Framework**: Gin (HTTP web framework)
- **Database**: PostgreSQL with SQLx driver
- **Architecture**: Clean architecture with Repository and Service layers

### Frontend
- **Framework**: Vue.js 3 with Composition API
- **Language**: TypeScript
- **Styling**: Tailwind CSS
- **Build Tool**: Vite
- **HTTP Client**: Axios

### Infrastructure
- **Containerization**: Docker
- **Orchestration**: Docker Compose
- **Reverse Proxy**: Traefik v3.1.0
- **Database**: PostgreSQL Alpine

## 🚀 Getting Started

### Prerequisites

- Docker and Docker Compose installed
- Go 1.23.0+ (for local development)
- Node.js 18+ (for local frontend development)

### Quick Start with Docker

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd stock_portfolio_tracker
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env file with your database credentials
   ```

3. **Start the application**
   ```bash
   docker-compose up -d
   ```

4. **Access the application**
   - Frontend: http://localhost:5173
   - Backend API: http://localhost:8080/api
   - Traefik Dashboard: http://localhost:8080

### Local Development

#### Backend (Go)
```bash
cd src
go mod download
go run main.go
```

#### Frontend (Vue.js)
```bash
cd src-client
npm install
npm run dev
```

## 📊 Database Schema

The application uses a simple but effective database schema:

```sql
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    type VARCHAR(5) NOT NULL CHECK (char_length(type) >= 3 AND char_length(type) <= 5),
    ticker VARCHAR(30) NOT NULL CHECK (char_length(ticker) >= 1 AND char_length(ticker) <= 30),
    volume REAL NOT NULL,
    price REAL NOT NULL,
    date DATE NOT NULL
);
```

## 🔧 API Endpoints

- `GET /api/transactions` - Retrieve all transactions
- `POST /api/transactions` - Create a new transaction
- `PUT /api/transactions/:id` - Update an existing transaction
- `DELETE /api/transactions/:id` - Delete a transaction

## 🎨 User Interface

The application features a clean, modern interface built with:
- **Dashboard**: Overview of all transactions with CRUD operations
- **Transaction Creation**: Form to add new buy/sell transactions
- **Transaction Editing**: Modify existing transaction details
- **Responsive Design**: Works seamlessly on desktop and mobile devices

## 🏛️ Project Structure

```
stock_portfolio_tracker/
├── src/                    # Go backend
│   ├── controller/         # HTTP controllers
│   ├── entity/            # Data models
│   ├── repository/         # Data access layer
│   ├── service/            # Business logic
│   ├── main.go            # Application entry point
│   └── go.mod             # Go dependencies
├── src-client/             # Vue.js frontend
│   ├── src/
│   │   ├── components/     # Vue components
│   │   ├── router/         # Vue Router configuration
│   │   └── utils/          # Utility functions
│   ├── package.json        # Node.js dependencies
│   └── vite.config.ts      # Vite configuration
├── docker-compose.yaml     # Docker services configuration
├── schema.sql              # Database schema
└── README.md               # This file
```

## 🧪 Testing

The application includes comprehensive testing:
- Backend unit tests for services and repositories
- Frontend component testing
- Integration tests for API endpoints

## 🚀 Deployment

The application is containerized and ready for deployment:
- **Development**: Use Docker Compose for local development
- **Production**: Deploy containers to your preferred cloud platform
- **Scaling**: Easy horizontal scaling with Docker Swarm or Kubernetes

## 🤝 Contributing

This project was developed as a CS50x final project, but contributions and improvements are welcome:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## 📝 License

This project is part of the CS50x course at Harvard University.

## 🎓 About CS50x

CS50x is Harvard University's introduction to computer science and programming. This project demonstrates the culmination of learning various programming concepts including:

- **Programming Fundamentals**: Variables, functions, control structures
- **Data Structures**: Arrays, linked lists, hash tables
- **Web Development**: HTML, CSS, JavaScript
- **Databases**: SQL, data modeling
- **Software Engineering**: Version control, testing, deployment

## 👨‍💻 Author

This project was developed as the final project for CS50x certification, demonstrating proficiency in full-stack web development and modern software engineering practices.

---

*Built with ❤️ for CS50x - Harvard University's Introduction to Computer Science*
