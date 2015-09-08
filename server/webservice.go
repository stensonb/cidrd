package server

import (
  "github.com/emicklei/go-restful"

  "github.com/stensonb/cidrd/server/filters"
  "github.com/stensonb/cidrd/server/endpoints"
)

func getContainer() *restful.Container {
  answer := restful.NewContainer()
  answer.Router(restful.CurlyRouter{})
  answer.Filter(filters.NCSACommonLogFormatLogger(logger))

  // register all endpoints
  endpoints.Register(answer)

  return answer
}
