package coverserver

import (
	"io"
	"log"
	"net/http"
)

type CoverageServer struct {
	server *http.ServeMux
	port   string
}

func New(port string, runchan chan bool, stopChan chan bool) *CoverageServer {

	mux := http.NewServeMux()
	mux.HandleFunc("/", root)
	mux.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "start")
		runchan <- true
	})

	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "stop")
		stopChan <- true
	})

	return &CoverageServer{server: mux, port: port}

}

func (s *CoverageServer) Run() {
	log.Println("--server run--")
	http.ListenAndServe(s.port, s.server)

}

func root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello i'am go1cove\n")
}
