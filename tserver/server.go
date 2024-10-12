package tserver

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	config
}

func NewServer(c config) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Port), nil))
}
