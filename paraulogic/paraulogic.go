package paraulogic

import (
	"strings"

	"github.com/adriangalera/lexir/parser"
)

type ParaulogicResult struct {
	Letters       []rune
	Words         map[string]string
	Tutis         []string
	CentralLetter rune
}

var numberOfLetters = 7
var vowels = []rune("aeiou")
var consonants = []rune("bcdfghjklmnpqrstvwxyz")

func GenerateParaulogic(dictionary parser.Dictionary) *ParaulogicResult {
	validCandidate := false
	var result *ParaulogicResult
	for !validCandidate {
		result = generateCandidate(dictionary)
		validCandidate = validate(result)
	}
	return result
}

func validate(result *ParaulogicResult) bool {
	if len(result.Letters) == 0 {
		return false
	}
	if len(result.Words) == 0 {
		return false
	}
	if len(result.Tutis) == 0 {
		return false
	}
	if result.CentralLetter == 0 {
		return false
	}
	return true
}

func pickLetters() *ParaulogicResult {
	result := ParaulogicResult{
		Letters: make([]rune, 0),
		Words:   make(map[string]string),
		Tutis:   make([]string, 0),
	}

	numberOfVowels := randIntInRange(1, numberOfLetters-1)
	numberOfConsonants := numberOfLetters - numberOfVowels

	selectedVowels := selectRandomlyFromRunes(vowels, numberOfVowels)
	selectedConsonants := selectRandomlyFromRunes(consonants, numberOfConsonants)

	letters := append(selectedVowels, selectedConsonants...)
	result.Letters = letters
	result.CentralLetter = selectRandomlyFromRunes(letters, 1)[0]
	return &result
}

func generateCandidate(dictionary parser.Dictionary) *ParaulogicResult {
	result := pickLetters()
	perms := permute(result.Letters, 2, numberOfLetters)

	for _, p := range perms {
		candidate := string(p)
		if dictionary.IsWordInDictionary(candidate) &&
			strings.Contains(candidate, string(result.CentralLetter)) {
			meanings := dictionary.GetMeanings(candidate)
			meaningsStr := meaningsToString((meanings))
			result.Words[candidate] = meaningsStr
			isTuti := isTuti(candidate, result.Letters)
			if isTuti {
				result.Tutis = append(result.Tutis, candidate)
			}
		}
	}
	return result
}

func meaningsToString(meanings []parser.Meaning) string {
	var allSenses []string
	for _, m := range meanings {
		allSenses = append(allSenses, m.Senses...)
	}

	return strings.Join(allSenses, ", ")
}

func isTuti(word string, letters []rune) bool {
	letterSet := make(map[rune]bool)
	for _, r := range word {
		letterSet[r] = true
	}

	for _, r := range letters {
		if !letterSet[r] {
			return false
		}
	}
	return true
}
