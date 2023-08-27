package main

import (
	"flag"
	"github.com/koding/websocketproxy"
	"log"
	"net/http"
	"net/url"
)

var (
	flagBackend = flag.String("backend", "", "Backend URL for proxying")
)

func main() {
	flag.Parse()

	u, err := url.Parse(*flagBackend)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("backend:%+v", u)

	err = http.ListenAndServe(":8084", websocketproxy.NewProxy(u))
	if err != nil {
		log.Fatalln(err)
	}
}
