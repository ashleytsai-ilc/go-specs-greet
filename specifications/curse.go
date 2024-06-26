package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func CurseSpecification(t testing.TB, meany MeanGreeter) {
	got, err := meany.Curse("Chris")
	assert.NoError(t, err)
	assert.Equal(t, "Go to hell, Chris!", got)
}
