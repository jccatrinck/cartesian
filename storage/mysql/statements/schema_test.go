package statements

import (
	"strings"
	"testing"

	_ "github.com/pingcap/parser/test_driver"
	"github.com/stretchr/testify/assert"
)

func TestSchema(t *testing.T) {
	assert.NotEmpty(t, strings.TrimSpace(Schema))
}
