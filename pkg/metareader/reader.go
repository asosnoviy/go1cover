package metareader

import (
	"encoding/xml"
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

	m.coverData.Data[Module{ModuleTUuid: m.configuration.Configuration.Uuid, ModuleType: "9b7bbbae-9771-46f2-9e4d-2489e0ffc702"}] = path.Join(m.srcpath, "Ext", "SessionModule.bsl")
	m.coverData.Data[Module{ModuleTUuid: m.configuration.Configuration.Uuid, ModuleType: "a4a9c1e2-1e54-4c7f-af06-4ca341198fac"}] = path.Join(m.srcpath, "Ext", "ExternalConnectionModule.bsl")
	m.coverData.Data[Module{ModuleTUuid: m.configuration.Configuration.Uuid, ModuleType: "a78d9ce3-4e0c-48d5-9863-ae7342eedf94"}] = path.Join(m.srcpath, "Ext", "OrdinaryApplicationModule.bsl")
	m.coverData.Data[Module{ModuleTUuid: m.configuration.Configuration.Uuid, ModuleType: "d22e852a-cf8a-4f77-8ccb-3548e7792bea"}] = path.Join(m.srcpath, "Ext", "ManagedApplicationModule.bsl")

	m.fillID("Documents", m.configuration.Configuration.ChildObjects.Document)
	m.fillID("CommonModules", m.configuration.Configuration.ChildObjects.CommonModule)

}

func (m *Metareader) CoverData() *coverData {

	return &m.coverData
}

func (m *Metareader) fillID(typeName string, objectNames []string) {

	for _, v := range objectNames {
		filename := path.Join(m.srcpath, typeName, v+".xml")
		file, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		switch typeName {
		case "Documents":
			metaData := metaDataObjectDocument{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleTUuid: metaData.Data.Uuid, ModuleType: "d1b64a2c-8078-4982-8190-8f81aefda192"}] = path.Join(m.srcpath, typeName, v, "Ext", "ModuleManager.bsl")
			m.coverData.Data[Module{ModuleTUuid: metaData.Data.Uuid, ModuleType: "a637f77f-3840-441d-a1c3-699c8c5cb7e0"}] = path.Join(m.srcpath, typeName, v, "Ext", "ObjectModule.bsl")
		case "CommonModules":
			metaData := metaDataObjectCommonModule{}
			xml.Unmarshal(file, &metaData)
			m.coverData.Data[Module{ModuleTUuid: metaData.Data.Uuid, ModuleType: "d5963243-262e-4398-b4d7-fb16d06484f6"}] = path.Join(m.srcpath, typeName, v, "Module.bsl")
		default:
			panic("uncnown metadatatype")
		}

	}

}
