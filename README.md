# Client-Server Application
This is a simple client-server application that allows a client to send a message through RabbitMQ to a server.
The server changes its state due to command obtained from the message.
It is only one way communication, from client to server. No response is sent back to the client.

## Requirements
- Golang 1.22
- RabbitMQ or Docker with compose plugin installed
- Makefile interpreter 

## Installation
- Clone the repository
- Run `make analyze` to check the code for errors
- Run `make test` to run the tests
- Run `make install` to build the client and server executables to the `bin` directory
- Run `docker compose up -d` to start RabbitMQ in a docker container
- Edit the `.env.dev` file to set the RabbitMQ connection credentials if needed

## Usage
- Run `make run` to start the server
- Run `make send` to start the client and send example messages to the server