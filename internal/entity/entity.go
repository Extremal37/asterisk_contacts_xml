package entity

import "encoding/xml"

type ContactDB struct {
	Number string `db:"name"`
	Name   string `db:"callerid"`
}

type ContactXML struct {
	XMLName      xml.Name        `xml:"contact"`
	Name         string          `xml:"name,attr"`
	Numbers      []ContactNumber `xml:"number"`
	Bandwidth    uint8           `xml:"bandwidth,attr"`
	Group        uint8           `xml:"group,attr"`
	Favorite     uint8           `xml:"favorite,attr"`
	Sticky       uint8           `xml:"sticky,attr"`
	Favoritetype uint8           `xml:"favoritetype,attr"`
}

type ContactNumber struct {
	Value string `xml:"value,attr"`
}

type ContactList struct {
	XMLName  xml.Name     `xml:"contactList"`
	Contacts []ContactXML `xml:"contact"`
}
