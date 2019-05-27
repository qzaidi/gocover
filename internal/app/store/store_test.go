package store

import (
	"os"
	"testing"
)

func TestPutGet(t *testing.T) {

	keys := []string{"d41d8cd98f00b204e9800998ecf8427e", "b1946ac92492d2347c6235b4d2611184"}

	store := NewStorage("temp.db")
	defer os.Remove("temp.db")

	for idx, key := range keys {
		store.PutCoverage(key, idx)
		if store.GetCoverage(key) != idx {
			t.Error("failed", key, idx)
			break
		}
	}

}

func TestBadPutGet(t *testing.T) {

	store := NewStorage("temp.db")
	defer os.Remove("temp.db")

	if err := store.PutCoverage("", 4); err == nil {
		t.Error("succeeded in storing value for empty string?")
	} else {
		t.Log(err)
	}

	if store.GetCoverage("") != -1 {
		t.Error("failed empty put")
	}

	if store.GetCoverage("helloworld") != -1 {
		t.Error("failed to get default value")
	}
}
