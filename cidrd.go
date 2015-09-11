package main

import (
  "github.com/stensonb/cidrd/config"
  "github.com/stensonb/cidrd/server"
)

func init() {

}

func main() {
  var cfg = config.GetConfig()
  web, err := server.New(cfg)
  if err != nil {
    panic(err)
  }
  web.Start()
  defer web.Stop()
}
