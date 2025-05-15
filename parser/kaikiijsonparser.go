package parser

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"runtime"
	"sync"
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

	// Parses JSON lines from input concurrently using a worker pool, extracts word meanings,
	// and adds them to the dictionary. Synchronizes access to the shared dictionary using a mutex.

	var mu sync.Mutex
	var wg sync.WaitGroup
	//100 is the channel buffer size. It means the channel can hold up to 100 strings without blocking.
	jobs := make(chan string, 100)

	// starts a worker goroutine for each CPU core available on the machine.
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for line := range jobs {
				var entry entry
				if err := json.Unmarshal([]byte(line), &entry); err != nil {
					log.Printf("Failed to parse line: %v", err)
					continue
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
				/*
					The parser does not know about the implementation details of the dictionary.
					In order to add words from multiple goroutines, the parser uses a mutex to ensure
					thread safety while adding a word
				*/
				mu.Lock()
				err := dictionary.AddWord(entry.Word, meaning)
				mu.Unlock()
				if err != nil {
					log.Printf("Failed to add word: %v", err)
				}
			}
		}()
	}

	for scanner.Scan() {
		line := scanner.Text()
		jobs <- line
	}
	close(jobs)
	wg.Wait()

	return nil
}
