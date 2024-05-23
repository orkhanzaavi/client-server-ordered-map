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

## Task

Your task is to write a **Client-Server application** that implements a server that reads commands from an external queue, and executes the command as well as client that accept commands from the standard input or a file and sends them to the external queue.

### Server

- Has a data structure that holds the data in the memory (Ordered Map for C++)
- Read messages (commands) from an external queue
- Server should be able to add an item, remove an item, get a single or all item from the data structure
- You should implement the ordered map by yourself and not import a go package that implements it
- ordered map data structure allow the server to get, add, and delete an object with O(1) and also scan all items based on the order they were inserted
- The server should read messages from the external queue while execute them in parallel. The execution of the commands should be executed in parallel as much as possible. For example if the server has 100 getItem, getAllItems requests they should not block each other and should be executed in parallel

### Client

- Should be configured from a command line or from a file (you decide)
- Can read data from a file or from a command line (you decide)
- Sends the messages to external queue
- All data is in the form of strings
- Clients can be run in parallel
- All keys and values are in the form of strings
- Clients can be added / removed / started while not inteferring to the server or other clients

### External Queue

- Can be Amazon Simple Queue Service (SQS) or RabbitMQ (you decide);

### Client and Server messages

- The messages represents commands the server should execute
    - addItem(’key’, ‘val’) - the client define the key and value (both are strings). The server stores the key and value in an ordered map data structure.
    - deleteItem(’key) - the client sends the key. The server removes the key from the ordered map data structure
    - getItem(’key’) - the client sends the key. The server get the key O(1) and print the key and its value to a file
    - getAllItems() - the client sends the commands. The server get all items based on the insert order and print the key and the value to a file

## Submission Guidelines

- Repository: Submit the project in GitHub repository
- Code: Should be compiled on arm64 or using docker on linuxamd64
- Assumptions: Document any assumptions made during the development process.

## Tips

- Break the project into small packages
- Define interfaces to the packages
- Have unit test (sample)
- Design it so it can the server can scale very quickly and serves many clients
- Think on parallelism, how to reduce contention between tasks
