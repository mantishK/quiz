package tree

import "testing"

func TestTree(t *testing.T) {
	words := []string{
		`aa`,
		`aah`,
		`aahed`,
		`aahing`,
		`aahs`,
		`aal`,
		`aalii`,
		`aaliis`,
		`aals`,
		`aardvark`,
		`aardvarks`,
		`aardwolf`,
		`aardwolves`,
		`aargh`,
		`aarrgh`,
		`aarrghh`,
		`a`,
	}
	tree := Tree{char: rune('a')}
	for _, word := range words {
		tree.AddWord(word)
	}
	if !tree.WordExists("aal") {
		t.Error("aal")
	}
	if tree.WordExists("aalsd") {
		t.Error("aalsd")
	}
	if tree.WordExists("j") {
		t.Error("j")
	}
	if tree.WordExists("aat") {
		t.Error("aat")
	}
	if tree.WordExists("a") {
		t.Error("aat")
	}
}
