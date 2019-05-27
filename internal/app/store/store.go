package store

import (
	"encoding/binary"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

type Storage interface {
	PutCoverage(key string, val int) error
	GetCoverage(key string) int
}

type Module struct {
	path string
	db   *bolt.DB
}

func NewStorage(path string) *Module {
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		log.Fatalln("failed to open db at", path)
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("coverage"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return &Module{path: path, db: db}
}

func (m *Module) PutCoverage(key string, val int) error {
	if key == "" || val < 0 || val > 100 {
		return fmt.Errorf("invalid key or value")
	}

	err := m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("coverage"))
		return b.Put([]byte(key), itob(val))
	})
	return err
}

func (m *Module) GetCoverage(key string) int {
	v := -1

	if key == "" {
		return v
	}

	err := m.db.View(func(tx *bolt.Tx) error {
		buf := tx.Bucket([]byte("coverage")).Get([]byte(key))
		if buf != nil {
			v = btoi(buf)
		}
		return nil
	})

	if err != nil {
		log.Println("GetCoverage:Error:", err, key)
	}

	return v
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(v []byte) int {
	return int(binary.BigEndian.Uint64(v))
}
