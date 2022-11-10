package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: loadSources(),
		ParseFunc: extractArticles,
		Exporters: []export.Exporter{&export.JSONLine{FileName: "out.jsonl"}},
	}).Start()
}

func loadSources() []string {
	var sources []string

	f, err := os.OpenFile("sources.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		s := sc.Text()
		if len(s) == 0 {
			continue
		}

		sources = append(sources, s)
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}

	return sources
}

func extractArticles(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("tr").Each(func(i int, s *goquery.Selection) {

		s.Find("td").Each(func(i int, s *goquery.Selection) {
			article := s.Find("a")
			url, exists := article.Attr("href")
			if !exists {
				return
			}

			if !strings.Contains(url, "/handle/") {
				return
			}

			name := article.Text()
			if len(name) == 0 {
				return
			}

			var fullURL string

			u, err := r.Request.URL.Parse(url)
			if err != nil {
				fullURL = url
			} else {
				fullURL = u.String()
			}

			g.Exports <- map[string]interface{}{
				"text": name,
				"url":  fullURL,
			}
		})

	})
}
