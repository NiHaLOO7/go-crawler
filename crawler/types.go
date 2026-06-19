package crawler
import (
	"sync"
)

type Page struct {
	URL string
	Links []string
	Depth int
}

type Crawler struct {
	MaxDepth int
	Workers int
	Visited map[string]bool
	Results []Page
	mu sync.Mutex
	wg sync.WaitGroup
	sem chan struct{}
}