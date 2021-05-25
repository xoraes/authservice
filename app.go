package main

import (
	_ "expvar"
	"flag"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/xoraes/dappauth/routers"
	"log"
	"net/http"
	"time"
)

func main() {

	nossl := flag.Bool("nossl", false, "run server in nossl mode")
	debug := flag.Bool("debug", false, "run server in debug mode")

	flag.Parse()
	logrus.SetLevel(logrus.InfoLevel)
	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	routes := routers.Routes()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	time.Sleep(10 * time.Second)
	if *nossl {
		logrus.Info("starting http server on port 8080")
		logrus.Fatal(http.ListenAndServe(":8080", routes))
	} else {
		logrus.Info("starting https server on port 8081")
		logrus.Fatal(http.ListenAndServeTLS(":8081", "server.crt", "server.key", routes))
	}
}
