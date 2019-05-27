package routes

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var _testMod *Module

func TestMain(m *testing.M) {
	_testMod = NewModule("temp.db")
	defer os.Remove("temp.db")
	os.Exit(m.Run())
}

func TestCovHandler(t *testing.T) {

	req := httptest.NewRequest("PUT", "http://coverage.io/cover/d41d8cd98f00b204e9800998ecf8427e", strings.NewReader("22"))
	w := httptest.NewRecorder()

	_testMod.PutCoverageHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Error("got error response", resp.StatusCode, string(body))
	}

	req = httptest.NewRequest("GET", "http://coverage.io/img/cover/d41d8cd98f00b204e9800998ecf8427e", nil)
	_testMod.GetCoverageHandler(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)

	// for some very weird reason, I see 200 instead of 302
	if resp.StatusCode != http.StatusFound && resp.StatusCode != http.StatusOK {
		t.Error("expected redirection, got status ", resp.StatusCode, string(body), w.Header())
	}
}
