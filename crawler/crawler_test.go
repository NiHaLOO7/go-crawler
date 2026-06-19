package crawler

import (
	"testing"
)

func TestNewCrawler(t *testing.T) {
	c := NewCrawler(3, 5)
	if c.MaxDepth != 3 {
		t.Errorf("expected MaxDepth 3, got %d", c.MaxDepth)
	}
	if c.Workers != 5 {
		t.Errorf("expected Workers 5, got %d", c.Workers)
	}
	if c.Visited == nil {
		t.Error("Visited map should be initialized")
	}
	if cap(c.sem) != 5 {
		t.Errorf("semaphore capacity should be 5, got %d", cap(c.sem))
	}
}

func TestCrawlSkipsVisited(t *testing.T) {
	c := NewCrawler(1, 2)
	c.mu.Lock()
	c.Visited["https://already-visited.com"] = true
	c.mu.Unlock()

	page := c.Crawl("https://already-visited.com", 0)
	if page != nil {
		t.Error("should return nil for already visited URL")
	}
}

func TestCrawlSkipsOverDepth(t *testing.T) {
	c := NewCrawler(2, 2)
	page := c.Crawl("https://example.com", 3)
	if page != nil {
		t.Error("should return nil when depth exceeds MaxDepth")
	}
}
