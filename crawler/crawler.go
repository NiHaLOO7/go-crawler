package crawler
import (
	// "fmt"
	"strings"
)

func NewCrawler(maxDepth int, workers int) *Crawler {
	return &Crawler{
		MaxDepth: maxDepth,
		Workers: workers,
		Visited: make(map[string]bool),
		sem: make(chan struct{}, workers),
	}
}

func (c *Crawler) Crawl(url string, depth int) *Page {
	c.mu.Lock()
	if c.Visited[url] || depth > c.MaxDepth {
		c.mu.Unlock()
		return nil
	}
	c.Visited[url] = true
	c.mu.Unlock()
	html, err := Fetch(url)
	if err != nil {
		return nil
	}
	links := ExtractLinks(html)
	// fmt.Println(links) 
	for _, link := range links {
		if strings.HasPrefix(link, "http") {
			// c.Crawl(link,depth+1)
			c.wg.Add(1)
			go func(l string) {
				c.sem <- struct{}{}
				defer func() { <-c.sem }() 
				defer c.wg.Done()
				c.Crawl(l, depth+1)
			}(link)
		}
	}
	// Crawl ke andar, return se pehle:
	page := Page{URL: url, Links: links, Depth: depth}
	c.mu.Lock()
	c.Results = append(c.Results, page)
	c.mu.Unlock()
	return &page
}

func (c  *Crawler) Run(url string) {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		c.Crawl(url, 0)
	}()
	c.wg.Wait()
}