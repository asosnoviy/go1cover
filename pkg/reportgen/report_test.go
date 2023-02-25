package reportgen

import (
	"os"
	"testing"

	coveragedata "github.com/asosnoviy/go1cover/internal/coverageData"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	generators := NewReporters([]string{"lcov", "generic"})
	generator := New(generators...)
	assert.EqualValues(t, generator.reporters, generators)

	generator.Report(&coveragedata.Coverage{})

	assert.FileExists(t, "lcov.info")
	assert.FileExists(t, "generic.xml")

	defer os.Remove("lcov.info")
	defer os.Remove("generic.xml")
}
