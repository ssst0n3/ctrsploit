package util

import "testing"

func TestReplaceELFDependencies(t *testing.T) {
	ReplaceELFDependencies("../helper/dirty-pipe/payload/constructor.so")
}
