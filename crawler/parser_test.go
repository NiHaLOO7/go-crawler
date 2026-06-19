package crawler

import "testing"

func TestExtractLinks(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected []string
	}{
		{
			name:     "single link",
			html:     `<a href="https://example.com">Example</a>`,
			expected: []string{"https://example.com"},
		},
		{
			name:     "multiple links",
			html:     `<a href="https://a.com">A</a><a href="https://b.com">B</a>`,
			expected: []string{"https://a.com", "https://b.com"},
		},
		{
			name:     "no links",
			html:     `<p>No links here</p>`,
			expected: nil,
		},
		{
			name:     "relative and absolute mixed",
			html:     `<a href="/about">About</a><a href="https://x.com">X</a>`,
			expected: []string{"/about", "https://x.com"},
		},
		{
			name:     "empty href",
			html:     `<a href="">Empty</a>`,
			expected: []string{""},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ExtractLinks(tc.html)
			if len(got) != len(tc.expected) {
				t.Fatalf("expected %d links, got %d: %v", len(tc.expected), len(got), got)
			}
			for i := range got {
				if got[i] != tc.expected[i] {
					t.Errorf("link[%d]: expected %q, got %q", i, tc.expected[i], got[i])
				}
			}
		})
	}
}
