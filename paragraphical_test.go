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

func Test_LongWords(t *testing.T) {
	text := `This is another really long sentence with an additionally1234567890123 long word.`

	expected := `This is another
really long sentence
with an
additionally1234567890123
long word.`

	assert.Equal(t, expected, Format(20, text))
}

func Test_Multiline(t *testing.T) {
	text := `This is some more super long text but
this time it's broken over
several lines.

Here is some more text to be wrapped.`

	expected := `This is some more
super long text but
this time it's
broken over
several lines.

Here is some more
text to be wrapped.`

	assert.Equal(t, expected, Format(20, text))
}
