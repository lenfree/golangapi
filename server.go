package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Get("/v1/api/status", status)
  m.Get("/v1/api/version", version)
  m.Run()
}
