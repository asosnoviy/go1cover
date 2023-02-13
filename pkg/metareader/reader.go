package metareader

import (
	"encoding/xml"
	"log"
	"os"
	"path"
)

type Metareader struct {
	configuration *configuration
	srcpath       string
	coverData     coverData
}

func New(src string) *Metareader {

	metaReader := Metareader{
		configuration: &configuration{},
		srcpath:       src,
		coverData:     coverData{Data: make(map[Module]string)},
	}

	return &metaReader
}

func (m *Metareader) Parse() {

	m.configuration.unmarshalConfig(m.srcpath)

	m.coverData.Data[Module{ModuleUuid: m.configuration.Configuration.Uuid, ModuleType: SessionModule}] = path.Join(m.srcpath, "Ext", "SessionModule.bsl")
	m.coverData.Data[Module{ModuleUuid: m.configuration.Configuration.Uuid, ModuleType: ExternalConnectionModule}] = path.Join(m.srcpath, "Ext", "ExternalConnectionModule.bsl")
	m.coverData.Data[Module{ModuleUuid: m.configuration.Configuration.Uuid, ModuleType: OrdinaryApplicationModule}] = path.Join(m.srcpath, "Ext", "OrdinaryApplicationModule.bsl")
	m.coverData.Data[Module{ModuleUuid: m.configuration.Configuration.Uuid, ModuleType: ManagedApplicationModule}] = path.Join(m.srcpath, "Ext", "ManagedApplicationModule.bsl")

	m.fillID("AccountingRegisters", m.configuration.Configuration.ChildObjects.AccountingRegister)
	m.fillID("AccumulationRegisters", m.configuration.Configuration.ChildObjects.AccumulationRegister)
	m.fillID("BusinessProcesses", m.configuration.Configuration.ChildObjects.BusinessProcess)
	m.fillID("CalculationRegisters", m.configuration.Configuration.ChildObjects.CalculationRegister)
	m.fillID("Catalogs", m.configuration.Configuration.ChildObjects.Catalog)
	m.fillID("ChartsOfAccounts", m.configuration.Configuration.ChildObjects.ChartOfAccounts)
	m.fillID("ChartsOfCalculationTypes", m.configuration.Configuration.ChildObjects.ChartOfCalculationTypes)
	m.fillID("ChartsOfCharacteristicTypes", m.configuration.Configuration.ChildObjects.ChartOfCharacteristicTypes)
	m.fillID("CommonCommands", m.configuration.Configuration.ChildObjects.CommonCommand)
	m.fillID("CommonForms", m.configuration.Configuration.ChildObjects.CommonForm)
	m.fillID("CommonModules", m.configuration.Configuration.ChildObjects.CommonModule)
	m.fillID("Constants", m.configuration.Configuration.ChildObjects.Constant)
	m.fillID("DataProcessors", m.configuration.Configuration.ChildObjects.DataProcessor)
	m.fillID("DocumentJournals", m.configuration.Configuration.ChildObjects.DocumentJournal)
	m.fillID("Documents", m.configuration.Configuration.ChildObjects.Document)
	m.fillID("Enums", m.configuration.Configuration.ChildObjects.Enum)
	m.fillID("ExchangePlans", m.configuration.Configuration.ChildObjects.ExchangePlan)
	m.fillID("ExternalDataSources", m.configuration.Configuration.ChildObjects.ExternalDataSource)
	m.fillID("FilterCriteria", m.configuration.Configuration.ChildObjects.FilterCriterion)
	m.fillID("HTTPServices", m.configuration.Configuration.ChildObjects.HTTPService)
	m.fillID("InformationRegisters", m.configuration.Configuration.ChildObjects.InformationRegister)
	m.fillID("Reports", m.configuration.Configuration.ChildObjects.Report)
	m.fillID("SettingsStorages", m.configuration.Configuration.ChildObjects.SettingsStorage)
	m.fillID("Tasks", m.configuration.Configuration.ChildObjects.Task)
	m.fillID("WebServices", m.configuration.Configuration.ChildObjects.WebService)

}

func (m *Metareader) CoverData() *coverData {

	return &m.coverData
}

func (m *Metareader) fillID(typeName string, objectNames []string) {

	for _, v := range objectNames {
		filename := path.Join(m.srcpath, typeName, v+".xml")
		file, err := os.ReadFile(filename)
		if err != nil {
			log.Println("file not fund", err)
		}

		switch typeName {
		case "AccountingRegisters":
			metaData := AccountingRegister{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ManagerModule.bsl")
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: RecordSetModule}] = path.Join(m.srcpath, typeName, v, "Ext", "RecordSetModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "AccumulationRegisters":
			metaData := AccumulationRegister{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ManagerModule.bsl")
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: RecordSetModule}] = path.Join(m.srcpath, typeName, v, "Ext", "RecordSetModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "CalculationRegisters":
			metaData := CalculationRegister{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ManagerModule.bsl")
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: RecordSetModule}] = path.Join(m.srcpath, typeName, v, "Ext", "RecordSetModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "InformationRegisters":
			metaData := InformationRegister{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ManagerModule.bsl")
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: RecordSetModule}] = path.Join(m.srcpath, typeName, v, "Ext", "RecordSetModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "Documents":
			metaData := Document{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ManagerModule.bsl")
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ObjectModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ObjectModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "ChartsOfAccounts":
			metaData := ChartOfAccounts{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ManagerModule.bsl")
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ObjectModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ObjectModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "ChartsOfCalculationTypes":
			metaData := ChartOfCalculationTypes{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ManagerModule.bsl")
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ObjectModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ObjectModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "ChartsOfCharacteristicTypes":
			metaData := ChartOfCharacteristicTypes{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ManagerModule.bsl")
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ObjectModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ObjectModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "BusinessProcesses":
			metaData := BusinessProcess{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ManagerModule.bsl")
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ObjectModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ObjectModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "Tasks":
			metaData := Task{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ManagerModule.bsl")
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ObjectModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ObjectModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "Catalogs":
			metaData := Catalog{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ManagerModule.bsl")
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ObjectModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ObjectModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "Reports":
			metaData := Report{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ManagerModule.bsl")
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ObjectModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ObjectModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "DataProcessors":
			metaData := DataProcessor{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ManagerModule.bsl")
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ObjectModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ObjectModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "CommonModules":
			metaData := CommonModuleObj{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: SimpleModule}] = path.Join(m.srcpath, typeName, v, "Module.bsl")
		case "WebServices":
			metaData := WebService{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: SimpleModule}] = path.Join(m.srcpath, typeName, v, "Module.bsl")
		case "HTTPServices":
			metaData := HTTPService{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: SimpleModule}] = path.Join(m.srcpath, typeName, v, "Module.bsl")
		case "FilterCriteria":
			metaData := FilterCriterion{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "ManagerModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "SettingsStorages":
			metaData := SettingsStorage{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModuleStorage}] = path.Join(m.srcpath, typeName, v, "ManagerModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
		case "ExchangePlans":
			metaData := ExchangePlan{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ManagerModule.bsl")
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ObjectModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ObjectModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "DocumentJournals":
			metaData := DocumentJournal{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "Ext", "ManagerModule.bsl")

			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "CommonCommands":
			metaData := CommonCommand{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: CommandModule}] = path.Join(m.srcpath, typeName, v, "CommandModule.bsl")
		case "Enums":
			metaData := Enum{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "ManagerModule.bsl")
			path := path.Join(m.srcpath, typeName, v)
			fillform(m, path, metaData.Data.ChildObjects.Form)
			fillcommand(m, path, metaData.Data.ChildObjects.Command)
		case "Constants":
			metaData := Constant{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ManagerModule}] = path.Join(m.srcpath, typeName, v, "ManagerModule.bsl")
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: ValueManagerModule}] = path.Join(m.srcpath, typeName, v, "ValueManagerModule.bsl")

		case "CommonForms":
			metaData := CommonForm{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: FormModule}] = path.Join(m.srcpath, typeName, v, "Form", "Module.bsl")

		case "ExternalDataSources":
			metaData := ExternalDataSource{}
			xml.Unmarshal(file, &metaData)

			// todo

		default:
			panic("unknown metadatatype " + typeName)
		}

	}
}

func fillform(m *Metareader, srcpath string, formNames []string) {

	for _, formName := range formNames {
		filename := path.Join(srcpath, "Forms", formName+".xml")
		file, err := os.ReadFile(filename)
		if err != nil {
			log.Println("file not fund", err)
		}

		metaData := Form{}
		xml.Unmarshal(file, &metaData)
		m.coverData.Data[Module{ModuleUuid: metaData.Data.Uuid, ModuleType: FormModule}] = path.Join(srcpath, "Forms", formName, "Form", "Module.bsl")
	}
}

func fillcommand(m *Metareader, srcpath string, Commands []struct {
	Uuid       string "xml:\"uuid,attr\""
	Properties struct {
		Name string "xml:\"Name\""
	} "xml:\"Properties\""
}) {
	for _, c := range Commands {
		m.coverData.Data[Module{ModuleUuid: c.Uuid, ModuleType: CommandModule}] = path.Join(srcpath, "Commands", c.Properties.Name, "Ext", "CommandModule.bsl")
	}
}
