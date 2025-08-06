# ClickHouse Parquet Loader

A simple project to:

1. Generate sample data and export it as a Parquet file using Python.
2. Read the Parquet file using Go and insert the data into ClickHouse running inside Docker.

---

## 📂 Project Structure

```
clickhouse-parquet-loader/
│
├── docker-compose.yml        # Runs ClickHouse in Docker
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
