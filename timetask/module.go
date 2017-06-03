package timetask

import (
	"encoding/xml"
)

type Module struct {
	ID   int    `xml:"moduleid"`
	Name string `xml:"modulename"`
}

type moduleXML struct {
	Modules []Module `xml:"response>item"`
}

func ParseXML(xml_str string) ([]Module, error) {
	modules := moduleXML{}
	err := xml.Unmarshal([]byte(xml_str), &modules)
	if err != nil {
		return nil, err
	}

	return modules.Modules, nil
}
