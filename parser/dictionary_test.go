package parser

import (
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestOnlyOneEntry(t *testing.T) {
	word := "japonés"
	expectedMeanings := []Meaning{
		{PartOfSpeech: "adj", PartOfSpeechTitle: "Adjetivo", Senses: []string{"Originario, relativo a, o propio de Japón."}},
	}

	implementations := []Dictionary{NewHashDictionary()}

	for _, dict := range implementations {
		error_parsing := Parse("./kaikki-test.json", dict)
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
}

func TestSameWordWithDifferentMeanings(t *testing.T) {
	word := "japonés"
	expectedMeanings := []Meaning{
		{PartOfSpeech: "noun", PartOfSpeechTitle: "Sustantivo masculino y femenino", Senses: []string{"Persona originaria de Japón."}},
		{PartOfSpeech: "adj", PartOfSpeechTitle: "Adjetivo", Senses: []string{"Originario, relativo a, o propio de Japón."}},
		{PartOfSpeech: "noun", PartOfSpeechTitle: "Sustantivo masculino", Senses: []string{"Lengua de filiación incierta, hablada en el Japón y partes de Palau."}},
	}

	implementations := []Dictionary{NewHashDictionary()}

	for _, dict := range implementations {
		error_parsing := Parse("./kaikki-test2.json", dict)
		if error_parsing != nil {
			t.Errorf("Error parsing!: %v", error_parsing)
		}

		if !dict.IsWordInDictionary(word) {
			t.Errorf("Word %s not found in dictionary", word)
		}

		meanings := dict.GetMeanings(word)

		// Sort the slices while comparing because the order is not guarantee due to the concurrency
		diff := cmp.Diff(meanings, expectedMeanings, cmpopts.SortSlices(func(a, b Meaning) bool {
			if a.PartOfSpeech != b.PartOfSpeech {
				return a.PartOfSpeech < b.PartOfSpeech
			}
			if a.PartOfSpeechTitle != b.PartOfSpeechTitle {
				return a.PartOfSpeechTitle < b.PartOfSpeechTitle
			}
			return len(a.Senses) > 0 && len(b.Senses) > 0 && a.Senses[0] < b.Senses[0]
		}))

		if diff != "" {
			t.Errorf("Found some difference! Difference: %s", diff)
		}
	}
}

func TestOneEntryWithDuplicatedMeaning(t *testing.T) {
	word := "japonés"
	expectedMeanings := []Meaning{
		{PartOfSpeech: "adj", PartOfSpeechTitle: "Adjetivo", Senses: []string{"Originario, relativo a, o propio de Japón."}},
	}

	implementations := []Dictionary{NewHashDictionary()}

	for _, dict := range implementations {
		error_parsing := Parse("./kaikki-test3.json", dict)
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
}

func TestFindWordsMatching(t *testing.T) {
	word := "japonés"

	implementations := []Dictionary{NewHashDictionary()}
	for _, dict := range implementations {
		error_parsing := Parse("./kaikki-test.json", dict)
		if error_parsing != nil {
			t.Errorf("Error parsing!: %v", error_parsing)
		}
		criteria := "j*pon*s"
		wordsMatching := dict.FindAllWordsMatching(criteria)

		if !slices.Contains(wordsMatching, word) {
			t.Errorf("Word %s not found by matching", word)
		}
	}
}

func TestNotFindWordsMatching(t *testing.T) {
	word := "austral"

	implementations := []Dictionary{NewHashDictionary()}
	for _, dict := range implementations {
		error_parsing := Parse("./kaikki-test.json", dict)
		if error_parsing != nil {
			t.Errorf("Error parsing!: %v", error_parsing)
		}
		criteria := "a*stral"
		wordsMatching := dict.FindAllWordsMatching(criteria)

		if slices.Contains(wordsMatching, word) {
			t.Errorf("Word %s found when it's not in the dictionary", word)
		}
	}
}
