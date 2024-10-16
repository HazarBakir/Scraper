package main

import (
	"fmt"

	"github.com/HazarBakir/ScraperTool/handlers"
	// "github.com/HazarBakir/ScraperTool/modules"
	"github.com/gocolly/colly"
)

func main() {
    c := colly.NewCollector(
        colly.AllowedDomains("goodreads.com", "www.goodreads.com"),
    )

    c.OnHTML("div.quote", handlers.GetQuotes())

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting:", r.URL)
    })

    c.OnResponse(func(r *colly.Response) {
        fmt.Println("Received response with status:", r.StatusCode)
    })

    c.OnError(func(r *colly.Response, err error) {
        fmt.Println("Request failed with status:", r.StatusCode)
        fmt.Println("Error:", err)
    })

    c.Visit("https://www.goodreads.com/quotes")
}
