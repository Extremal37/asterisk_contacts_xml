package entity

import "encoding/xml"

type ContactDB struct {
	number string `db:"name"`
	name   string `db:"callerid"`
}

type ContactXML struct {
	XMLName      xml.Name `xml:"contact"`
	Name         string   `xml:"name,attr"`
	Bandwidth    uint8    `xml:"bandwidth,attr"`
	Group        uint8    `xml:"group,attr"`
	Favorite     uint8    `xml:"favorite,attr"`
	Sticky       uint8    `xml:"sticky,attr"`
	Favoritetype uint8    `xml:"favoritetype,attr"`
	Number       string   `xml:"number"`
}
