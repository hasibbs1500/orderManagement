This is a simple order management system built using Go,Gin and GORM for PostgreSQL. It provides RESTful APIs for Creating Orders, Listing Orders and Cancelling Orders. This application also provides authentication functionality allowing users to login and logout.
## Prerequisites
- Docker
## Getting started
Build and start the application using Docker Compose:
````
docker-compose up --build
````
This command will:
- Build the Docker image for the application and the PostgreSQL database.
- Start the containers and set up the required network.
- Start the application server.
- The application server will be available at `http://localhost:8080`