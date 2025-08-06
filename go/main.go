package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	ck "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go-source/local"
)

type Row struct {
	ID    int64   `parquet:"name=id, type=INT64"`
	Name  string  `parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Score float64 `parquet:"name=score, type=DOUBLE"`
}

func main() {
	// Locate parquet file
	parquetPath := filepath.Join("/home/hussain/go-clickhouse-parquet/parquet-files", "sample.parquet")
	if _, err := os.Stat(parquetPath); os.IsNotExist(err) {
		log.Fatalf("Parquet file not found at %s. Run `make generate` first.", parquetPath)
	}

	// Read parquet
	fr, err := local.NewLocalFileReader(parquetPath)
	if err != nil {
		log.Fatal("Failed to open parquet:", err)
	}
	pr, err := reader.NewParquetReader(fr, new(Row), 4)
	if err != nil {
		log.Fatal("Failed to create parquet reader:", err)
	}
	defer pr.ReadStop()
	defer fr.Close()

	num := int(pr.GetNumRows())
	rows := make([]Row, num)
	if err = pr.Read(&rows); err != nil {
		log.Fatal("Failed to read parquet data:", err)
	}

	fmt.Printf("Read %d rows from parquet\n", num)

	// Insert into ClickHouse
	insertData(rows)
}

func insertData(rows []Row) {
	ctx := context.Background()
	conn, err := ck.Open(&ck.Options{
		Addr: []string{"localhost:9000"},
		Auth: ck.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
	})
	if err != nil {
		log.Fatal("ClickHouse connection failed:", err)
	}

	// Ensure table exists
	createTable := `
	CREATE TABLE IF NOT EXISTS scores (
		id Int64,
		name String,
		score Float64
	) ENGINE = MergeTree() ORDER BY id;
	`
	if err := conn.Exec(ctx, createTable); err != nil {
		log.Fatal("Table creation failed:", err)
	}

	batch, err := conn.PrepareBatch(ctx, "INSERT INTO scores (id, name, score)")
	if err != nil {
		log.Fatal("Batch prep failed:", err)
	}
	for _, r := range rows {
		if err := batch.Append(r.ID, r.Name, r.Score); err != nil {
			log.Fatal("Append failed:", err)
		}
	}
	if err := batch.Send(); err != nil {
		log.Fatal("Batch send failed:", err)
	}

	fmt.Println("Data inserted into ClickHouse successfully!")
}
