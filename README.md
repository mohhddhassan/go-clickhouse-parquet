# ClickHouse Parquet Loader

A simple project to:

1. Generate sample data and export it as a Parquet file using Python.
2. Read the Parquet file using Go and insert the data into ClickHouse running inside Docker.

---

## 📂 Project Structure

```
clickhouse-parquet-loader/
│
├── docker-compose.yml        # Runs ClickHouse in Docker# Go ClickHouse Parquet Loader

A lightweight utility built in Go to read data from Parquet files and insert it into [ClickHouse](https://clickhouse.com).

## Features

* Reads data from Parquet files using [parquet-go](https://github.com/xitongsys/parquet-go)
* Inserts data into ClickHouse via the official [ClickHouse Go driver](https://github.com/ClickHouse/clickhouse-go)
* Handles automatic table creation (if it doesn’t exist)
* Clean and minimal setup with Dockerized ClickHouse

---

## Project Structure

```
go-clickhouse-parquet/
├── go/                # Main Go code
│   ├── main.go        # Entry point
│   ├── model.go       # Structs mapped to Parquet
│   └── loader.go      # ClickHouse interaction
├── .gitignore
└── README.md
```

---

## Prerequisites

* [Go 1.20+](https://go.dev/dl/)
* [Docker](https://www.docker.com/) (for ClickHouse)
* Parquet file(s) to test with

---

## Setup

### **1. Clone Repo**

```bash
git clone https://github.com/<your-username>/go-clickhouse-parquet.git
cd go-clickhouse-parquet/go
```

### **2. Install Dependencies**

```bash
go mod tidy
```

### **3. Run ClickHouse with Docker**

```bash
docker run -d --name clickhouse-server \
    -p 9000:9000 -p 8123:8123 clickhouse/clickhouse-server:23.11
```

### **4. Run the App**

```bash
go run main.go
```

---

## Configuration

Modify database connection parameters in `main.go`:

```go
const (
    clickhouseHost = "localhost"
    clickhousePort = 8080
    clickhouseUser = "default"
    clickhousePassword = ""
    clickhouseDB = "default"
)
```

---

## License

MIT License © 2025 Mohamed Hussain S

│
├── python/                   # Python script for data generation
│   └── generate_parquet.py
│
├── go/                       # Go code to read parquet and insert into ClickHouse
│   ├── go.mod
│   ├── go.sum
│   └── main.go
│
└── parquet-files/            # Generated parquet files
    └── sample.parquet
```

---

## 🚀 How to Run

### 1. Start ClickHouse

```bash
docker-compose up -d
```

### 2. Generate Parquet File (Python)

```bash
cd python
python3 generate_parquet.py
```

### 3. Insert Data into ClickHouse (Go)

```bash
cd go
go run main.go
```

---

## 🛠 Tech Stack

* **Python** – Parquet file generation (`pandas`, `pyarrow`)
* **Go** – Reading parquet (`github.com/xitongsys/parquet-go`) & inserting into ClickHouse
* **ClickHouse** – Columnar database
* **Docker** – Containerized setup

---

## 🤝 Contributing

Feel free to fork this repo, improve it, and open pull requests.

---

## 📜 License

MIT License
