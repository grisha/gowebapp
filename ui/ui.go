package ui

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/grisha/gowebapp/model"
)

type Config struct {
	Assets http.FileSystem
}

func Start(cfg Config, m *model.Model, listener net.Listener) {

	server := &http.Server{
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16}

	http.Handle("/", indexHandler(m))
	http.Handle("/people", peopleHandler(m))
	http.Handle("/js/", http.FileServer(cfg.Assets))

	go server.Serve(listener)
}

const (
	cdnReact           = "https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"
	cdnReactDom        = "https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"
	cdnBabelStandalone = "https://cdnjs.cloudflare.com/ajax/libs/babel-standalone/6.24.0/babel.min.js"
	cdnAxios           = "https://cdnjs.cloudflare.com/ajax/libs/axios/0.16.1/axios.min.js"
)

const indexHTML = `
<!DOCTYPE HTML>
<html>
  <head>
    <meta charset="utf-8">
    <title>Simple Go Web App</title>
  </head>
  <body>
    <div id='root'></div>
    <script src="` + cdnReact + `"></script>
    <script src="` + cdnReactDom + `"></script>
    <script src="` + cdnBabelStandalone + `"></script>
    <script src="` + cdnAxios + `"></script>
    <script src="/js/app.jsx" type="text/babel"></script>
  </body>
</html>
`

func indexHandler(m *model.Model) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, indexHTML)
	})
}

func peopleHandler(m *model.Model) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		people, err := m.People()
		if err != nil {
			http.Error(w, "This is an error", http.StatusBadRequest)
			return
		}

		js, err := json.Marshal(people)
		if err != nil {
			http.Error(w, "This is an error", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, string(js))
	})
}
