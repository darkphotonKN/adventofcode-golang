package concurrency

import "fmt"

type Webcrawler interface {
	CrawlSites([]string) (websites, error)
}

type websites map[string][]string
type webcrawler struct {
	siteUrls       []string
	crawledContent []websites
}

func (wc *webcrawler) CrawlSites([]string) (websites, error) {

	return nil, nil
}

func NewWebCrawler() Webcrawler {
	return &webcrawler{}
}

func RunCrawler() {
	fmt.Println("Test")
}
