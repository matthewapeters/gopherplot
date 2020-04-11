package gopherplot

import (
	"log"
	"net/http"
)

type server struct {
	DataSpace *DataSpace
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
func (s *server) GetStatus(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}

// Serve starts the server
func (d *DataSpace) Serve() {
	s := &server{DataSpace: d}
	http.Handle("/", s)
	http.HandleFunc("/status", s.GetStatus)
	log.Fatal(http.ListenAndServe(":8282", nil))
}
