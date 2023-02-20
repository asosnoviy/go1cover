package coverserver

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type CoverageClient struct {
	url string
}

func NewClient(port string) *CoverageClient {

	return &CoverageClient{url: "http://localhost:" + port}
}

func (c *CoverageClient) Stop() {
	r, err := http.Get(c.url + "/stop")
	if err != nil {
		log.Println(err)
		return
	}

	b, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	fmt.Println(string(b))

}
