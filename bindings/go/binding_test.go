package tree_sitter_openfga_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_openfga "github.com/sebb3/tree-sitter-openfga/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_openfga.Language())
	if language == nil {
		t.Errorf("Error loading OpenFGA grammar")
	}
}
