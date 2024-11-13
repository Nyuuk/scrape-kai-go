package helpers

import (
	"bytes"
	"log"
	"mime/multipart"
	"net/http"

)

func SendNotif(number string, msg string) string {
	config,_ := LoadConfig("config.yaml")
	url := config.ApiWhatsapp.Url

	var requestBody bytes.Buffer

	writer := multipart.NewWriter(&requestBody)
	writer.WriteField("number", number)
	writer.WriteField("text", msg)
	writer.Close()

	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		log.Println("error ketika sending notification new request", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("x-api-key", config.ApiWhatsapp.X_Api_Key)
	req.Header.Set("x-api-secret", config.ApiWhatsapp.X_Api_Secret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error ketika sending notification", err)
		return "Unkown"
	}
	defer resp.Body.Close()

	log.Println("response status notif:", resp.Status)

	return resp.Status
}