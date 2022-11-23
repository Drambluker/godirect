package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type TLS struct {
	CertFile string `json:"certFile"`
	KeyFile  string `json:"keyFile"`
}

type Config struct {
	Scheme      string            `json:"schema"`
	Host        string            `json:"host"`
	Port        uint              `json:"port"`
	SecurePort  uint              `json:"securePort"`
	ContextPath string            `json:"contextPath"`
	TLS         *TLS              `json:"tls"`
	Rules       map[string]string `json:"rules"`
}

func NewConfig(filepath string) *Config {
	reader, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()

	jsonData, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Configuration:\n%s", string(jsonData))

	var config Config
	json.Unmarshal(jsonData, &config)
	return &config
}
