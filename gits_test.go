package gits

import "testing"

func TestClone(t *testing.T) {
	err := Clone("git@github.com:JobV/gits.git", "./")
	if err != nil {
		t.Errorf("Can't clone: %v", err)
	}
}
