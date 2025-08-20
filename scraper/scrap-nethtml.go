package scraper

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func ScrapNethtml() {
	// URL pro scraping
	url := "https://www.theguardian.com/lifeandstyle/series/sudoku"
	
	// HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Chyba při HTTP requestu:", err)
	}
	defer resp.Body.Close()

	// Kontrola status kódu
	if resp.StatusCode != http.StatusOK {
		log.Fatal("HTTP status:", resp.StatusCode)
	}

	// Parsování HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal("Chyba při parsování HTML:", err)
	}

	// Extrakce odkazů s konkrétní CSS třídou
	links := extractLinks(doc)
	
	fmt.Println("Nalezené odkazy s class='dcr-2yd10d':")
	for _, link := range links {
		fmt.Println("-", link)
	}
}

// Rekurzivní funkce pro extrakci odkazů s konkrétní CSS třídou
func extractLinks(n *html.Node) []string {
	var links []string
	
	// Pokud je to <a> tag s konkrétní class a href atributem
	if n.Type == html.ElementNode && n.Data == "a" {
		var hasTargetClass bool
		var hrefValue string
		
		// Projdeme všechny atributy
		for _, attr := range n.Attr {
			if attr.Key == "class" && attr.Val == "dcr-2yd10d" {
				hasTargetClass = true
			}
			if attr.Key == "href" {
				hrefValue = attr.Val
			}
		}
		
		// Přidáme odkaz pouze pokud má správnou třídu a href
		if hasTargetClass && hrefValue != "" && !strings.HasPrefix(hrefValue, "#") {
			links = append(links, hrefValue)
		}
	}
	
	// Rekurzivně procházíme všechny child elementy
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		links = append(links, extractLinks(child)...)
	}
	
	return links
}
