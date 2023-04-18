// fake debug server client

package fdbc

import (
	"testing"
	"time"

	"github.com/h2non/gock"
)

func TestFdbc_Init(t *testing.T) {

	defer gock.Off()

	debugClient := New("http://example.com")

	gock.New(debugClient.debuggerURL).
		MatchParam("cmd", "attachDebugUI").
		Post("e1crdbg/rdbg").
		Reply(200)

	gock.New(debugClient.debuggerURL).
		MatchParam("cmd", "initSettings").
		Post("e1crdbg/rdbg").
		Reply(200)

	gock.New(debugClient.debuggerURL).
		MatchParam("cmd", "setAutoAttachSettings").
		Post("e1crdbg/rdbg").
		Reply(200)

	debugClient.Init()

	gock.New(debugClient.debuggerURL).
		MatchParam("cmd", "setMeasureMode").
		Post("e1crdbg/rdbg").
		Reply(200)

	debugClient.timeChan = time.NewTimer(10 * time.Minute).C
	debugClient.Attach()

	gock.New(debugClient.debuggerURL).
		MatchParams(map[string]string{"cmd": "pingDebugUIParams", "dbgui": debugClient.idOfDebuggerUI}).
		Post("e1crdbg/rdbg").
		Reply(200)

	ping(debugClient)

	gock.New(debugClient.debuggerURL).
		MatchParams(map[string]string{"cmd": "pingDebugUIParams", "dbgui": debugClient.idOfDebuggerUI}).
		Post("e1crdbg/rdbg").
		Reply(204)

	ping(debugClient)

}
