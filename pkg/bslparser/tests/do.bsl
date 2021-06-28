Перем А;

Процедура Инициализация()
	А = Новый Массив;
	А.Добавить(0);
	А.Добавить(1);
	А.Добавить(2);
КонецПроцедуры

Процедура ЦиклДляКаждого()

	Для
	Каждого
	Б
	Из
	ПолучитьА
	(
	А
	)
	.
	А
	Цикл
	КонецЦикла;

	Для
	Каждого
	Б
	Из
	А
	Цикл
		Если Б = "1" Тогда
			Продолжить;
		КонецЕсли;
		Если Б = "2" Тогда
			Возврат;
		КонецЕсли;
	КонецЦикла;

КонецПроцедуры

Процедура ЦиклДля()

	В = 1;
	Для
	Б
	=
	-
	В
	По
	В
	+
	1
	Цикл
	КонецЦикла;

	Для
	Б
	=
	0
	По
	А
	.
	Количество()
	Цикл
		Если Б = 0 Тогда
			Продолжить;
		КонецЕсли;
	КонецЦикла;

	Для
	Б
	=
	0
	По
	КоличествоА()
	Цикл
		Если Б > 0 Тогда
			Прервать;
		КонецЕсли;
	КонецЦикла;

КонецПроцедуры

Процедура ЦиклПока()
	Б = 1;
	Пока
	Б
	=
	А
	.
	Получить
	(
	1
	)
	Цикл
	Б = 0;
	КонецЦикла;

КонецПроцедуры

Функция КоличествоА()
	Возврат А.Количество()
КонецФункции

Функция ПолучитьА(А = 0)
	Возврат Новый Структура("А", А);
КонецФункции

Инициализация();
ЦиклДляКаждого();
ЦиклДля();
ЦиклПока();