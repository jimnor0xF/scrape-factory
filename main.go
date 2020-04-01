package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	//"golang.org/x/oauth2/google"
	//"golang.org/x/oauth2"
	"github.com/gocolly/colly"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func main() {
	//menuItems := make([]menuItem, 5)
	allMenuItems := []menuItem{}
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div.week-container", func(e *colly.HTMLElement) {
		e.ForEach("div.day", func(_ int, e *colly.HTMLElement) {
			tempMenuItem := menuItem{}
			var day string

			day = e.ChildText("h2")

			if day == "" {
				return
			}
			tempMenuItem.Day = day
			//fmt.Printf("%s\n", day)

			e.ForEach("div.menu-row", func(_ int, e *colly.HTMLElement) {
				var dish string
				var lunchCategory string
				dish = e.ChildText("div.element.description.col-md-4.col-print-5")
				lunchCategory = e.ChildText("div.element.title.col-md-4.col-print-3")

				if dish == "" {
					return
				}

				if lunchCategory == "" {
					return
			        }

				if lunchCategory != "The Factorys lunch" {
					tempMenuItem.Category = lunchCategory
				}

				if dish != "Välj din egen lunch från varma och kalla tillbehör och komplettera med proteinet" {
					tempMenuItem.Description = dish
				}

				if tempMenuItem.Day != "" && tempMenuItem.Category != "" && tempMenuItem.Description != "" {
					//fmt.Printf("%s\n", tempMenuItem.Day)
					//fmt.Printf("%s\n", tempMenuItem.Category)
					//fmt.Printf("%s\n", tempMenuItem.Description)
					//fmt.Printf("END MENU\n")
					allMenuItems = append(allMenuItems, tempMenuItem)
				}
			})
		})
		//TODO: Figure out how to make service account access calendar and make func that handles this. Calendar is empty for some reason.
		//var credFile = "./credentials.json"
		//cred, err := ioutil.ReadFile(credFile)

		//if err != nil {
		//	log.Fatalf("Unable to read JSON credentials config %v", err)
		//}

		//conf, err := google.JWTConfigFromJSON(cred, "https://www.googleapis.com/auth/calendar")

		//if err != nil {
		//	log.Fatal(err)
		//}

		//client := conf.Client(oauth2.NoContext)

		//srv, err := calendar.New(client)
		//if err != nil {
		//	log.Fatalf("Unable to retrieve Calender client: %v", err)

		//}
		//t := time.Now().Format(time.RFC3339)
		//events, err := srv.Events.List("primary").ShowDeleted(false).
		//	SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
		//if err != nil {
		//	log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
		//}
		//fmt.Println("Upcoming events:")
		//if len(events.Items) == 0 {
		//	fmt.Println("No upcoming events found.")
		//} else {
		//	for _, item := range events.Items {
		//		date := item.Start.DateTime
		//		if date == "" {
		//			date = item.Start.Date
		//		}
		//		fmt.Printf("%v (%v)\n", item.Summary, date)
		//	}
		//}

	})

	c.Visit("https://ericsson.foodbycoor.se/the-factory/restaurangen/restaurangens-meny?active_week=0")
}

type menuItem struct {
	Day         string
	Category    string
	Description string
}
