package gopherplot_test

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	gp "github.com/matthewapeters/gopherplot"
)

func TestServer(t *testing.T) {
	d := gp.DataSpace{}

	doTest := make(chan bool)
	go func() {
		go d.Serve()
		<-time.After(3 * time.Second)
		close(doTest)
		<-time.After(15 * time.Second)
	}()
	<-doTest

	resp, err := http.Get("http://localhost:8282/status")
	if err != nil {
		t.Error(err)
	}
	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("reading response %s", err)
		resp.Body.Close()
		return
	}
	if string(output) != "OK" {
		t.Errorf("%s != OK", string(output))
		resp.Body.Close()
		return
	}
	resp.Body.Close()
}
