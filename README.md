# Go API to Get Current Toronto Time and Store in MySQL

This Go project provides an API endpoint to retrieve the current time in Toronto and store the timestamp in a MySQL database. The API returns the time in JSON format and logs each request in a MySQL table.

## Features

- **API Endpoint**: `/current-time` that returns the current time in Toronto.
- **Timezone Conversion**: Uses Go's `time` package to convert UTC time to Toronto's local time (`America/Toronto`).
- **Database Logging**: Stores the timestamp of each API request in a MySQL database.

## Requirements

- **Go** (1.18+ recommended)
- **MySQL** Database
- **Go MySQL Driver** (`github.com/go-sql-driver/mysql`)

## Installation

### 1. Install Go

Make sure you have Go installed on your machine. You can download it from the official website: [https://golang.org/dl/](https://golang.org/dl/).

### 2. Install MySQL

You need to have a MySQL database running. You can download and install MySQL from [here](https://dev.mysql.com/downloads/installer/).

### 3. Set Up MySQL Database

1. **Create a new MySQL database** named `time_db`:

   ```sql
   CREATE DATABASE time_db;

2. **Create a table for storing the timestamp**   
   ```sql
   CREATE TABLE time_log (
    id INT AUTO_INCREMENT PRIMARY KEY,
    timestamp DATETIME
    );

### 4. Setup environment variables in the terminal of visual studio/IDE
    ```bash
    $env:DB_USER = "root"
    $env:DB_PASSWORD = Enter_yout_pasword_here
    $env:DB_NAME = "time_db"
    $env:DB_ADDRESS = "127.0.0.1"


### 5. Test the code

run main.go 

open browser localhost:8080/current-time
or

curl http://localhost:8080/current-time


now check the mysql table 




the timestamp is successfully stored in table 


