package main

import (
	"fmt"
	"github.com/veezex/serv/settings"
	"log"
	"net/http"
)

func main() {
	sett := settings.Parse()

	fs := http.FileServer(http.Dir(sett.Path))
	http.Handle("/", fs)

	err := http.ListenAndServe(fmt.Sprintf(":%d", sett.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
