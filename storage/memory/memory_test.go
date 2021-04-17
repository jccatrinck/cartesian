package memory

import (
	"log"
	"math/rand"
	"os"
	"testing"
)

var m = New()

func TestMain(m *testing.M) {
	err := setup()

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func setup() (err error) {
	err = setupPoints()

	if err != nil {
		return
	}

	return
}

func randomRange(min, max int) int {
	return rand.Intn(max-min) + min
}
