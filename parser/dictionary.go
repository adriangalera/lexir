package parser

type Meaning struct {
	PartOfSpeech      string
	PartOfSpeechTitle string
	Senses            []string
}

type Dictionary interface {
	AddWord(word string, meaning Meaning) error
	IsWordInDictionary(word string) bool
	FindAllWordsMatching(wordQuery string) []string
	GetMeanings(word string) []Meaning
}

func NewHashDictionary() Dictionary {
	return &hashDictionary{}
}
