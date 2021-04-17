package mysql

import (
	"log"
	"math/rand"
	"net"
	"os"
	"testing"
	"time"
)

var m *MySQL

func TestMain(m *testing.M) {
	err := setup()

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func setup() (err error) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(i) * time.Second)

		m, err = New()

		if err == nil {
			break
		}

		if _, ok := err.(*net.OpError); !ok {
			break
		}

		log.Println("Waiting for MySQL server startup.")
	}

	if err != nil {
		log.Printf("%T: %+v\n", err, err)
		return
	}

	err = setupPoints()

	if err != nil {
		return
	}

	return
}

func randomRange(min, max int) int {
	return rand.Intn(max-min) + min
}
