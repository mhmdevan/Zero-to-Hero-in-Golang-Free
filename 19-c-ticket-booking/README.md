# 🎟️ Ticket Booking System

A scalable and secure Ticket Booking System built with Go, Gin, MySQL, and integrated with Stripe for payment processing. The system is containerized using Docker and deployed with Kubernetes for high availability and scalability.

## Features
- **User Registration & Authentication**: Secure user management with hashed passwords and JWT authentication.
- **Event Management**: CRUD operations for managing events.
- **Ticket Booking**: Book tickets for events and track booking status.
- **Payment Processing**: Integrated with Stripe for handling payments.
- **Persistent Storage**: MySQL for data storage with proper error handling and data validation.
- **Containerization**: Dockerized for easy deployment.
- **Orchestration**: Kubernetes configuration for scaling and managing services.

---

## Project Structure

/ticket-booking

├── /controllers
│   booking_controller.go
│   event_controller.go
│   payment_controller.go
│   user_controller.go

├── /models
│   booking.go
│   event.go
│   user.go

├── /database
│   connection.go

├── /middleware
│   auth_middleware.go

├── /utils
│   jwt.go

├── main.go
├── Dockerfile
├── deployment.yaml

├── mysql-deployment.yaml
├── README.md

---

## Setup Instructions
- **Prerequisites**
- **Go**: Install Go
- **MySQL**: Make sure you have a running MySQL instance or use the provided Kubernetes configuration.
- **Stripe**: Sign up for a Stripe account and get your API keys.
- **Docker**: Install Docker
- **Kubernetes**: Install kubectl and set up a Kubernetes cluster (e.g., using Minikube or a cloud provider).

## 1. Set Up Environment Variables
Create a .env file in the root of your project and add the following:

env  
STRIPE_SECRET_KEY=your_stripe_secret_key  
DB_HOST=localhost  
DB_PORT=3306  
DB_USER=appuser  
DB_PASSWORD=apppassword  
DB_NAME=ticket_booking_db  

Note: For Kubernetes deployment, set these variables in the deployment.yaml file.

## 2. Run Locally
Install Dependencies  
``` go mod tidy ```  

Run the Application  
``` go run main.go ```

## 3. Using Docker
Build the Docker Image
``` docker build -t ticket-booking-app . ```

Run the Docker Container
``` docker run -p 8080:8080 -e STRIPE_SECRET_KEY=your_stripe_secret_key ticket-booking-app ```

## 4. Using Kubernetes
Apply the MySQL Deployment
``` kubectl apply -f mysql-deployment.yaml ```
Apply the Ticket Booking App Deployment
``` kubectl apply -f deployment.yaml ```


## Future Enhancements
- **Add more payment options: Integrate PayPal or other gateways.**
- **Improve error handling: Implement a more comprehensive error-handling strategy.**
- **Add monitoring: Use tools like Prometheus and Grafana for monitoring.**
- **Automate deployment: Use CI/CD pipelines for automated builds and deployments.**

Happy coding! 🎉
