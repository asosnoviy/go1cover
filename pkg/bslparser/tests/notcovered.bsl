// Известные исключения


Процедура ЦиклБезЕдиногоПрохода()

	Для А = 0 По 1 Цикл
		Прервать;
	КонецЦикла // Будет всегда красная
КонецПроцедуры

Функция УсловияВсеВеткиВозврат(Условие)

	Если Условие Тогда
		Возврат 1;
	Иначе
		Возврат 1;
	КонецЕсли; // Будет всегда красная

КонецФункции

ЦиклБезЕдиногоПрохода();
УсловияВсеВеткиВозврат(Истина);
УсловияВсеВеткиВозврат(Ложь);
