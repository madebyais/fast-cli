package main

import "os"

// DEFAULT_FUNCTION_FILE is the default schema that will be written for FAST serverless application
var DEFAULT_FUNCTION_FILE = `// This file contains the basic sample of plugin that will be used as part of FAST serverless application
package main

import (
  "net/http"
  "github.com/labstack/echo"
)

type {plugin_name} struct{}

func (p *{plugin_name}) Call(c echo.Context) error {
  return c.String(http.StatusOK, "Hello world.")
}

var Function {plugin_name}
`

// GOPATH is the default go path
var GOPATH = os.Getenv("GOPATH")
