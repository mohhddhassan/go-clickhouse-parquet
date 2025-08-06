import pandas as pd

def main():
    data = {
        "id": [1, 2, 3, 4, 5],
        "name": ["Alice", "Bob", "Charlie", "David", "Eva"],
        "score": [85.5, 90.0, 88.5, 76.0, 92.0]
    }
    df = pd.DataFrame(data)

    output_path = "/home/hussain/go-clickhouse-parquet/parquet-files/sample.parquet"
    df.to_parquet(output_path, engine="pyarrow", index=False)
    print(f"Parquet file generated at: {output_path}")

if __name__ == "__main__":
    main()
