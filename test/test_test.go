package test_test

import (
	"testing"

	"github.com/pucora/httpcache"
	"github.com/pucora/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
