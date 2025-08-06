# Go + Python Parquet to ClickHouse 🚀

This project demonstrates how to **generate Parquet files using Python** and **ingest them into ClickHouse** using a **Go application**.
It also comes with a **Docker Compose** setup for easy local ClickHouse deployment.

---

## 🐳 What’s Inside?

* **Python** script to generate sample Parquet data (`generate_parquet.py`)
* **Go** application to read the Parquet file and insert data into ClickHouse (`main.go`)
* **Docker Compose** file to spin up ClickHouse quickly
* A sample Parquet file (`sample.parquet`) for quick testing

---

## 🔧 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/mohhddhassan/go-clickhouse-parquet.git
cd go-clickhouse-parquet
```

### 2. Generate a sample Parquet file

```bash
cd python
python3 generate_parquet.py
```

### 3. Start ClickHouse using Docker Compose

```bash
docker-compose up -d
```

### 4. Run the Go app to insert data into ClickHouse

```bash
cd go
go run main.go
```

---

## 📂 File Structure

```
go-clickhouse-parquet/
├── docker-compose.yml
├── parquet-files/
│   └── sample.parquet
├── python/
│   └── generate_parquet.py
└── go/
    ├── go.mod
    ├── go.sum
    └── main.go
```

---

## 🧠 What I Learned

* Working with **Parquet files** programmatically
* Using **Go to interact with ClickHouse**
* Deploying **ClickHouse with Docker Compose**

---

## 📝 Notes

This is a **beginner-friendly experiment** for exploring Python (data generation), Go (ETL), and ClickHouse (OLAP database). Future improvements may include streaming pipelines and handling larger datasets.
