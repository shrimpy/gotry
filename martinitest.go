package main

import (
    "os" 
    "github.com/go-martini/martini"
)

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world from martini!"
  })
  m.RunOnAddr(":" + os.Getenv("HTTP_PLATFORM_PORT"))
}