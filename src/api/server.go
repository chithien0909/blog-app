package api

import (
	"../api/router"
	"../config"
	"../auto"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("\n\tListening [::]:%d", config.PORT)
	listen(config.PORT)
}

func listen(port int)  {

	r := router.NEW()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
