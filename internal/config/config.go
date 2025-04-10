package config

import (
	"encoding/xml"
	"fmt"
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
	Entity  []EntityParam `xml:"entity"`
}

type EntityParam struct {
	Name     string `xml:"name"`
	Quantity int    `xml:"quantity"`
}

func ParseConfig(file string) (*Config, error) {
	cfg := new(Config)
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = xml.Unmarshal(data, cfg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("--- Unmarshal ---\n\n")
	fmt.Printf("%+v\n", cfg)

	return cfg, nil
}
