package config

import "encoding/xml"

type Config struct {
	XmlName xml.Name `xml:"simulation:config"`
	Amount  []struct {
		Grass int `xml:"Grass"`
		Tree  int `xml:"Tree"`
		Sheep int `xml:"Sheep"`
		Fox   int `xml:"Fox"`
		Rock  int `xml:"Rock"`
	} `xml:"amount"`
}
