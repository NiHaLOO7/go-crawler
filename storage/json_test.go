package storage

import (
	"os"
	"testing"
)

func TestSaveJSON(t *testing.T) {
	os.MkdirAll("data", 0755)
	defer os.RemoveAll("data")

	type TestPage struct {
		URL   string
		Links []string
	}

	data := []TestPage{
		{URL: "https://test.com", Links: []string{"https://a.com"}},
	}

	err := SaveJSON("https://www.test.com/path", data)
	if err != nil {
		t.Fatalf("SaveJSON failed: %v", err)
	}

	if _, err := os.Stat("data/test.json"); os.IsNotExist(err) {
		t.Error("expected data/test.json to be created")
	}
}
