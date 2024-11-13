package helpers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly/v2"
)

func GetCookie() []*http.Cookie {
	c := colly.NewCollector()
	var cookies []*http.Cookie

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Requesting to get Cookie:", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		// tampilin cookie
		fmt.Println("Cookies:")
		cookies = c.Cookies(r.Request.URL.String())
		for _, cookie := range cookies {
			fmt.Printf("%s=%s\n", cookie.Name, cookie.Value)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit("https://booking.kai.id")
	if err != nil {
		log.Fatal(err)
	}
	return cookies
}