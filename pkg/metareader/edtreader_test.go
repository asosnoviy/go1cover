package metareader

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var edtreader *EdtReader

func init() {
	path := "../../testresourse/cf/edt/src"
	edtreader = NewEdtReader(path)
	edtreader.Parse()

}

func TestNewEdtReader(t *testing.T) {
	assert.NotEmpty(t, edtreader.configuration.Uuid)
	assert.Len(t, edtreader.configuration.AccountingRegisters, 1)
	assert.Len(t, edtreader.configuration.AccumulationRegisters, 1)
	assert.Len(t, edtreader.configuration.BusinessProcesses, 1)
	assert.Len(t, edtreader.configuration.CalculationRegisters, 1)
	assert.Len(t, edtreader.configuration.Catalogs, 1)
	assert.Len(t, edtreader.configuration.ChartsOfAccounts, 1)
	assert.Len(t, edtreader.configuration.ChartsOfCalculationTypes, 1)
	assert.Len(t, edtreader.configuration.ChartsOfCharacteristicTypes, 1)
	assert.Len(t, edtreader.configuration.CommonCommands, 1)
	assert.Len(t, edtreader.configuration.CommonForms, 2)
	assert.Len(t, edtreader.configuration.CommonModules, 7)
	assert.Len(t, edtreader.configuration.Constants, 2)
	assert.Len(t, edtreader.configuration.DocumentJournals, 1)
	assert.Len(t, edtreader.configuration.Documents, 1)
	assert.Len(t, edtreader.configuration.Enums, 1)
	assert.Len(t, edtreader.configuration.ExchangePlans, 1)
	assert.Len(t, edtreader.configuration.ExternalDataSources, 1)
	assert.Len(t, edtreader.configuration.FilterCriteria, 1)
	assert.Len(t, edtreader.configuration.HttpServices, 1)
	assert.Len(t, edtreader.configuration.InformationRegisters, 2)
	assert.Len(t, edtreader.configuration.Reports, 1)
	assert.Len(t, edtreader.configuration.Tasks, 1)
	assert.Len(t, edtreader.configuration.WebServices, 1)
	assert.Len(t, edtreader.configuration.ExchangePlans, 1)
}
func TestEdtMeta(t *testing.T) {

	CheckMeta(t, "Configuration", "46c7c1d0-b04d-4295-9b04-ae3207c18d29", ManagedApplicationModule, SessionModule, ExternalConnectionModule, OrdinaryApplicationModule)

	CheckMeta(t, "AccountingRegisters/РегистрБухгалтерии1", "e5930f2f-15d9-48a1-ac69-379ad990b02a", ManagerModule, RecordSetModule)
	CheckMeta(t, "AccountingRegisters/РегистрБухгалтерии1/Forms/ФормаСписка", "c5274498-e1c2-4113-a354-a8731848cb58", FormModule)

	CheckMeta(t, "AccumulationRegisters/РегистрНакопления1", "8ea07f36-d671-4649-bc7a-94daa939e77f", ManagerModule, RecordSetModule)
	CheckMeta(t, "AccumulationRegisters/РегистрНакопления1/Forms/ФормаСписка", "6adc9271-744f-4cda-88c9-59156ece5b89", FormModule)

	CheckMeta(t, "BusinessProcesses/БизнесПроцесс1", "560a32ca-028d-4b88-b6f2-6b7212bf31f8", ManagerModule, ObjectModule)
	CheckMeta(t, "BusinessProcesses/БизнесПроцесс1/Forms/ФормаБизнесПроцесса", "8c82c4f6-2c9d-4cd2-bad9-6767a2f1326c", FormModule)

	CheckMeta(t, "CalculationRegisters/РегистрРасчета1", "90587c7d-b950-4476-ac14-426e4a83d9c4", ManagerModule, RecordSetModule)
	CheckMeta(t, "CalculationRegisters/РегистрРасчета1/Forms/ФормаСписка", "ccb50988-1274-49b8-a2df-e0e254189782", FormModule)
	CheckMeta(t, "CalculationRegisters/РегистрРасчета1/Recalculations/Перерасчет1", "b6bf09de-cb9b-4d7f-a755-044d4d832b7e", RecordSetModule)

	CheckMeta(t, "Catalogs/Справочник1", "eeef463d-d5e7-42f2-ae53-10279661f59d", ManagerModule, ObjectModule)
	CheckMeta(t, "Catalogs/Справочник1/Forms/ФормаЭлемента", "175b035e-ee35-4fdf-a8b4-c30ce49dee61", FormModule)
	CheckMeta(t, "Catalogs/Справочник1/Forms/ФормаСписка", "1feb8a5b-989e-440d-afe3-9472183c335c", FormModule)
	CheckMeta(t, "Catalogs/Справочник1/Forms/ФормаВыбора", "3b3346ea-1b32-4a2c-b3a6-1a9651659f0c", FormModule)
	CheckMeta(t, "Catalogs/Справочник1/Commands/Команда1", "342ec3c7-82d4-42bb-a5ff-8a756f110744", CommandModule)
	CheckMeta(t, "Catalogs/Справочник1/Commands/Команда2", "9ff91e05-0f21-407f-a391-32f1a50b8724", CommandModule)

	CheckMeta(t, "ChartsOfAccounts/ПланСчетов1", "2766f353-abd2-4e7f-9a95-53f05c83f5d4", ManagerModule, ObjectModule)
	CheckMeta(t, "ChartsOfAccounts/ПланСчетов1/Forms/ФормаСчета", "b9a2eb3b-ba5a-4968-87db-e5e9d0f7172e", FormModule)

	CheckMeta(t, "ChartsOfCalculationTypes/ПланВидовРасчета1", "1755c534-9ccd-49c4-9f8b-2aa066424aaa", ManagerModule, ObjectModule)

	CheckMeta(t, "ChartsOfCharacteristicTypes/ПланВидовХарактеристик1", "f53a24c3-f1dc-43b7-8dcf-eeb8c0b7f452", ManagerModule, ObjectModule)

	CheckMeta(t, "CommonCommands/ОбщаяКоманда1", "a608f796-f58e-4f8a-b63f-272342b32f35", CommandModule)

	CheckMeta(t, "CommonForms/Константы", "91e1be9a-ecd2-4a39-9213-84592c7d50db", FormModule)
	CheckMeta(t, "CommonForms/Форма", "d13d6a3f-17d9-4ba0-8759-4e5a14a9ea1b", FormModule)

	CheckMeta(t, "CommonModules/ГлобальныйОбщийМодуль", "9e9c021c-bdbd-4804-a53a-9442ba9eb18c", SimpleModule)
	CheckMeta(t, "CommonModules/ОбщийМодульВызовСервера", "96504c9f-4126-439a-9c62-19f43d260c82", SimpleModule)
	CheckMeta(t, "CommonModules/ОбщийМодульКлиентСервер", "b999a18b-2c78-4f41-8b36-a229e148cf05", SimpleModule)

	CheckMeta(t, "Constants/ДатаЗавершенияФоновогоЗадания", "d42ac573-3c10-4f7e-8b10-dadbc45c196c", ManagerModule, ValueManagerModule)
	CheckMeta(t, "Constants/Константа1", "61e6a6f2-7057-4e93-96c3-7bd2559217f4", ManagerModule, ValueManagerModule)

	CheckMeta(t, "ExternalDataSources/ВнешнийИсточникДанных1/Cubes/Куб1", "3a2f77a9-afc0-4f86-bc78-d0aa4dd2b7d3", ManagerModule)

	CheckMeta(t, "ExternalDataSources/ВнешнийИсточникДанных1/Tables/Таблица1", "a2daba70-c996-4cd4-9478-00f901ac7ee6", ManagerModule, ObjectModule, RecordSetModule)
	CheckMeta(t, "ExternalDataSources/ВнешнийИсточникДанных1/Tables/Таблица1/Forms/ФормаОбъекта", "b61902e8-e31a-4552-baea-206b787001dc", FormModule)

	CheckMeta(t, "ExchangePlans/ПланОбмена1", "242cb07d-3d2b-4689-b590-d3ed23ac9d10", ManagerModule, ObjectModule)
	CheckMeta(t, "FilterCriteria/КритерийОтбора1", "6e9d3381-0607-43df-866d-14ee5d65a294", ManagerModule)

}

func CheckMeta(t *testing.T, foldername string, Uuid string, ModuleTypes ...string) {

	for _, ModuleType := range ModuleTypes {

		module := Module{ModuleUuid: Uuid, ModuleType: ModuleType}
		assert.Contains(t, edtreader.coverData.Data, module)

		localfilepath := path.Join(edtreader.srcpath, foldername, Filenames[ModuleType])
		assert.Equal(t, localfilepath, edtreader.coverData.Data[module])
		assert.Contains(t, localfilepath, ".bsl")
	}

}
