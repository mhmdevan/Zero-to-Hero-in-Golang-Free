# File Upload and Download Service 📁

A simple web server written in Go for uploading and downloading files. This project demonstrates file handling in Go by allowing users to upload a file via a web form and download it from the server.

## Features
- **File Upload**: Upload files through a web form with a file size limit of 10 MB.
- **File Download**: Download the uploaded file directly from the server.
- **HTML Interface**: Simple HTML form for uploading files.

## Prerequisites
- **Go** installed on your system. [Download Go](https://golang.org/dl/)

## Usage

### 1. Clone the Repository
Clone or create the `upload.go` and `upload.html` files as shown above.

### 2. Set Up the Directory Structure
Ensure the `uploads` directory exists to store uploaded files:
```bash
mkdir uploads