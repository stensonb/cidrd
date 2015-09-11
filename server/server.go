package server

import(
  "os"
  "log"
  "fmt"
  "errors"
  "net/http"
  "github.com/emicklei/go-restful"

  "github.com/stensonb/cidrd/config"
  "github.com/stensonb/cidrd/db"
  "github.com/stensonb/cidrd/model"
  "github.com/stensonb/cidrd/controller"
  "github.com/stensonb/cidrd/server/filters"
)

var logger *log.Logger = log.New(os.Stdout, "", 0)

type WebServer struct {
  Port string
  Container *restful.Container
  Model *model.Model
  Controller *controller.Controller
}

func New(cfg *config.Config) (*WebServer, error) {
  // create the DB connection
  handle, err := db.New(cfg)
  if err != nil {
    return &WebServer{}, errors.New("Failed to create DB connection.")
  }

  // initialize the data model
  m := model.New(handle)

  // start with a standard web Container
  container := newContainer()

  // initialize the controllers
  c := controller.New(cfg, m, container)

  return &WebServer{Port: fmt.Sprint(cfg.Port),
    Container: container,
    Model: m,
    Controller: c,
  }, nil
}

func newContainer() *restful.Container {
  answer := restful.NewContainer()
  answer.EnableContentEncoding(true)  // enable GZIP encoding
  answer.Router(restful.CurlyRouter{})
  answer.Filter(filters.NCSACommonLogFormatLogger(logger))

  return answer
}

func (w *WebServer) Start() {
  server := &http.Server{Addr: ":" + w.Port, Handler: w.Container}
  logger.Printf("Listening on localhost:%s", w.Port)
 	logger.Fatal(server.ListenAndServe())
}

func (w *WebServer) Stop() {
  // TODO: db cleanup, closing
  w.Model.Close()
}
