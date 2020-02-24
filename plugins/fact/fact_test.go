package fact

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	j = json{
		Text: "fact",
	}
)

func TestString(t *testing.T) {
	assert.Equal(t, j.String(), "fact")
}
