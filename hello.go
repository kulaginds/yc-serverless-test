package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var helloMessage = []byte(`hello world`)

func main() {
	var err error

	httpPort := 8080
	httpPortEnv := os.Getenv("HTTP_PORT")
	if httpPortEnv != "" {
		var newHttpPort int

		newHttpPort, err = strconv.Atoi(httpPortEnv)
		if err != nil {
			log.Printf("can't parse HTTP_PORT: %v\n", err)
		} else {
			httpPort = newHttpPort
		}
	}

	http.HandleFunc("/", helloHandler)

	addr := fmt.Sprintf(":%d", httpPort)

	log.Printf("starting listening requests on %s\n", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("start webserver error: %v\n", err)
	}
}

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write(helloMessage)
}
