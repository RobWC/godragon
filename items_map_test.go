package godragon

import "testing"

func TestItemBuildTree(t *testing.T) {

	err := ItemBuildTree()
	if err != nil {
		t.Fatal(err)
	}

}
