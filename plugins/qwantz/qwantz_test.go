package qwantz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testComic(t *testing.T) {
	c := comic{
		url:  "test.com",
		desc: "comic desc",
	}

	expected := "comic desc - test.com"

	assert.Equal(t, c.String(), expected)
}
