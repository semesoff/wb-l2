package main

import (
	"os"
	"strings"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	url := "http://example.com"
	filepath := "example.html"
	if err := downloadFile(url, filepath); err != nil {
		t.Errorf("downloadFile failed: %v", err)
	}
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		t.Errorf("file not found: %s", filepath)
	}
	os.Remove(filepath)
}

func TestExtractLinks(t *testing.T) {
	html := `<html><body><a href="/link1">Link1</a><img src="/image1.jpg"></body></html>`
	links, err := extractLinks("http://example.com", strings.NewReader(html))
	if err != nil {
		t.Errorf("extractLinks failed: %v", err)
	}
	expected := []string{"http://example.com/link1", "http://example.com/image1.jpg"}
	for i, link := range links {
		if link != expected[i] {
			t.Errorf("expected %s, got %s", expected[i], link)
		}
	}
}

func TestDownloadPage(t *testing.T) {
	url := "http://example.com"
	visited := make(map[string]bool)
	if err := downloadPage(url, visited); err != nil {
		t.Errorf("downloadPage failed: %v", err)
	}
	if _, err := os.Stat("example.com"); os.IsNotExist(err) {
		t.Errorf("directory not found: example.com")
	}
	os.RemoveAll("example.com")
}
