package gopherplot

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"text/template"
	"time"
)

var singleImageTemplage string = `
<!DOCTYPE html><html lang="en"><head></head>
<body onLoad='setTimeout(() => {window.location.reload(true);},1000);'>
<table><tr><td><img src="data:image/png;base64,{{.Image}}"></td></tr></table>{{.TS}}
</body></html> 
`

func writeSingleImagePage(w http.ResponseWriter, image *image.Image) {
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *image); err != nil {
		fmt.Println("writeImageWithTemplate unable to encode image", err)
		log.Fatalln("unable to encode image.")
	}
	originalImage64 := base64.StdEncoding.EncodeToString(buffer.Bytes())

	if tmpl, err := template.New("image").Parse(singleImageTemplage); err != nil {
		log.Println("unable to parse single image template.", err)
	} else {
		data := map[string]interface{}{
			"Image": originalImage64,
			"TS":    time.Now()}
		tmpl.Execute(w, data)
	}
}

type server struct {
	DataSpace *DataSpace
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		writeSingleImagePage(w, s.DataSpace.Render())
	}
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
