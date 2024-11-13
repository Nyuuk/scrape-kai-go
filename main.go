package main

import (
	"github.com/gocolly/colly"
	"fmt"
	"strings"
)

func main() {
	// Inisialisasi collector
    c := colly.NewCollector()

	// Variabel untuk menyimpan hasil scraping
    var keretaDitemukan bool
    var statusTiket string

    // Mencari elemen dengan nama kereta
    c.OnHTML("div.col-one div.name", func(e *colly.HTMLElement) {
        namaKereta := e.Text
        if strings.Contains(strings.ToUpper(namaKereta), "AIRLANGGA") {
            keretaDitemukan = true
            fmt.Println("Kereta ditemukan:", namaKereta)
        }
    })

    // Mengecek status tiket pada elemen dengan kelas sisa-kursi
    c.OnHTML("small.sisa-kursi", func(e *colly.HTMLElement) {
        statusTiket = e.Text
        fmt.Println("Status Tiket:", statusTiket)
    })

    // Menetapkan user-agent dan header lainnya
    c.OnRequest(func(r *colly.Request) {
        // Menambahkan header sesuai dengan request di curl
        r.Headers.Set("Referer", "https://booking.kai.id/search?origination=%2BkiJ843HDK8x5CQIFzAxcQ%3D%3D&destination=eTzVVt1Dij2khlCW%2Ff9NbA%3D%3D&tanggal=yyU3sXLRgq2%2FkAsmJHQ4mppIGVTyX6qFviVNfTUMJZc%3D&adult=0SoHWQxybcl0Od8oS%2F%2B3uA%3D%3D&infant=1%2BCBzaRdtY8Be5JgBh3LrA%3D%3D&book_type=")
        r.Headers.Set("Upgrade-Insecure-Requests", "1")
        r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36")
        r.Headers.Set("sec-ch-ua", `"Not)A;Brand";v="99", "Google Chrome";v="127", "Chromium";v="127"`)
        r.Headers.Set("sec-ch-ua-mobile", "?0")
        r.Headers.Set("sec-ch-ua-platform", `"Linux"`)
    })

    // Membuat request GET ke URL yang diinginkan
    url := "https://booking.kai.id/search?origination=%2BkiJ843HDK8x5CQIFzAxcQ%3D%3D&destination=eTzVVt1Dij2khlCW%2Ff9NbA%3D%3D&tanggal=yyU3sXLRgq2%2FkAsmJHQ4mppIGVTyX6qFviVNfTUMJZc%3D&adult=0SoHWQxybcl0Od8oS%2F%2B3uA%3D%3D&infant=1%2BCBzaRdtY8Be5JgBh3LrA%3D%3D&book_type="
    c.Visit(url)

	// Menampilkan hasil
    if keretaDitemukan {
        fmt.Println("Kereta Airlangga ditemukan.")
    } else {
        fmt.Println("Kereta Airlangga tidak ditemukan.")
    }

    fmt.Println("Status Tiket:", statusTiket)
}