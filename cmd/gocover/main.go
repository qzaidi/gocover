package main

import (
	"github.com/qzaidi/gocover/internal/app/routes"
	"log"
	"net/http"
)

func main() {

	m := routes.NewModule("coverage.db")

	http.HandleFunc("/img/cover/", m.GetCoverageHandler)
	http.HandleFunc("/cover/", m.PutCoverageHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
