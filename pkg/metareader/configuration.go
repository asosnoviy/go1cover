package metareader

import (
	"encoding/xml"
	"os"
	"path"
)

type configuration struct {
	Configuration struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			CommonModule []string `xml:"CommonModule"`
			Document     []string `xml:"Document"`
		} `xml:"ChildObjects"`
	} `xml:"Configuration"`
}

func (c *configuration) unmarshalConfig(configurationPatch string) {

	configurationFile, err := os.ReadFile(path.Join(configurationPatch, "Configuration.xml"))
	if err != nil {
		panic(err)
	}

	xml.Unmarshal(configurationFile, &c)

}

type metaDataObjectDocument struct {
	Data struct {
		Uuid string `xml:"uuid,attr"`
	} `xml:"Document"`
}

type metaDataObjectCommonModule struct {
	Data struct {
		Uuid string `xml:"uuid,attr"`
	} `xml:"CommonModule"`
}

type coverData struct {
	Data map[Module]string
}

type Module struct {
	ModuleTUuid string
	ModuleType  string
}
