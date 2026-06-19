package main
import (
    "fmt"
    "github.com/NiHaLOO7/go-crawler/crawler"
	"github.com/NiHaLOO7/go-crawler/storage"
	"time"
	"flag"
)

func main() {
	start := time.Now()
	url := flag.String("url", "https://google.com", "URL to crawl")
	depth := flag.Int("depth", 1, "Max crawl depth")
	workers := flag.Int("workers", 10, "Number of concurrent workers")
	flag.Parse()
	c := crawler.NewCrawler(*depth, *workers)
	c.Run(*url)
	fmt.Println("Total pages crawled:", len(c.Results))
    for _, page := range c.Results {
        fmt.Println(page.URL, "- depth:", page.Depth, "- links:", len(page.Links))
	}
	fmt.Println("Time taken:", time.Since(start))
	storage.SaveJSON(*url, c.Results)
}