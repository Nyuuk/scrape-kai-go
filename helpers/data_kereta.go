package helpers

import (
	"github.com/gocolly/colly/v2"
)

type DataKereta struct {
	NameKereta string
	StatusTiket string
}

func GetDataKereta(nameKerata string) ([]DataKereta, error) {
	config,_ := LoadConfig("config.yaml")
	// Inisialisasi collector
    c := colly.NewCollector()

	data := []DataKereta{}

    // Mencari elemen dengan nama kereta dan mengecek status tiket jika ditemukan
    c.OnHTML("div.data-block.list-kereta", func(e *colly.HTMLElement) {
        // Mengecek nama kereta
		kereta := e.ChildText("div.col-one div.name")
		statusTiket := e.ChildText("small.sisa-kursi")

		data = append(data, DataKereta{
			NameKereta: kereta,
			StatusTiket: statusTiket,
		})
    })

    // Menetapkan user-agent dan header lainnya
    c.OnRequest(func(r *colly.Request) {
        // Menambahkan header sesuai dengan request di curl
        r.Headers.Set("Referer", config.UrlBooking)
        r.Headers.Set("Upgrade-Insecure-Requests", "1")
        r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36")
        r.Headers.Set("sec-ch-ua", `"Not)A;Brand";v="99", "Google Chrome";v="127", "Chromium";v="127"`)
        r.Headers.Set("sec-ch-ua-mobile", "?0")
        r.Headers.Set("sec-ch-ua-platform", `"Linux"`)
    })

    // Membuat request GET ke URL yang diinginkan
    url := config.UrlBooking
    err := c.Visit(url)

	if err != nil {
		return []DataKereta{}, err
	}
	return data, nil
}