# Chat Messenger Backend

This repository contains the backend application for a chat messaging system built using Go and Clean Architecture.

## Features

- **Real-time messaging:** Implement real-time messaging functionality using WebSockets.
- **User authentication:** Securely authenticate users and manage user sessions.
- **User management:** Create, update, and delete user profiles.
- **Group chats:** Implement group chats where multiple users can participate. [IN DEVELOPMENT]
- **Message history:** Store and retrieve chat message history. [IN DEVELOPMENT]
- **Notifications:** Send notifications for new messages and mentions. [IN DEVELOPMENT]
- **File upload:** Allow users to upload and share files in chat conversations. [IN DEVELOPMENT]

## Acknowledgements

- [Go](https://golang.org)
- [Gin](https://github.com/gin-gonic/gin)
- [Gorilla WebSockets](https://github.com/gorilla/websocket)
- [Viper](https://github.com/spf13/viper)
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)

## Project Description&Structure:
REST API and Websockets with custom JWT-based authentication system. Core functionality is about creating and managing messages on chat.

## Structure:
4 Domain layers:
 - **Models layer**
 - **Repository layer**
 - **UseCase layer**
 - **Delivery layer**

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/YnMann/chat_backend.git

3. Navigate to the project directory:
   ```bash
   cd chat_backend

4. Install dependencies:
   ```bash
   go mod tidy

## Usage
1. Build the project:
   ```bash
   make
   ```
2. Start the backend server
   ```bash
   app.exe

## License
This project is licensed under the MIT License 
