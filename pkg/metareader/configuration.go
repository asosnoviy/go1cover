package metareader

import (
	"encoding/xml"
	"log"
	"os"
	"path"
)

type configuration struct {
	Configuration struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			CommonModule               []string `xml:"CommonModule"`
			Catalog                    []string `xml:"Catalog"`
			Document                   []string `xml:"Document"`
			FilterCriterion            []string `xml:"FilterCriterion"`
			ExchangePlan               []string `xml:"ExchangePlan"`
			WebService                 []string `xml:"WebService"`
			HTTPService                []string `xml:"HTTPService"`
			SettingsStorage            []string `xml:"SettingsStorage"`
			CommonCommand              []string `xml:"CommonCommand"`
			Constant                   []string `xml:"Constant"`
			DocumentJournal            []string `xml:"DocumentJournal"`
			Enum                       []string `xml:"Enum"`
			Report                     []string `xml:"Report"`
			DataProcessor              []string `xml:"DataProcessor"`
			InformationRegister        []string `xml:"InformationRegister"`
			AccumulationRegister       []string `xml:"AccumulationRegister"`
			ChartOfCharacteristicTypes []string `xml:"ChartOfCharacteristicTypes"`
			ChartOfAccounts            []string `xml:"ChartOfAccounts"`
			AccountingRegister         []string `xml:"AccountingRegister"`
			ChartOfCalculationTypes    []string `xml:"ChartOfCalculationTypes"`
			CalculationRegister        []string `xml:"CalculationRegister"`
			BusinessProcess            []string `xml:"BusinessProcess"`
			Task                       []string `xml:"Task"`
			ExternalDataSource         []string `xml:"ExternalDataSource"`
			CommonForm                 []string `xml:"CommonForm"`
		} `xml:"ChildObjects"`
	} `xml:"Configuration"`
}

func (c *configuration) unmarshalConfig(configurationPatch string) {

	configurationFile, err := os.ReadFile(path.Join(configurationPatch, "Configuration.xml"))
	if err != nil {
		log.Fatal(err)
	}

	xml.Unmarshal(configurationFile, &c)

}
