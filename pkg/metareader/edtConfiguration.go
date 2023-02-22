package metareader

import (
	"encoding/xml"
	"log"
	"os"
	"path"
)

type EdtConfiguration struct {
	Uuid                        string   `xml:"uuid,attr"`
	CommonModules               []string `xml:"commonModules"`
	ExchangePlans               []string `xml:"exchangePlans"`
	FilterCriteria              []string `xml:"filterCriteria"`
	WebServices                 []string `xml:"webServices"`
	HttpServices                []string `xml:"httpServices"`
	CommonCommands              []string `xml:"commonCommands"`
	Constants                   []string `xml:"constants"`
	CommonForms                 []string `xml:"commonForms"`
	Catalogs                    []string `xml:"catalogs"`
	Documents                   []string `xml:"documents"`
	Enums                       []string `xml:"enums"`
	DocumentJournals            []string `xml:"documentJournals"`
	Reports                     []string `xml:"reports"`
	InformationRegisters        []string `xml:"informationRegisters"`
	AccumulationRegisters       []string `xml:"accumulationRegisters"`
	ChartsOfCharacteristicTypes []string `xml:"chartsOfCharacteristicTypes"`
	ChartsOfAccounts            []string `xml:"chartsOfAccounts"`
	AccountingRegisters         []string `xml:"accountingRegisters"`
	ChartsOfCalculationTypes    []string `xml:"chartsOfCalculationTypes"`
	CalculationRegisters        []string `xml:"calculationRegisters"`
	BusinessProcesses           []string `xml:"businessProcesses"`
	Tasks                       []string `xml:"tasks"`
	ExternalDataSources         []string `xml:"externalDataSources"`
}

func (c *EdtConfiguration) unmarshalConfig(configurationPatch string) {

	filepath := path.Join(configurationPatch, "Configuration", "Configuration.mdo")
	configurationFile, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	xml.Unmarshal(configurationFile, &c)

}

type edtData struct {
	Uuid           string           `xml:"uuid,attr"`
	Recalculations []Recalculations `xml:"recalculations"`
	Forms          []Forms          `xml:"forms"`
	Commands       []Commands       `xml:"commands"`
	Tables         []string         `xml:"tables"`
	Cubes          []string         `xml:"cubes"`
}

type Forms struct {
	Uuid string `xml:"uuid,attr"`
	Name string `xml:"name"`
}

type Recalculations struct {
	Uuid string `xml:"uuid,attr"`
	Name string `xml:"name"`
}

type Commands struct {
	Uuid string `xml:"uuid,attr"`
	Name string `xml:"name"`
}

type Tables struct {
	Uuid string `xml:"uuid,attr"`
	Name string `xml:"name"`
}

type Cubes struct {
	Uuid string `xml:"uuid,attr"`
	Name string `xml:"name"`
}

func (s *edtData) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

// type Data struct {
// 	Uuid         string `xml:"uuid,attr"`
// 	ChildObjects struct {
// 		Form    []string `xml:"Form"`
// 		Command []struct {
// 			Uuid       string `xml:"uuid,attr"`
// 			Properties struct {
// 				Name string `xml:"Name"`
// 			} `xml:"Properties"`
// 		}
// 	}
// }
