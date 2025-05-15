package parser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type entry struct {
	Word              string  `json:"word"`
	Senses            []sense `json:"senses"`
	PartOfSpeech      string  `json:"pos"`
	PartOfSpeechTitle string  `json:"pos_title"`
}

type sense struct {
	Glosses []string `json:"glosses"`
}

func Parse(filename string, dictionary Dictionary) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var entry entry
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			return fmt.Errorf("failed to parse line: %v", err)
		}

		var senses []string
		for _, sense := range entry.Senses {
			senses = append(senses, sense.Glosses...)
		}

		meaning := Meaning{
			PartOfSpeech:      entry.PartOfSpeech,
			PartOfSpeechTitle: entry.PartOfSpeechTitle,
			Senses:            senses,
		}
		err_adding := dictionary.AddWord(entry.Word, meaning)
		if err_adding != nil {
			return err_adding
		}
	}
	return nil
}
