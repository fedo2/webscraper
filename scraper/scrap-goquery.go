package scraper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func scrap_goquery() {
	url := "https://www.theguardian.com/lifeandstyle/series/sudoku"
	
	// HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Chyba při HTTP requestu:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("HTTP status:", resp.StatusCode)
	}

	// Vytvoří GoQuery dokument
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Chyba při parsování HTML:", err)
	}

	fmt.Println("=== GoQuery Scraping ===")
	
	// Najde všechny odkazy s konkrétní třídou pomocí CSS selektoru
	doc.Find("a.dcr-2yd10d").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			title := s.Text() // Text uvnitř <a> tagu
			fmt.Printf("%d. %s\n   -> %s\n", i+1, title, href)
		}
	})
	
	// Bonus: Najdeme také titulky článků
	fmt.Println("\n=== Titulky článků ===")
	doc.Find("h3").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		if title != "" {
			fmt.Printf("- %s\n", title)
		}
	})
}