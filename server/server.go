package server

import(
  "os"
  "log"
  "fmt"
  "net/http"
  "github.com/emicklei/go-restful"

  "github.com/stensonb/cidrd/config"
)

var logger *log.Logger = log.New(os.Stdout, "", 0)

type WebServer struct {
  Port string
  Container *restful.Container
}

func New(cfg config.Config) WebServer {
  return WebServer{Port: fmt.Sprint(cfg.Port),
    Container: getContainer(),
  }
}

func (w *WebServer) Start() {
  server := &http.Server{Addr: ":" + w.Port, Handler: w.Container}
  logger.Printf("Listening on localhost:%s", w.Port)
 	logger.Fatal(server.ListenAndServe())
}
