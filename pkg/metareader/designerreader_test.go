package metareader

import (
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type metaobject struct {
	name   string
	id     string
	folder string
}

var src = "../../testresourse/cf/designer/src/"
var reader *DesignerReader

func init() {
	reader = NewDesignerReader(src)
	reader.Parse()
	src, _ = filepath.Abs(src)
}

func TestReader(t *testing.T) {

	const (
		РегистрБухгалтерии1_ФормаСписка    = "c5274498-e1c2-4113-a354-a8731848cb58"
		БизнесПроцесс1                     = "560a32ca-028d-4b88-b6f2-6b7212bf31f8"
		БизнесПроцесс1_ФормаБизнесПроцесса = "8c82c4f6-2c9d-4cd2-bad9-6767a2f1326c"
	)

	meta := reader.coverData.Data

	module := Module{ModuleUuid: РегистрБухгалтерии1_ФормаСписка, ModuleType: FormModule}
	assert.Contains(t, meta, module)
	assert.Equal(t, path.Join(src, "AccountingRegisters", "РегистрБухгалтерии1", "Forms", "ФормаСписка", "Ext", "Form", "Module.bsl"), meta[module])

	module = Module{ModuleUuid: БизнесПроцесс1_ФормаБизнесПроцесса, ModuleType: FormModule}
	assert.Contains(t, meta, module)
	assert.Equal(t, path.Join(src, "BusinessProcesses", "БизнесПроцесс1", "Forms", "ФормаБизнесПроцесса", "Ext", "Form", "Module.bsl"), meta[module])
}

func TestRegisters(t *testing.T) {

	meta := reader.coverData.Data

	var registers []metaobject
	registers = append(registers, metaobject{"РегистрБухгалтерии1", "e5930f2f-15d9-48a1-ac69-379ad990b02a", "AccountingRegisters"})
	registers = append(registers, metaobject{"РегистрРасчета1", "90587c7d-b950-4476-ac14-426e4a83d9c4", "CalculationRegisters"})
	registers = append(registers, metaobject{"РегистрСведений1", "184d9d78-9523-4cfa-9542-a7ba72efe4dd", "InformationRegisters"})
	registers = append(registers, metaobject{"РегистрНакопления1", "8ea07f36-d671-4649-bc7a-94daa939e77f", "AccumulationRegisters"})

	for _, v := range registers {
		checkRegister(t, meta, &v)
	}
}

func TestInstans(t *testing.T) {

	meta := reader.coverData.Data

	var instants []metaobject
	instants = append(instants, metaobject{"Справочник1", "eeef463d-d5e7-42f2-ae53-10279661f59d", "Catalogs"})
	instants = append(instants, metaobject{"БизнесПроцесс1", "560a32ca-028d-4b88-b6f2-6b7212bf31f8", "BusinessProcesses"})
	instants = append(instants, metaobject{"Документ1", "ce4fb46b-4af7-493e-9fcb-76ad8c4f8acd", "Documents"})
	instants = append(instants, metaobject{"ПланОбмена1", "242cb07d-3d2b-4689-b590-d3ed23ac9d10", "ExchangePlans"})
	instants = append(instants, metaobject{"Отчет1", "34d3754d-298c-4786-92f6-a487db249fc7", "Reports"})
	instants = append(instants, metaobject{"Обработка1", "a7c57ba0-75d8-487d-b8ea-ae5083d8a503", "DataProcessors"})
	instants = append(instants, metaobject{"Задача1", "c251fcec-ec02-4ef4-8f70-4d70db6631ea", "Tasks"})
	instants = append(instants, metaobject{"ПланСчетов1", "2766f353-abd2-4e7f-9a95-53f05c83f5d4", "ChartsOfAccounts"})
	instants = append(instants, metaobject{"ПланВидовРасчета1", "1755c534-9ccd-49c4-9f8b-2aa066424aaa", "ChartsOfCalculationTypes"})
	instants = append(instants, metaobject{"ПланВидовХарактеристик1", "f53a24c3-f1dc-43b7-8dcf-eeb8c0b7f452", "ChartsOfCharacteristicTypes"})

	for _, v := range instants {
		checkInstant(t, meta, &v)
	}
}

func TestSimple(t *testing.T) {

	meta := reader.coverData.Data

	var simple []metaobject
	simple = append(simple, metaobject{"ПростойОбщийМодуль", "1be4af7e-334e-49fa-a9f9-d80c737ff954", "CommonModules"})
	simple = append(simple, metaobject{"WebСервис1", "d7f9b06b-0799-486e-adff-c45a2d5b8101", "WebServices"})
	simple = append(simple, metaobject{"HTTPСервис1", "3f029e1e-5a9e-4446-b74f-cbcb79b1e2fe", "HTTPServices"})

	for _, v := range simple {
		checkSimple(t, meta, &v)
	}
}

func TestEnum(t *testing.T) {

	meta := reader.coverData.Data

	enum := metaobject{id: "f11f3441-4b64-4344-b1a0-0e4b3e466e03", name: "Перечисление1", folder: "Enums"}

	module := Module{ModuleUuid: enum.id, ModuleType: ManagerModule}
	assert.Contains(t, meta, module)
	assert.Equal(t, path.Join(src, enum.folder, enum.name, "Ext", "ManagerModule.bsl"), meta[module])
}

func TestFilterCriteria(t *testing.T) {

	meta := reader.coverData.Data

	FilterCriteria := metaobject{id: "6e9d3381-0607-43df-866d-14ee5d65a294", name: "КритерийОтбора1", folder: "FilterCriteria"}

	module := Module{ModuleUuid: FilterCriteria.id, ModuleType: ManagerModule}
	assert.Contains(t, meta, module)
	assert.Equal(t, path.Join(src, FilterCriteria.folder, FilterCriteria.name, "Ext", "ManagerModule.bsl"), meta[module])
}

func TestConstant(t *testing.T) {

	meta := reader.coverData.Data

	enum := metaobject{id: "61e6a6f2-7057-4e93-96c3-7bd2559217f4", name: "Константа1", folder: "Constants"}

	module := Module{ModuleUuid: enum.id, ModuleType: ManagerModule}
	assert.Contains(t, meta, module)
	assert.Equal(t, path.Join(src, enum.folder, enum.name, "Ext", "ManagerModule.bsl"), meta[module])

	module = Module{ModuleUuid: enum.id, ModuleType: ValueManagerModule}
	assert.Contains(t, meta, module)
	assert.Equal(t, path.Join(src, enum.folder, enum.name, "Ext", "ValueManagerModule.bsl"), meta[module])
}

func TestCommonForms(t *testing.T) {

	meta := reader.coverData.Data

	form := metaobject{id: "91e1be9a-ecd2-4a39-9213-84592c7d50db", name: "Константы", folder: "CommonForms"}

	module := Module{ModuleUuid: form.id, ModuleType: FormModule}
	assert.Contains(t, meta, module)
	assert.Equal(t, path.Join(src, form.folder, form.name, "Ext", "Form", "Module.bsl"), meta[module])
}

func TestCommonCommands(t *testing.T) {

	meta := reader.coverData.Data

	command := metaobject{id: "a608f796-f58e-4f8a-b63f-272342b32f35", name: "ОбщаяКоманда1", folder: "CommonCommands"}

	module := Module{ModuleUuid: command.id, ModuleType: CommandModule}
	assert.Contains(t, meta, module)
	assert.Equal(t, path.Join(src, command.folder, command.name, "Ext", "CommandModule.bsl"), meta[module])
}

func TestCommands(t *testing.T) {

	meta := reader.coverData.Data

	var commands []metaobject

	commands = append(commands, metaobject{id: "342ec3c7-82d4-42bb-a5ff-8a756f110744", name: "Команда1", folder: "Catalogs/Справочник1"})

	for _, command := range commands {
		module := Module{ModuleUuid: command.id, ModuleType: CommandModule}
		assert.Contains(t, meta, module)
		assert.Equal(t, path.Join(src, command.folder, "Commands", command.name, "Ext", "CommandModule.bsl"), meta[module])
	}
}

func TestSettingsStorage(t *testing.T) {

	meta := reader.coverData.Data

	enum := metaobject{id: "e7a9947d-7565-4681-b75c-c9a229b45042", name: "ХранилищеНастроек1", folder: "SettingsStorages"}

	module := Module{ModuleUuid: enum.id, ModuleType: ManagerModuleStorage}
	assert.Contains(t, meta, module)
	assert.Equal(t, path.Join(src, enum.folder, enum.name, "Ext", "ManagerModule.bsl"), meta[module])
}

func TesDocumentJournals(t *testing.T) {

	meta := reader.coverData.Data

	enum := metaobject{id: "c6743657-4787-40de-9a45-2493c630f626", name: "ЖурналДокументов1", folder: "DocumentJournals"}

	module := Module{ModuleUuid: enum.id, ModuleType: ManagerModule}
	assert.Contains(t, meta, module)
	assert.Equal(t, path.Join(src, enum.folder, enum.name, "Ext", "ManagerModule.bsl"), meta[module])
}

func checkRegister(t *testing.T, meta map[Module]string, m *metaobject) {
	module := Module{ModuleUuid: m.id, ModuleType: RecordSetModule}
	assert.Contains(t, meta, module)
	assert.Equal(t, path.Join(src, m.folder, m.name, "Ext", "RecordSetModule.bsl"), meta[module])

	module = Module{ModuleUuid: m.id, ModuleType: ManagerModule}
	assert.Contains(t, meta, module)
	assert.Equal(t, path.Join(src, m.folder, m.name, "Ext", "ManagerModule.bsl"), meta[module])
}

func checkInstant(t *testing.T, meta map[Module]string, m *metaobject) {
	module := Module{ModuleUuid: m.id, ModuleType: ObjectModule}
	assert.Contains(t, meta, module)
	assert.Equal(t, path.Join(src, m.folder, m.name, "Ext", "ObjectModule.bsl"), meta[module])

	module = Module{ModuleUuid: m.id, ModuleType: ManagerModule}
	assert.Contains(t, meta, module)
	assert.Equal(t, path.Join(src, m.folder, m.name, "Ext", "ManagerModule.bsl"), meta[module])
}

func checkSimple(t *testing.T, meta map[Module]string, m *metaobject) {
	module := Module{ModuleUuid: m.id, ModuleType: SimpleModule}
	assert.Contains(t, meta, module)
	assert.Equal(t, path.Join(src, m.folder, m.name, "Ext", "Module.bsl"), meta[module])
}
