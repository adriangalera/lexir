package paraulogic

import (
	"slices"
	"strings"
	"testing"

	"github.com/adriangalera/lexir/parser"
)

func TestPermutations(t *testing.T) {
	min := 2
	max := 4
	permutations := permute([]rune{'a', 'c', 's'}, min, max)

	var strPerms []string
	for _, r := range permutations {
		strPerms = append(strPerms, string(r))
	}

	expected := []string{"casa", "cas", "sac"}

	for _, want := range expected {
		if !slices.Contains(strPerms, want) {
			t.Errorf("Expected permutation %q not found", want)
		}
	}

	for _, combo := range strPerms {
		if len(combo) < min {
			t.Errorf("Found combo %s which is shorter than minimum %d ", combo, min)
		}
		if len(combo) > max {
			t.Errorf("Found combo %s which is larger than maximum %d ", combo, max)
		}
	}
}

func TestTuti(t *testing.T) {
	runes := []rune("maegiln")
	if !isTuti("maligne", runes) {
		t.Error("maligne not detected as tuti!")
	}
	if isTuti("amalgama", runes) {
		t.Error("amalgama detected as tuti!")
	}
}

func TestParaulogicGeneration(t *testing.T) {
	dict := parser.NewHashDictionary()
	err := parser.Parse("../dictionaries/kaikki-spanish.jsonl", dict)
	if err != nil {
		t.Error("Could not parse dictionary")
	}
	paraulogic := GenerateParaulogic(dict)
	if !validate(paraulogic) {
		t.Error("The algorithm has generated an invalid paraulogic")
	}

	for k := range paraulogic.Words {
		letter := string(paraulogic.CentralLetter)
		if !strings.Contains(k, letter) {
			t.Errorf("Word %s does not contain %s\n", k, letter)
		}
	}
}
