# crawler

A simple website crawler made with go

## Usage

1. Clone the repo
```bash
git clone https://github.com/dagregi/crawler
```

2. Compile the package
```bash
go build -o crawler
```

3. Run the executable
```bash
./crawler <baseURL> <maxConcurrency> <maxPages>
# example
./crawler https://crawler-test.com 5 30
```

**Warning**: Unless you want your IP blocked don't send large amounts of requests
