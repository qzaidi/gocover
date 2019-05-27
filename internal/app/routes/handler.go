package routes

import (
	"fmt"
	"github.com/qzaidi/gocover/internal/app/store"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Module struct {
	store store.Storage
}

func NewModule(path string) *Module {
	store := store.NewStorage(path)
	return &Module{store: store}
}

func (m *Module) PutCoverageHandler(w http.ResponseWriter, r *http.Request) {
	uniqId := strings.TrimPrefix(r.URL.Path, "/cover/")
	log.Println("inside putCoverage handler", uniqId)

	if r.Method != http.MethodPut {
		http.Error(w, "expected put request", 500)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if len(b) == 0 {
		http.Error(w, "missing request body", 500)
		return
	}

	cov, _ := strconv.Atoi(string(b))

	if cov < 0 || cov > 100 {
		http.Error(w, "Invalid Coverage value", 400)
	}

	err = m.store.PutCoverage(uniqId, cov)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
	w.WriteHeader(200)
}

func (m *Module) GetCoverageHandler(w http.ResponseWriter, r *http.Request) {

	uniqId := strings.TrimPrefix(r.URL.Path, "/img/cover/")
	log.Println("inside coverage handler")

	v := m.store.GetCoverage(uniqId)
	url := fmt.Sprintf("https://img.shields.io/badge/Go Coverage-%d%%25-brightgreen.svg?longCache=true&style=flat", v)
	log.Println("read from db", v, url)

	http.Redirect(w, r, url, 302)
}
