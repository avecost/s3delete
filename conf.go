package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port       int        `json:"port"`
	Bucket     Bucket     `json:"bucket"`
	Postgresql Postgresql `json:"postgresql"`
}

type Bucket struct {
	Endpoint  string `json:"endpoint"`
	Region    string `json:"region"`
	Name      string `json:"name"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

type Postgresql struct {
	Database string `json:"database"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func GetConfig(configFile string) (*Config, error) {

	var c Config

	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil

}
