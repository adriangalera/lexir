package parser

import (
	"reflect"
	"regexp"
	"strings"
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
	var words []string
	regex := strings.ReplaceAll(wordQuery, "*", ".")
	var re = regexp.MustCompile(regex)
	for k := range h.wordMap {
		if len(re.FindStringIndex(k)) > 0 {
			words = append(words, k)
		}
	}
	return words
}
func (h *hashDictionary) GetMeanings(word string) []Meaning {
	if !h.IsWordInDictionary(word) {
		return []Meaning{}
	}
	return h.wordMap[word]
}

func (h *hashDictionary) AllWords() []string {
	keys := make([]string, 0, len(h.wordMap))
	for k := range h.wordMap {
		keys = append(keys, k)
	}
	return keys
}
