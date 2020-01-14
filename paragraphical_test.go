package paragraphical

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SimpleWrap(t *testing.T) {
	text := `This is a fairly long string that needs to be broken in to smaller lines in the terminal.`

	expected20 := `This is a fairly
long string that
needs to be broken
in to smaller lines
in the terminal.`
	assert.Equal(t, expected20, Format(20, text))

	expected30 := `This is a fairly long string
that needs to be broken in to
smaller lines in the terminal.`
	assert.Equal(t, expected30, Format(30, text))
}
