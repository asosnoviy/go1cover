package fdbc

import (
	"encoding/xml"
	"fmt"
)

type lineInfo struct {
	LineNo string `xml:"lineNo"`
}

type moduleID struct {
	ObjectID   string `xml:"objectID"`
	PropertyID string `xml:"propertyID"`
}

type moduleData struct {
	ModuleID moduleID   `xml:"moduleID"`
	LineInfo []lineInfo `xml:"lineInfo"`
}

type measure struct {
	ModuleData []moduleData `xml:"moduleData"`
}

type result struct {
	CmdID   string  `xml:"cmdID"`
	Measure measure `xml:"measure"`
}

type Response struct {
	XMLName xml.Name `xml:"response"`
	Result  []result `xml:"result"`
}

func (v *Response) Unmarshal(data []byte) {

	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

}
