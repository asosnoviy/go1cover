// fake debug server client
package fdbc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type fdbc struct {
	client         *http.Client
	infoBaseAlias  string
	idOfDebuggerUI string
	debuggerURL    string
	Storage        map[ModuleData][]LIneCoverage
}

func New() *fdbc {

	client := &http.Client{}
	fdbg := &fdbc{
		client:         client,
		infoBaseAlias:  "DefAlias",
		idOfDebuggerUI: "1090f54e-4f23-4193-b005-5e59fe488bdf",
		debuggerURL:    "http://127.0.0.1:1760",
		Storage:        make(map[ModuleData][]LIneCoverage),
	}

	return fdbg
}

func (f *fdbc) Init() {

	doRequest(f.client, f.debuggerURL+"/e1crdbg/rdbg?cmd=attachDebugUI", attachDebugUI(f))
	doRequest(f.client, f.debuggerURL+"/e1crdbg/rdbg?cmd=initSettings", initSettings(f))
	doRequest(f.client, f.debuggerURL+"/e1crdbg/rdbg?cmd=setAutoAttachSettings", setAutoAttachSettings(f))
}

func (f *fdbc) Attach() {
	doRequest(f.client, f.debuggerURL+"/e1crdbg/rdbg?cmd=setMeasureMode", setMeasureMode(f))
	for i := 0; i < 2; i++ {
		pingResult := doRequest(f.client, f.debuggerURL+"/e1crdbg/rdbg?cmd=pingDebugUIParams&dbgui="+f.idOfDebuggerUI, "")

		// dumpResponse, _ := httputil.DumpResponse(pingResult, true)
		// fmt.Println(string(dumpResponse))

		pingResultComplete(f, pingResult)
		time.Sleep(10 * time.Second)
	}

}

func (f *fdbc) Deattach() {
	doRequest(f.client, f.debuggerURL+"/e1crdbg/rdbg?cmd=setMeasureMode", setMeasureModeOff(f))
}

func doRequest(client *http.Client, url string, body string) *http.Response {

	req := newRequest(url, body)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	// defer resp.Body.Close()
	return resp
}

func newRequest(url string, body string) *http.Request {

	req, _ := http.NewRequest("POST", url, strings.NewReader(body))
	req.Header.Set("1c-Applicationname-Type", "1C:Enterprise DT")
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("User-Agent", "1CV8")

	return req
}

func pingResultComplete(f *fdbc, pingResp *http.Response) {

	if pingResp.StatusCode != 200 {
		fmt.Println("--Ping--")
		return
	}

	fmt.Println("--HasResult--")

	body, _ := io.ReadAll(pingResp.Body)

	measure := Response{}
	measure.Unmarshal(body)

	for _, r := range measure.Result {
		for _, md := range r.Measure.ModuleData {
			module := ModuleData{ObjectID: md.ModuleID.ObjectID, PropertyID: md.ModuleID.PropertyID}
			for _, li := range md.LineInfo {

				data := f.Storage[module]
				f.Storage[module] = append(data, LIneCoverage{LineNo: li.LineNo, Covered: true})

			}

		}
	}

}
