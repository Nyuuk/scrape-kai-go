package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Nyuuk/scrape-kai-go/helpers"
)

func main() {
    countLoop := 0
    for {
        status := run()
        if status {
            fmt.Println("Ketemu kereta yang dicari, count:", countLoop)
            break
        }
        log.Println("Loop count bertambah", countLoop)
        countLoop++
        time.Sleep(time.Second * 10)
    }
}

func run() bool {
    config,_ := helpers.LoadConfig("config.yaml")
    kereta := config.TargetKereta
    var hasilKereta, statusTiket string
	// helpers.GetCookie()
    // data, err := helpers.GetDataKereta("Airlangga")
    var data []helpers.DataKereta
    var errorLoadDataKereta error
    loopCount := 0
    maxLoop := 5

    for loopCount < maxLoop {
        datas, err := helpers.GetDataKereta(kereta)
        if len(datas) > 0 {
            data = datas
            break
        }
        log.Println(err)
        loopCount++
        // log.Println("Loop count bertambah", loopCount)
        errorLoadDataKereta = err
    }

    // fmt.Println("data:", data)
    // fmt.Println("loopCount:", loopCount)

    for _, d := range data {
        // fmt.Println("d:", d)
        if strings.Contains(strings.ToUpper(d.NameKereta), strings.ToUpper(kereta)) {
            log.Println("Kereta ditemukan:", d.NameKereta, "Status tiket:", d.StatusTiket)
            hasilKereta = d.NameKeretaGet
            statusTiket = d.StatusTiket
            break
        }
    }

    // log.Println(len(data), loopCount == maxLoop, loopCount, maxLoop)
    if len(data) == 0 && loopCount == maxLoop {
        // sending message to admin
        log.Println("Error get data kereta", errorLoadDataKereta.Error())
        helpers.SendNotif(config.NumberAdmin, errorLoadDataKereta.Error())
        return false
    }

    if statusTiket != "Habis" && statusTiket == "Tersedia" {
        // log.Println("Status tiket:", statusTiket, "Kereta:", kereta)
        notifMsg := fmt.Sprintf("Laporan kereta untuk %s: \n\nNama Kereta: *%s* \nStatus Tiket: *%s*\n\nSilahkan melakukan pemesanan segera takutnya terbooking sama yang lain ya *%s*", kereta, hasilKereta, statusTiket, config.Name)
        helpers.SendNotif(config.Number, notifMsg)
        helpers.SendNotif(config.NumberAdmin, notifMsg)
        return true
    }
    return false
}