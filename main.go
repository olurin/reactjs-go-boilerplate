package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"./modules"
	"gopkg.in/macaron.v1"
)

type Configure struct {
	port    int
	root    string
	private bool
}

var gcfg = Configure{}

var m *macaron.Macaron

func init() {
	m = macaron.Classic()
//	m.Use(modules.Public)
	m.Use(modules.Renderer)

	flag.IntVar(&gcfg.port, "port", 8000, "Which port to listen")
	flag.BoolVar(&gcfg.private, "private", false, "Only listen on lookback interface, otherwise listen on all interface")
}

func initRouters() {
	m.Get("/", func(ctx *macaron.Context) {
		ctx.HTML(200, "homepage", nil)
	})
	m.Use(macaron.Static("dist"))
}

func main() {
	flag.Parse()
	initRouters()

	http.Handle("/", m)

	int := ":" + strconv.Itoa(gcfg.port)
	p := strconv.Itoa(gcfg.port)
	mesg := "; please visit http://127.0.0.1:" + p
	if gcfg.private {
		int = "localhost" + int
		log.Printf("listens on 127.0.0.1@" + p + mesg)
	} else {
		log.Printf("listens on 0.0.0.0@" + p + mesg)
	}
	if err := http.ListenAndServe(int, nil); err != nil {
		log.Fatal(err)
	}
}
