package helpers

import (
	"gopkg.in/yaml.v3"
	"os"
)

type ConfigHelper struct {
	Name string `yaml:"name"`
	Number string `yaml:"number"`
	NumberAdmin string `yaml:"number_admin"`
	TargetKereta string `yaml:"target_kereta"`
	UrlBooking string `yaml:"url_booking"`
	ApiWhatsapp struct {
		X_Api_Key string `yaml:"x_api_key"`
		X_Api_Secret string `yaml:"x_api_secret"`
		Url string `yaml:"url"`
	} `yaml:"api_whatsapp"`
}

func LoadConfig(filename string) (*ConfigHelper, error) {
	config := &ConfigHelper{}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}