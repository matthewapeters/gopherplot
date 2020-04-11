package gopherplot

import (
	"net/http" 
	"log"
)

type server struct {
	DataSpace 
}

func (s &server) ServeHTTP(w http.ResponseWriter, r http.Request){

}

func Serve(){
	s:= &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8282", nil)))
}
