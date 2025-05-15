package parser

import (
	"reflect"
)

type hashDictionary struct {
	wordMap map[string][]Meaning
}

func (h *hashDictionary) AddWord(word string, meaning Meaning) error {
	if h.wordMap == nil {
		h.wordMap = make(map[string][]Meaning)
	}

	var meaningAdded = false
	for _, existing := range h.wordMap[word] {
		if reflect.DeepEqual(existing, meaning) {
			meaningAdded = true
			break
		}
	}

	if !meaningAdded {
		h.wordMap[word] = append(h.wordMap[word], meaning)
	}

	return nil
}
func (h *hashDictionary) IsWordInDictionary(word string) bool {
	_, ok := h.wordMap[word]
	return ok
}
func (h *hashDictionary) FindAllWordsMatching(wordQuery string) []string {
	return []string{}
}
func (h *hashDictionary) GetMeanings(word string) []Meaning {
	if !h.IsWordInDictionary(word) {
		return []Meaning{}
	}
	return h.wordMap[word]
}
