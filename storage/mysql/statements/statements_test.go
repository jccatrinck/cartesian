package statements

import (
	"regexp"
	"testing"

	"github.com/pingcap/parser"
	_ "github.com/pingcap/parser/test_driver"
	"github.com/stretchr/testify/assert"
)

var bindRemover = regexp.MustCompile(`\:[a-z]+`)

func TestStatements(t *testing.T) {
	statements := []string{
		LoadPoints,
		PointsByDistancePK,
		PointsByDistanceSpatial,
	}

	for _, statement := range statements {
		statement = bindRemover.ReplaceAllString(statement, "?")

		_, warnings, err := parser.New().Parse(statement, "", "")

		assert.NoError(t, err)
		assert.Empty(t, warnings, "There are warnings %v", warnings)
	}
}
