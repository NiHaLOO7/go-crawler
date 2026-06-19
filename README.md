# Go Crawler

A concurrent web crawler built from scratch in Go. Crawls websites up to a configurable depth using goroutines, with a semaphore-based worker pool for controlled concurrency.

## Features

- **Concurrent crawling** — goroutines with WaitGroup for synchronization
- **Worker pool** — semaphore pattern (buffered channel) to limit concurrent requests
- **Thread-safe** — Mutex-protected visited map and results slice
- **Configurable** — CLI flags for URL, depth, and worker count
- **JSON output** — crawl results saved to `data/` directory

## Architecture

```
main.go              CLI entry point (flags + orchestration)
crawler/
├── types.go         Page & Crawler structs
├── crawler.go       Core crawl logic (recursive, concurrent)
├── fetcher.go       HTTP fetcher (net/http)
└── parser.go        Link extractor (manual href parsing)
storage/
└── json.go          JSON file writer (saves to data/)
```

## Usage

```bash
# Default: crawl google.com, depth 1, 10 workers
go run main.go

# Custom URL with depth 2 and 20 workers
go run main.go -url https://example.com -depth 2 -workers 20
```

## Output

Crawl results are saved as JSON in the `data/` folder, named after the domain:

```json
[
  {
    "URL": "https://example.com",
    "Links": ["https://example.com/about", "https://example.com/contact"],
    "Depth": 0
  }
]
```

## Concurrency Model

```
main goroutine
  └── Run() spawns initial Crawl()
        └── For each link found:
              └── New goroutine (guarded by semaphore channel)
                    └── Recursive Crawl() at depth+1
```

- **Semaphore**: buffered channel of size `workers` — blocks when pool is full
- **Mutex**: protects `Visited` map and `Results` slice from race conditions
- **WaitGroup**: ensures main waits for all goroutines to complete

## Data Structures & Concepts

| Concept | Usage |
|---------|-------|
| Goroutines | One per URL to crawl |
| sync.Mutex | Protect shared state |
| sync.WaitGroup | Wait for all crawls to finish |
| Buffered Channel | Semaphore for worker limit |
| Map | Track visited URLs |
| Slice | Collect results |

## Built With

- Go standard library only (no external dependencies for core logic)
