package plugin

import (
  "github.com/emicklei/go-restful"

	"github.com/stensonb/cidrd/config"
	"github.com/stensonb/cidrd/model"
)

// Endpoint defines the common interface for the different controller endpoints
type Endpoint interface {
  Register(*restful.Container)
  Configure(*model.Model, *config.Config)
  Name() string
}
