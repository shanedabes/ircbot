package donger

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tdc = dongers{
	"c1": []string{
		"d1",
		"d2",
	},
	"c2": []string{
		"d3",
		"d4",
	},
}

type fakeRandSource struct{}

func (f fakeRandSource) Int63() int64 {
	return 1
}

func (f fakeRandSource) Seed(seed int64) {}

func TestDongerRandomCat(t *testing.T) {
	r := rand.New(fakeRandSource{})

	assert.Equal(t, tdc.randomCat(r), "c1")
}

func TestDongerRandom(t *testing.T) {
	r := rand.New(fakeRandSource{})

	assert.Equal(t, tdc.random(r), "d1")
}

func TestDongerRandomFromCat(t *testing.T) {
	r := rand.New(fakeRandSource{})

	assert.Equal(t, tdc.randomFromCat(r, "c2"), "d3")
}

func TestDongerRandomFromCatError(t *testing.T) {
	r := rand.New(fakeRandSource{})

	assert.Equal(t, tdc.randomFromCat(r, "c3"), "No c3 category")
}
