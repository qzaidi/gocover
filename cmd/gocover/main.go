package main

import (
	"github.com/CrowdSurge/banner"
	"github.com/qzaidi/gocover/internal/app/routes"
	"log"
	"net/http"
)

func main() {

	banner.Print("gocover")
	m := routes.NewModule("coverage.db")

	http.HandleFunc("/img/cover/", m.GetCoverageHandler)
	http.HandleFunc("/cover/", m.PutCoverageHandler)

	port := ":9000"
	log.Println("listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
