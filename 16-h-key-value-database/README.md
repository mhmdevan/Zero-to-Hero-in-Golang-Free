# 🗄️ Simple Key-Value Database in Go

A lightweight, in-memory Key-Value Database inspired by Redis, built in Go. This simple project provides basic functionality for storing, retrieving, and deleting key-value pairs. Perfect for learning about in-memory databases and building custom data stores! 🚀

## Features

- **SET**: Store a value associated with a key.
- **GET**: Retrieve the value associated with a key.
- **DEL**: Delete a key-value pair.
- **EXIT**: Exit the program.
- **Thread Safety**: Uses `sync.RWMutex` to ensure safe concurrent access.

## Getting Started

### Prerequisites

- **Go**: Make sure you have [Go](https://golang.org/doc/install) installed on your system.