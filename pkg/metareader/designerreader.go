package metareader

import (
	"encoding/xml"
	"log"
	"os"
	"path"
	"path/filepath"
)

type DesignerReader struct {
	ReaderCommon
	configuration *configuration
}

type ReaderCommon struct {
	srcpath   string
	coverData coverData
}

func NewDesignerReader(src string) *DesignerReader {

	path, _ := filepath.Abs(src)
	metaReader := DesignerReader{
		configuration: &configuration{},
		ReaderCommon: ReaderCommon{
			srcpath:   path,
			coverData: coverData{Data: make(map[Module]string)},
		},
	}

	return &metaReader
}

func (m *DesignerReader) Parse() {

	m.configuration.unmarshalConfig(m.srcpath)

	m.coverData.Data[Module{ModuleUuid: m.configuration.Configuration.Uuid, ModuleType: SessionModule}] = path.Join(m.srcpath, "Ext", "SessionModule.bsl")
	m.coverData.Data[Module{ModuleUuid: m.configuration.Configuration.Uuid, ModuleType: ExternalConnectionModule}] = path.Join(m.srcpath, "Ext", "ExternalConnectionModule.bsl")
	m.coverData.Data[Module{ModuleUuid: m.configuration.Configuration.Uuid, ModuleType: OrdinaryApplicationModule}] = path.Join(m.srcpath, "Ext", "OrdinaryApplicationModule.bsl")
	m.coverData.Data[Module{ModuleUuid: m.configuration.Configuration.Uuid, ModuleType: ManagedApplicationModule}] = path.Join(m.srcpath, "Ext", "ManagedApplicationModule.bsl")

	m.fillObjects("AccountingRegisters", m.configuration.Configuration.ChildObjects.AccountingRegister, &AccountingRegister{}, ManagerModule, RecordSetModule)
	m.fillObjects("AccumulationRegisters", m.configuration.Configuration.ChildObjects.AccumulationRegister, &AccumulationRegister{}, ManagerModule, RecordSetModule)
	m.fillObjects("InformationRegisters", m.configuration.Configuration.ChildObjects.InformationRegister, &InformationRegister{}, ManagerModule, RecordSetModule)
	// m.fillID("CalculationRegisters", m.configuration.Configuration.ChildObjects.CalculationRegister, &CalculationRegister{}, ManagerModule, RecordSetModule)

	m.fillObjects("BusinessProcesses", m.configuration.Configuration.ChildObjects.BusinessProcess, &BusinessProcess{}, ManagerModule, ObjectModule)
	m.fillObjects("Catalogs", m.configuration.Configuration.ChildObjects.Catalog, &Catalog{}, ManagerModule, ObjectModule)
	m.fillObjects("Tasks", m.configuration.Configuration.ChildObjects.Task, &Task{}, ManagerModule, ObjectModule)
	m.fillObjects("Reports", m.configuration.Configuration.ChildObjects.Report, &Report{}, ManagerModule, ObjectModule)
	m.fillObjects("DataProcessors", m.configuration.Configuration.ChildObjects.DataProcessor, &DataProcessor{}, ManagerModule, ObjectModule)
	m.fillObjects("DocumentJournals", m.configuration.Configuration.ChildObjects.DocumentJournal, &DocumentJournal{}, ManagerModule, ObjectModule)
	m.fillObjects("Documents", m.configuration.Configuration.ChildObjects.Document, &Document{}, ManagerModule, ObjectModule)
	m.fillObjects("Enums", m.configuration.Configuration.ChildObjects.Enum, &Enum{}, ManagerModule, ObjectModule)
	m.fillObjects("ExchangePlans", m.configuration.Configuration.ChildObjects.ExchangePlan, &ExchangePlan{}, ManagerModule, ObjectModule)
	m.fillObjects("ChartsOfAccounts", m.configuration.Configuration.ChildObjects.ChartOfAccounts, &ChartOfAccounts{}, ManagerModule, ObjectModule)
	m.fillObjects("ChartsOfCalculationTypes", m.configuration.Configuration.ChildObjects.ChartOfCalculationTypes, &ChartOfCalculationTypes{}, ManagerModule, ObjectModule)
	m.fillObjects("ChartsOfCharacteristicTypes", m.configuration.Configuration.ChildObjects.ChartOfCharacteristicTypes, &ChartOfCharacteristicTypes{}, ManagerModule, ObjectModule)

	m.fillObjects("FilterCriteria", m.configuration.Configuration.ChildObjects.FilterCriterion, &FilterCriterion{}, ManagerModule)

	m.fillObjects("CommonModules", m.configuration.Configuration.ChildObjects.CommonModule, &CommonModuleObj{}, SimpleModule)
	m.fillObjects("WebServices", m.configuration.Configuration.ChildObjects.WebService, &WebService{}, SimpleModule)
	m.fillObjects("HTTPServices", m.configuration.Configuration.ChildObjects.HTTPService, &HTTPService{}, SimpleModule)

	m.fillObjects("SettingsStorages", m.configuration.Configuration.ChildObjects.SettingsStorage, &SettingsStorage{}, ManagerModuleStorage)

	m.fillObjects("Constants", m.configuration.Configuration.ChildObjects.Constant, &Constant{}, ValueManagerModule, ManagerModule)

	m.fillObjects("CommonCommands", m.configuration.Configuration.ChildObjects.CommonCommand, &CommonCommand{}, CommandModule)

	m.fillID("CalculationRegisters", m.configuration.Configuration.ChildObjects.CalculationRegister)
	m.fillID("CommonForms", m.configuration.Configuration.ChildObjects.CommonForm)
	m.fillID("ExternalDataSources", m.configuration.Configuration.ChildObjects.ExternalDataSource)

}

func (m *DesignerReader) Files() []string {

	var files []string
	for _, v := range m.coverData.Data {
		files = append(files, v)
	}
	return files
}

func (m *DesignerReader) CoverData() *coverData {
	return &m.coverData
}

func (m *DesignerReader) fillObjects(typeName string, objectNames []string, metaDataType metaDataInt, ModuleTypes ...string) {

	for _, objectName := range objectNames {

		metaFilename := path.Join(m.srcpath, typeName, objectName+".xml")

		metaDataType.ReadMetaFile(metaFilename)

		for _, ModuleType := range ModuleTypes {
			m.coverData.Data[Module{ModuleUuid: metaDataType.GetData().Uuid, ModuleType: ModuleType}] = path.Join(m.srcpath, typeName, objectName, "Ext", Filenames[ModuleType])
		}
		path := path.Join(m.srcpath, typeName, objectName)
		fillform(m, path, metaDataType.GetData().ChildObjects.Form)
		fillcommand(m, path, metaDataType.GetData().ChildObjects.Command)

	}
}

func (m *DesignerReader) fillID(typeName string, objectNames []string) {

	for _, v := range objectNames {
		filename := path.Join(m.srcpath, typeName, v+".xml")
		file, err := os.ReadFile(filename)
		if err != nil {
			log.Println("file not fund", err)
		}

		switch typeName {

		case "CalculationRegisters":
			metaData := CalculationRegister{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "Ext", ManagerModuleFile)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: RecordSetModule}] = path.Join(m.srcpath, typeName, v, "Ext", RecordSetModuleFile)

			// todo recalc
			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)

		case "CommonForms":
			metaData := CommonForm{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: FormModule}] = path.Join(m.srcpath, typeName, v, "Ext", "Form", ModuleFile)

		case "ExternalDataSources":
			metaData := ExternalDataSource{}
			xml.Unmarshal(file, &metaData)

		default:
			panic("unknown metadatatype " + typeName)
		}

	}
}

func fillform(m *DesignerReader, srcpath string, formNames []string) {

	for _, formName := range formNames {
		filename := path.Join(srcpath, "Forms", formName+".xml")
		file, err := os.ReadFile(filename)
		if err != nil {
			log.Println("file not fund", err)
		}

		metaData := Form{}
		xml.Unmarshal(file, &metaData)
		m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: FormModule}] = path.Join(srcpath, "Forms", formName, "Ext", "Form", ModuleFile)
	}
}

func fillcommand(m *DesignerReader, srcpath string, Commands []struct {
	Uuid       string "xml:\"uuid,attr\""
	Properties struct {
		Name string "xml:\"Name\""
	} "xml:\"Properties\""
}) {
	for _, c := range Commands {
		m.coverData.Data[Module{ModuleUuid: c.Uuid, ModuleType: CommandModule}] = path.Join(srcpath, "Commands", c.Properties.Name, "Ext", CommandModuleFile)
	}
}
