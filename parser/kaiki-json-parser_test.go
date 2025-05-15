package parser

import (
	"testing"

	"github.com/warpfork/go-wish/cmp"
)

func TestHashMapDictionary(t *testing.T) {

	word := "japonés"
	expectedMeanings := []Meaning{
		{PartOfSpeech: "adj", PartOfSpeechTitle: "Adjetivo", Senses: []string{"Originario, relativo a, o propio de Japón."}},
	}

	dict := hashDictionary{}
	error_parsing := Parse("./kaikki-test.json", &dict)
	if error_parsing != nil {
		t.Errorf("Error parsing!: %v", error_parsing)
	}

	if !dict.IsWordInDictionary(word) {
		t.Errorf("Word %s not found in dictionary", word)
	}

	meanings := dict.GetMeanings(word)

	diff := cmp.Diff(meanings, expectedMeanings)
	if diff != "" {
		t.Errorf("Found some difference! Difference: %s", diff)
	}
}
