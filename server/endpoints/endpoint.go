package endpoints

import(
  "github.com/emicklei/go-restful"

  "github.com/stensonb/cidrd/server/endpoints/class"
)

func Register(c *restful.Container) {
  class.New().Register(c)
  //netblock.New().Register(c)
}
