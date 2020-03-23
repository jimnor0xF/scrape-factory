package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div.week-container", func(e *colly.HTMLElement) {
		e.ForEach("div.day", func(_ int, e *colly.HTMLElement) {
			var day string

			day = e.ChildText("h2")

			if day == "" {
				return
			}

			e.ForEach("div.menu-row", func(_ int, e *colly.HTMLElement) {
				var tempDish string
				var templunchCategory string
				tempDish = e.ChildText("div.element.description.col-md-4.col-print-5")
				templunchCategory = e.ChildText("div.element.title.col-md-4.col-print-3")

				if tempDish == "" {
					return
				}

				if templunchCategory != "The Factorys lunch" {
					fmt.Printf("%s\n", templunchCategory)
				}

				if tempDish != "Välj din egen lunch från varma och kalla tillbehör och komplettera med proteinet" {
					fmt.Printf("%s\n", tempDish)
				}

			})
		})
	})

	c.Visit("https://ericsson.foodbycoor.se/the-factory/restaurangen/restaurangens-meny?active_week=0")
}

type menuItem struct {
	Day         string
	Category    string
	Description string
}
