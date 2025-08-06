# Go ClickHouse Parquet Loader

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
    clickhousePort = 9000
    clickhouseUser = "default"
    clickhousePassword = ""
    clickhouseDB = "default"
)
```

---

## License

MIT License © 2025 Mohamed Hussain S
