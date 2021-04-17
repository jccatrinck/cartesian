package points

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/jccatrinck/cartesian/services/points/model"

	"github.com/stretchr/testify/assert"
)

func TestWalker(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	e := json.NewEncoder(buf)

	buf.WriteString("[")

	for i := 0; i < 1000; i++ {
		if i > 0 {
			buf.WriteString(",")
		}

		err := e.Encode(model.Point{X: rand.Int(), Y: rand.Int()})
		assert.NoError(t, err)
	}

	buf.WriteString("]")

	reader := bytes.NewReader(buf.Bytes())

	walker, err := NewWalker(reader)
	assert.NoError(t, err)

	err = walker.Run(func(chunk []model.Point) error {
		assert.NotEmpty(t, chunk)
		return nil
	})
	assert.NoError(t, err)
}
