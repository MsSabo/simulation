package config

import (
	"encoding/xml"
	"os"
)

// список обьектов симуляции
const (
	Grass string = "grass"
	Tree  string = "tree"
	Rock  string = "rock"
	Sheep string = "sheep"
	Fox   string = "fox"
)

type Config struct {
	XmlName xml.Name      `xml:"simulation"`
	MapSize MapSize       `xml:"mapSize"`
	Entity  []EntityParam `xml:"entity"`
}

type MapSize struct {
	Xlength int `xml:"xlength"`
	Ylength int `xml:"ylength"`
}
type EntityParam struct {
	Name     string `xml:"name"`
	Quantity int    `xml:"quantity"`
}

func ParseConfig(file string) (*Config, error) {
	cfg := new(Config)
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
