// fake debug server client
package fdbc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Fdbc struct {
	client         *http.Client
	infoBaseAlias  string
	idOfDebuggerUI string
	debuggerURL    string
	timeChan       *time.Ticker
	stopChan       chan bool
	sync.Mutex
	Storage map[ModuleData][]LIneCoverage
}

func New(debuggerURL string) *Fdbc {

	client := &http.Client{}
	fdbg := &Fdbc{
		client:         client,
		infoBaseAlias:  "DefAlias",
		idOfDebuggerUI: "1090f54e-4f23-4193-b005-5e59fe488bdf",
		debuggerURL:    debuggerURL,
		Storage:        make(map[ModuleData][]LIneCoverage),
		timeChan:       time.NewTicker(500 * time.Millisecond),
		stopChan:       make(chan bool, 2),
	}

	return fdbg
}

func (f *Fdbc) Init() {

	if len(f.debuggerURL) == 0 {

		log.Println("Empty debugger URL debug connect skipped")
		return

	}

	doRequest(f.client, f.debuggerURL+"/e1crdbg/rdbg?cmd=attachDebugUI", attachDebugUI(f))
	doRequest(f.client, f.debuggerURL+"/e1crdbg/rdbg?cmd=initSettings", initSettings(f))
	doRequest(f.client, f.debuggerURL+"/e1crdbg/rdbg?cmd=setAutoAttachSettings", setAutoAttachSettings(f))
}

func (f *Fdbc) Attach() {

	if len(f.debuggerURL) == 0 {

		log.Println("Empty debugger URL debug connect skipped")
		return

	}

	doRequest(f.client, f.debuggerURL+"/e1crdbg/rdbg?cmd=setMeasureMode", setMeasureMode(f))

	go func() {

		for {
			select {
			case <-f.stopChan:
				{
					return
				}
			case <-f.timeChan.C:
				{
					pingResult := doRequest(f.client, f.debuggerURL+"/e1crdbg/rdbg?cmd=pingDebugUIParams&dbgui="+f.idOfDebuggerUI, "")
					pingResultComplete(f, pingResult)
				}

			}
		}
	}()

}

func (f *Fdbc) Deattach() {

	if len(f.debuggerURL) == 0 {

		log.Println("Empty debugger URL debug connect skipped")
		return

	}
	doRequest(f.client, f.debuggerURL+"/e1crdbg/rdbg?cmd=setMeasureMode", setMeasureModeOff(f))
}

func (f *Fdbc) Stop() {

	f.timeChan.Stop()
	f.stopChan <- true
}

func (f *Fdbc) LinesCount() int {
	linescount := 0
	for _, v := range f.Storage {

		linescount += len(v)
	}
	return linescount
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

func pingResultComplete(f *Fdbc, pingResp *http.Response) {

	if pingResp.StatusCode != 200 {
		fmt.Println("--Ping--")
		return
	}

	fmt.Println("--HasResult--")

	body, _ := io.ReadAll(pingResp.Body)

	measure := Response{}
	measure.Unmarshal(body)

	// fmt.Println(measure)

	for _, r := range measure.Result {
		for _, md := range r.Measure.ModuleData {
			module := ModuleData{ObjectID: md.ModuleID.ObjectID, PropertyID: md.ModuleID.PropertyID}
			for _, li := range md.LineInfo {

				data := f.Storage[module]
				f.Mutex.Lock()
				f.Storage[module] = append(data, LIneCoverage{LineNo: li.LineNo, Covered: true})
				f.Mutex.Unlock()
			}

		}
	}

}
