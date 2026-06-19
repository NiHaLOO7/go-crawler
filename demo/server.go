package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NiHaLOO7/go-crawler/crawler"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PageResponse struct {
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Total   int    `json:"total"`
	Posts   []Post `json:"posts"`
	Next    string `json:"next,omitempty"`
}

type CrawlRequest struct {
	URL     string `json:"url"`
	Depth   int    `json:"depth"`
	Workers int    `json:"workers"`
}

type CrawlResponse struct {
	URL         string        `json:"url"`
	Depth       int           `json:"depth"`
	Workers     int           `json:"workers"`
	TotalPages  int           `json:"total_pages"`
	Results     []crawler.Page `json:"results"`
}

var allPosts []Post

func init() {
	for i := 1; i <= 50; i++ {
		allPosts = append(allPosts, Post{
			ID:    i,
			Title: fmt.Sprintf("Post #%d", i),
			Body:  fmt.Sprintf("This is the body of post number %d.", i),
		})
	}
}

func crawlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	var req CrawlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}
	if req.Depth < 1 {
		req.Depth = 1
	}
	if req.Workers < 1 {
		req.Workers = 5
	}

	c := crawler.NewCrawler(req.Depth, req.Workers)
	c.Run(req.URL)

	resp := CrawlResponse{
		URL:        req.URL,
		Depth:      req.Depth,
		Workers:    req.Workers,
		TotalPages: len(c.Results),
		Results:    c.Results,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	perPage := 10
	total := len(allPosts)

	start := (page - 1) * perPage
	end := start + perPage
	if start >= total {
		start = total
	}
	if end > total {
		end = total
	}

	resp := PageResponse{
		Page:    page,
		PerPage: perPage,
		Total:   total,
		Posts:   allPosts[start:end],
	}

	if end < total {
		resp.Next = fmt.Sprintf("http://localhost:8080/api/posts?page=%d", page+1)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `<!DOCTYPE html>
<html>
<head><title>Go Crawler Demo</title></head>
<body>
<h1>Demo Server</h1>
<p>Paginated API for testing the crawler.</p>
<ul>
<li><a href="/api/posts?page=1">Posts Page 1</a></li>
<li><a href="/api/posts?page=2">Posts Page 2</a></li>
<li><a href="/api/posts?page=3">Posts Page 3</a></li>
<li><a href="/api/posts?page=4">Posts Page 4</a></li>
<li><a href="/api/posts?page=5">Posts Page 5</a></li>
</ul>
</body>
</html>`
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/posts", postsHandler)
	http.HandleFunc("/crawl", crawlHandler)

	fmt.Println("Demo server running on http://localhost:8080")
	fmt.Println("  HTML index:     http://localhost:8080/")
	fmt.Println("  Paginated API:  http://localhost:8080/api/posts?page=1")
	fmt.Println("  Crawl API:      POST http://localhost:8080/crawl")
	fmt.Println(`    Body: {"url": "https://example.com", "depth": 2, "workers": 10}`)
	http.ListenAndServe(":8080", nil)
}
