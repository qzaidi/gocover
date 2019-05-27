package main

import (
	"log"
	"net/http"
  "github.com/qzaidi/gocover/internal/app/routes"
)

func main() {

	m := routes.NewModule()

	http.HandleFunc("/img/cover/", m.GetCoverageHandler)
	http.HandleFunc("/cover/", m.PutCoverageHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
