package metareader

import (
	"path"
	"path/filepath"
	"strings"
)

type EdtReader struct {
	ReaderCommon
	configuration *EdtConfiguration
}

func NewEdtReader(src string) *EdtReader {
	//todo linux/windows
	path, _ := filepath.Abs(src)
	return &EdtReader{
		configuration: &EdtConfiguration{},
		ReaderCommon: ReaderCommon{
			srcpath:   path,
			coverData: coverData{Data: make(map[Module]string)},
		},
	}
}

func (m *EdtReader) Parse() {
	m.configuration.unmarshalConfig(m.srcpath)

	m.coverData.Data[Module{ModuleUuid: m.configuration.Uuid, ModuleType: SessionModule}] = path.Join(m.srcpath, "Configuration", Filenames[SessionModule])
	m.coverData.Data[Module{ModuleUuid: m.configuration.Uuid, ModuleType: ManagedApplicationModule}] = path.Join(m.srcpath, "Configuration", Filenames[ManagedApplicationModule])
	m.coverData.Data[Module{ModuleUuid: m.configuration.Uuid, ModuleType: OrdinaryApplicationModule}] = path.Join(m.srcpath, "Configuration", Filenames[OrdinaryApplicationModule])
	m.coverData.Data[Module{ModuleUuid: m.configuration.Uuid, ModuleType: ExternalConnectionModule}] = path.Join(m.srcpath, "Configuration", Filenames[ExternalConnectionModule])

	m.fillObjects("AccountingRegisters", m.configuration.AccountingRegisters, ManagerModule, RecordSetModule)
	m.fillObjects("AccumulationRegisters", m.configuration.AccumulationRegisters, ManagerModule, RecordSetModule)
	m.fillObjects("BusinessProcesses", m.configuration.BusinessProcesses, ManagerModule, ObjectModule)
	m.fillObjects("CalculationRegisters", m.configuration.CalculationRegisters, ManagerModule, RecordSetModule)
	m.fillObjects("Catalogs", m.configuration.Catalogs, ManagerModule, ObjectModule)
	m.fillObjects("ChartsOfAccounts", m.configuration.ChartsOfAccounts, ManagerModule, ObjectModule)
	m.fillObjects("ChartsOfCalculationTypes", m.configuration.ChartsOfCalculationTypes, ManagerModule, ObjectModule)
	m.fillObjects("ChartsOfCharacteristicTypes", m.configuration.ChartsOfCharacteristicTypes, ManagerModule, ObjectModule)
	m.fillObjects("CommonCommands", m.configuration.CommonCommands, CommandModule)
	m.fillObjects("CommonForms", m.configuration.CommonForms, FormModule)
	m.fillObjects("CommonModules", m.configuration.CommonModules, SimpleModule)
	m.fillObjects("Constants", m.configuration.Constants, ManagerModule, ValueManagerModule)
	m.fillObjects("DocumentJournals", m.configuration.DocumentJournals, ManagerModule)
	m.fillObjects("Documents", m.configuration.Documents, ManagerModule, ObjectModule)
	m.fillObjects("Enums", m.configuration.Enums, ManagerModule)
	m.fillObjects("ExchangePlans", m.configuration.ExchangePlans, ManagerModule, ObjectModule)
	m.fillObjects("ExchangePlans", m.configuration.FilterCriteria, ManagerModule)
	m.fillObjects("HttpServices", m.configuration.HttpServices, SimpleModule)
	m.fillObjects("InformationRegisters", m.configuration.InformationRegisters, ManagerModule, RecordSetModule)
	m.fillObjects("Reports", m.configuration.Reports, ManagerModule, ObjectModule)
	m.fillObjects("Tasks", m.configuration.Tasks, ManagerModule, ObjectModule)
	m.fillObjects("WebServices", m.configuration.WebServices, SimpleModule)
	m.fillObjects("ExternalDataSources", m.configuration.ExternalDataSources, ManagerModule)

}

func (m *EdtReader) fillObjects(typeName string, objectNames []string, ModuleTypes ...string) {

	for _, objectName := range objectNames {

		nameParts := strings.Split(objectName, ".")
		shortName := nameParts[len(nameParts)-1]
		metaFilename := path.Join(m.srcpath, typeName, shortName, shortName+".mdo")

		edtData := &edtData{}
		edtData.ReadMetaFile(metaFilename)

		for _, ModuleType := range ModuleTypes {
			m.coverData.Data[Module{ModuleUuid: edtData.Uuid, ModuleType: ModuleType}] = path.Join(m.srcpath, typeName, shortName, Filenames[ModuleType])
		}

		for _, form := range edtData.Forms {
			m.coverData.Data[Module{ModuleUuid: form.Uuid, ModuleType: FormModule}] = path.Join(m.srcpath, typeName, shortName, "Forms", form.Name, Filenames[FormModule])
		}

		for _, command := range edtData.Commands {
			m.coverData.Data[Module{ModuleUuid: command.Uuid, ModuleType: CommandModule}] = path.Join(m.srcpath, typeName, shortName, "Commands", command.Name, Filenames[CommandModule])
		}
		for _, recalc := range edtData.Recalculations {
			m.coverData.Data[Module{ModuleUuid: recalc.Uuid, ModuleType: RecordSetModule}] = path.Join(m.srcpath, typeName, shortName, "Recalculations", recalc.Name, Filenames[RecordSetModule])
		}

		if len(edtData.Cubes) > 0 {
			m.fillObjects(path.Join(typeName, shortName, "Cubes"), edtData.Cubes, ManagerModule)
		}
		if len(edtData.Tables) > 0 {
			m.fillObjects(path.Join(typeName, shortName, "Tables"), edtData.Tables, ManagerModule, RecordSetModule, ObjectModule)
		}
	}
}
