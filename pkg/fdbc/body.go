package fdbc

func attachDebugUI(f *Fdbc) string {
	return `<?xml version="1.0" encoding="UTF-8"?>
	<request xmlns:response="http://v8.1c.ru/8.3/debugger/debugRDBGRequestResponse">
	  <response:infoBaseAlias>` + f.infoBaseAlias + `</response:infoBaseAlias>
	  <response:idOfDebuggerUI>` + f.idOfDebuggerUI + `</response:idOfDebuggerUI>
	</request>`
}

func initSettings(f *Fdbc) string {
	return `<?xml version="1.0" encoding="UTF-8"?>
	<request xmlns:response="http://v8.1c.ru/8.3/debugger/debugRDBGRequestResponse">
	  <response:infoBaseAlias>` + f.infoBaseAlias + `</response:infoBaseAlias>
	  <response:idOfDebuggerUI>` + f.idOfDebuggerUI + `</response:idOfDebuggerUI>
	</request>`
}

func setMeasureMode(f *Fdbc) string {
	return `<?xml version="1.0" encoding="UTF-8"?>
	<request xmlns:response="http://v8.1c.ru/8.3/debugger/debugRDBGRequestResponse">
	  <response:infoBaseAlias>` + f.infoBaseAlias + `</response:infoBaseAlias>
	  <response:idOfDebuggerUI>` + f.idOfDebuggerUI + `</response:idOfDebuggerUI>
	  <response:measureModeSeanceID>4d049139-8eca-4502-90e4-5c1cbb7cac9c</response:measureModeSeanceID>
	</request>`
}

func setMeasureModeOff(f *Fdbc) string {
	return `<?xml version="1.0" encoding="UTF-8"?>
	<request xmlns:response="http://v8.1c.ru/8.3/debugger/debugRDBGRequestResponse">
	  <response:infoBaseAlias>` + f.infoBaseAlias + `</response:infoBaseAlias>
	  <response:idOfDebuggerUI>` + f.idOfDebuggerUI + `</response:idOfDebuggerUI>
	</request>`
}

func setAutoAttachSettings(f *Fdbc) string {
	return `<?xml version="1.0" encoding="UTF-8"?>
	<request xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:attach="http://v8.1c.ru/8.3/debugger/debugAutoAttach" xmlns:response="http://v8.1c.ru/8.3/debugger/debugRDBGRequestResponse">
	  <response:infoBaseAlias>` + f.infoBaseAlias + `</response:infoBaseAlias>
	  <response:idOfDebuggerUI>` + f.idOfDebuggerUI + `</response:idOfDebuggerUI>
	  <response:autoAttachSettings>
		<attach:targetType>Client</attach:targetType>
		<attach:targetType>Server</attach:targetType>
		<attach:targetType>ServerEmulation</attach:targetType>
		<attach:targetType>WEBClient</attach:targetType>
		<attach:targetType>JOB</attach:targetType>
		<attach:targetType>JobFileMode</attach:targetType>
		<attach:areaName xsi:nil="true"/>
	  </response:autoAttachSettings>
	</request>`
}
