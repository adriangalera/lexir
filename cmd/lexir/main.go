package main

import (
	"fmt"
	"log"

	"github.com/adriangalera/lexir/parser"
	"github.com/spf13/cobra"
)

func main() {
	var inputFile string

	var rootCmd = &cobra.Command{
		Use:   "lexir",
		Short: "Lexir generation",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Loading file: %s...\n", inputFile)
			dict := parser.NewHashDictionary()
			errorParsing := parser.Parse(inputFile, dict)
			if errorParsing != nil {
				log.Fatalf("Could not parse %s", inputFile)
			} else {
				fmt.Printf("Successfully loaded %d words\n", len(dict.AllWords()))
			}
		},
	}

	var err error

	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file path")
	err = rootCmd.MarkFlagRequired("input")
	if err != nil {
		log.Fatalf("Could not initialize. Error: %s", err)
	}

	err = rootCmd.Execute()
	if err != nil {
		log.Fatalf("Could not initialize. Error: %s", err)
	}
}
