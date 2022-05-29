package main

import (
	"log"

	"github.com/kangsorang/gotest/pkg/apiserver"
)

func main() {
	config := apiserver.NewConfig()
	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
