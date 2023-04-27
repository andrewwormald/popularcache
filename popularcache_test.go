package popularcache_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"popularcache"
)

// TestCache is a general acceptance test for the package
func TestCache(t *testing.T) {
	type item struct {
		ID    string
		Value string
	}

	c := popularcache.New[item]()
	c.Add("1", item{ID: "1", Value: "item 1"})
	c.Add("2", item{ID: "2", Value: "item 2"})
	c.Add("3", item{ID: "3", Value: "item 3"})
	c.Add("4", item{ID: "4", Value: "item 4"})

	// Calling collect moves the item to the front
	c.Collect("2")

	require.Equal(t, []item{
		{ID: "2", Value: "item 2"},
		{ID: "4", Value: "item 4"},
		{ID: "3", Value: "item 3"},
		{ID: "1", Value: "item 1"},
	}, c.List())
}
