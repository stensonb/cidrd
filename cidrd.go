package main

import (
  "github.com/stensonb/cidrd/config"
  "github.com/stensonb/cidrd/server"
)

func init() {

}

func main() {
  var cfg = config.GetConfig()
  var web = server.New(cfg)
  web.Start()
}
