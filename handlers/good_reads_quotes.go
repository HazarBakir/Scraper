package handlers

import (
	"github.com/gocolly/colly"
    // "github.com/HazarBakir/ScraperTool/modules"
	"fmt"
	"strings"
)

func GetQuotes() colly.HTMLCallback {
    return func(e *colly.HTMLElement) {
		rawQuoteText := e.ChildText("div.quoteText")
        quoteText := strings.TrimSpace(strings.Split(rawQuoteText, "―")[0])

        author := strings.TrimSpace(e.ChildText("span.authorOrTitle"))
        if quoteText != "" && author != "" {
            fmt.Printf("Quote: %s\nAuthor: %s\n----------------------------------------\n", quoteText, author)
        }
    }
}

