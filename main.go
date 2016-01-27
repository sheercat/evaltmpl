package main

import (
	"flag"
	"fmt"
	"log"

	// _ "github.com/golang/glog"

	"github.com/labstack/echo"
	"github.com/sheercat/evaltmpl/middleware"
	"github.com/sheercat/evaltmpl/renderer"
	"github.com/sheercat/evaltmpl/router"
)

var portNumber = flag.Int("port", 3000, "port number.")

func main() {
	flag.Parse()

	e := echo.New()

	middleware.Set(e)
	router.Set(e)
	renderer.Set(e, router.Routes)

	log.Println(fmt.Sprintf("listen :%d", *portNumber))
	e.Run(fmt.Sprintf(":%d", *portNumber))
}
