# JWT-Based Security for Node.js RESTful APIs

This guide illustrates how to implement JSON Web Tokens (JWT) for securing a
Node.js RESTful CRUD API. The tutorial focuses on integrating JWT with an
existing minimalist API that utilizes MongoDB as its underlying database.

## Components

### 1. User Model and Controller

- The **User model** defines the structure of the data related to users.
- The **User controller** encapsulates the business logic required for database
  interactions.

### 2. Database Connectivity

- The **db file** establishes the connection between the application and the
  MongoDB database.

### 3. Application Bootstrap

- The **app file** serves as the entry point for initializing the application.

### 4. Server Configuration

- The **server file** is responsible for starting the server and specifying the
  port on which it should listen.

### 5. Authentication

- The **auth folder** encompasses the configuration for user registration,
  login, as well as token signing and verification processes.
