package metareader

import (
	"encoding/xml"
	"log"
	"os"
)

type metaDataInt interface {
	ReadMetaFile(string)
	GetData() Data
}

type Data struct {
	Uuid         string `xml:"uuid,attr"`
	ChildObjects struct {
		Form    []string `xml:"Form"`
		Command []struct {
			Uuid       string `xml:"uuid,attr"`
			Properties struct {
				Name string `xml:"Name"`
			} `xml:"Properties"`
		}
	}
}

type Document struct {
	Data Data `xml:"Document"`
}

func (s *Document) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *Document) GetData() Data {

	return s.Data
}

type BusinessProcess struct {
	Data Data `xml:"BusinessProcess"`
}

func (s *BusinessProcess) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *BusinessProcess) GetData() Data {

	return s.Data
}

type Catalog struct {
	Data Data `xml:"Catalog"`
}

func (s *Catalog) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *Catalog) GetData() Data {

	return s.Data
}

type Report struct {
	Data Data `xml:"Report"`
}

func (s *Report) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *Report) GetData() Data {

	return s.Data
}

type DataProcessor struct {
	Data Data `xml:"DataProcessor"`
}

func (s *DataProcessor) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *DataProcessor) GetData() Data {

	return s.Data
}

type InformationRegister struct {
	Data Data `xml:"InformationRegister"`
}

func (s *InformationRegister) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *InformationRegister) GetData() Data {

	return s.Data
}

type AccumulationRegister struct {
	Data Data `xml:"AccumulationRegister"`
}

func (s *AccumulationRegister) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *AccumulationRegister) GetData() Data {

	return s.Data
}

type AccountingRegister struct {
	Data Data `xml:"AccountingRegister"`
}

func (s *AccountingRegister) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *AccountingRegister) GetData() Data {

	return s.Data
}

type ChartOfCalculationTypes struct {
	Data Data `xml:"ChartOfCalculationTypes"`
}

func (s *ChartOfCalculationTypes) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *ChartOfCalculationTypes) GetData() Data {

	return s.Data
}

type Task struct {
	Data Data `xml:"Task"`
}

func (s *Task) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *Task) GetData() Data {

	return s.Data
}

type Cube struct {
	Data Data `xml:"Cube"`
}

func (s *Cube) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *Cube) GetData() Data {

	return s.Data
}

type ExternalDataSource struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Table []string `xml:"Table"`
			Cube  []string `xml:"Cube"`
		}
	} `xml:"ExternalDataSource"`
}

type CalculationRegister struct {
	Data struct {
		Uuid         string `xml:"uuid,attr"`
		ChildObjects struct {
			Form          []string `xml:"Form"`
			Recalculation []string `xml:"Recalculation"`
			Command       []struct {
				Uuid       string `xml:"uuid,attr"`
				Properties struct {
					Name string `xml:"Name"`
				} `xml:"Properties"`
			}
		}
	} `xml:"CalculationRegister"`
}

type ChartOfCharacteristicTypes struct {
	Data Data `xml:"ChartOfCharacteristicTypes"`
}

func (s *ChartOfCharacteristicTypes) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *ChartOfCharacteristicTypes) GetData() Data {

	return s.Data
}

type ChartOfAccounts struct {
	Data Data `xml:"ChartOfAccounts"`
}

func (s *ChartOfAccounts) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *ChartOfAccounts) GetData() Data {

	return s.Data
}

type ExchangePlan struct {
	Data Data `xml:"ExchangePlan"`
}

func (s *ExchangePlan) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *ExchangePlan) GetData() Data {

	return s.Data
}

type DocumentJournal struct {
	Data Data `xml:"DocumentJournal"`
}

func (s *DocumentJournal) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *DocumentJournal) GetData() Data {

	return s.Data
}

type Enum struct {
	Data Data `xml:"Enum"`
}

func (s *Enum) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *Enum) GetData() Data {

	return s.Data
}

type CommonModuleObj struct {
	Data Data `xml:"CommonModule"`
}

func (s *CommonModuleObj) ReadMetaFile(filename string) {

	unmarshal(filename, s)
}

func (s *CommonModuleObj) GetData() Data {

	return s.Data
}

type SettingsStorage struct {
	Data Data `xml:"SettingsStorage"`
}

func (s *SettingsStorage) ReadMetaFile(filename string) {
	unmarshal(filename, s)
}

func (s *SettingsStorage) GetData() Data {
	return s.Data
}

type Table struct {
	Data Data `xml:"Table"`
}

func (s *Table) ReadMetaFile(filename string) {
	unmarshal(filename, s)
}

func (s *Table) GetData() Data {
	return s.Data
}

type FilterCriterion struct {
	Data Data `xml:"FilterCriterion"`
}

func (s *FilterCriterion) ReadMetaFile(filename string) {
	unmarshal(filename, s)
}

func (s *FilterCriterion) GetData() Data {
	return s.Data
}

type WebService struct {
	Data Data `xml:"WebService"`
}

func (s *WebService) ReadMetaFile(filename string) {
	unmarshal(filename, s)
}

func (s *WebService) GetData() Data {
	return s.Data
}

type HTTPService struct {
	Data Data `xml:"HTTPService"`
}

func (s *HTTPService) ReadMetaFile(filename string) {
	unmarshal(filename, s)
}

func (s *HTTPService) GetData() Data {
	return s.Data
}

type Form struct {
	Data struct {
		Uuid string `xml:"uuid,attr"`
	} `xml:"Form"`
}

type CommonForm struct {
	Data struct {
		Uuid string `xml:"uuid,attr"`
	} `xml:"CommonForm"`
}

type CommonCommand struct {
	Data Data `xml:"CommonCommand"`
}

func (s *CommonCommand) ReadMetaFile(filename string) {
	unmarshal(filename, s)
}

func (s *CommonCommand) GetData() Data {
	return s.Data
}

type Constant struct {
	Data Data `xml:"Constant"`
}

func (s *Constant) ReadMetaFile(filename string) {
	unmarshal(filename, s)
}

func (s *Constant) GetData() Data {
	return s.Data
}
func unmarshal(filename string, s interface{}) {

	metafile, err := os.ReadFile(filename)
	if err != nil {
		log.Println("file not fund", err)

	}

	xml.Unmarshal(metafile, s)
}

const (
	CommandModule             string = "078a6af8-d22c-4248-9c33-7e90075a3d2c"
	ObjectModule              string = "a637f77f-3840-441d-a1c3-699c8c5cb7e0"
	ManagerModule             string = "d1b64a2c-8078-4982-8190-8f81aefda192"
	ManagerModuleStorage      string = "0c8cad23-bf8c-468e-b49e-12f1927c048b"
	FormModule                string = "32e087ab-1491-49b6-aba7-43571b41ac2b"
	RecordSetModule           string = "9f36fd70-4bf4-47f6-b235-935f73aab43f"
	ValueManagerModule        string = "3e58c91f-9aaa-4f42-8999-4baf33907b75"
	ManagedApplicationModule  string = "d22e852a-cf8a-4f77-8ccb-3548e7792bea"
	SessionModule             string = "9b7bbbae-9771-46f2-9e4d-2489e0ffc702"
	ExternalConnectionModule  string = "a4a9c1e2-1e54-4c7f-af06-4ca341198fac"
	OrdinaryApplicationModule string = "a78d9ce3-4e0c-48d5-9863-ae7342eedf94"
	SimpleModule              string = "d5963243-262e-4398-b4d7-fb16d06484f6"
)

const (
	ManagerModuleFile      string = "ManagerModule.bsl"
	RecordSetModuleFile    string = "RecordSetModule.bsl"
	ObjectModuleFile       string = "ObjectModule.bsl"
	ModuleFile             string = "Module.bsl"
	CommandModuleFile      string = "CommandModule.bsl"
	ValueManagerModuleFile string = "ValueManagerModule.bsl"
)

var Filenames map[string]string = map[string]string{
	ManagerModule:        ManagerModuleFile,
	RecordSetModule:      RecordSetModuleFile,
	ObjectModule:         ObjectModuleFile,
	SimpleModule:         ModuleFile,
	ManagerModuleStorage: ManagerModuleFile,
	ValueManagerModule:   ValueManagerModuleFile,
	CommandModule:        CommandModuleFile,
}

type coverData struct {
	Data map[Module]string
}

type Module struct {
	ModuleUuid string
	ModuleType string
}
