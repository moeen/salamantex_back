package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Postgres struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	DB   string `yaml:"db"`
}

type Redis struct {
	Address string `yaml:"address"`
	TxQueue string `yaml:"tx_queue"`
}

type Config struct {
	JWTSecret string   `yaml:"jwt_secret"`
	Postgres  Postgres `yaml:"postgres"`
	Redis     Redis    `yaml:"redis"`
}

var config *Config

func SetConfig() {
	source, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(source, &config)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetConfig() *Config {
	return config
}
