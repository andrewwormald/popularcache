package popularcache_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/andrewwormald/popularcache"
)

// TestCache is a general acceptance test for the package
func TestCacheCollect(t *testing.T) {
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

func TestCacheAdd(t *testing.T) {
	type item struct {
		ID    string
		Value string
	}

	c := popularcache.New[item]()
	c.Add("1", item{ID: "1", Value: "item 1"})
	c.Add("2", item{ID: "2", Value: "item 2"})
	c.Add("3", item{ID: "3", Value: "item 3"})
	c.Add("4", item{ID: "4", Value: "item 4"})

	require.Equal(t, []item{
		{ID: "4", Value: "item 4"},
		{ID: "3", Value: "item 3"},
		{ID: "2", Value: "item 2"},
		{ID: "1", Value: "item 1"},
	}, c.List())
}

func TestCacheTrimLast(t *testing.T) {
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

	// Calling TrimRight removes 1 item from the very end
	c.TrimRight(1)

	require.Equal(t, []item{
		{ID: "2", Value: "item 2"},
		{ID: "4", Value: "item 4"},
		{ID: "3", Value: "item 3"},
	}, c.List())

	c.TrimRight(1)

	require.Equal(t, []item{
		{ID: "2", Value: "item 2"},
		{ID: "4", Value: "item 4"},
	}, c.List())

	c.Add("5", item{ID: "5", Value: "item 5"})
	c.Add("6", item{ID: "6", Value: "item 6"})
	c.Add("7", item{ID: "7", Value: "item 7"})

	c.TrimRight(1)

	require.Equal(t, []item{
		{ID: "7", Value: "item 7"},
		{ID: "6", Value: "item 6"},
		{ID: "5", Value: "item 5"},
		{ID: "2", Value: "item 2"},
	}, c.List())
}
