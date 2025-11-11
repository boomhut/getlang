package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/boomhut/getlang"
	// When using this example code in your own project,
	// you should import getlang with:
)

// This example demonstrates how to use getlang to detect
// the language of a text provided via command line arguments.
//
// To use this example in your own project:
// 1. Import the getlang package: import "github.com/boomhut/getlang"
// 2. Use getlang.FromString() to detect language
// 3. Access language info via the returned Info object

func main() {
	// Check if command line arguments were provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: getlang \"Text to detect language\"")
		fmt.Println("Example: getlang \"Hello world\"")
		fmt.Println("\nSupported languages include:")
		fmt.Println("  English, Spanish, French, German, Italian, Portuguese")
		fmt.Println("  Russian, Ukrainian, Arabic, Hindi, Chinese, Japanese")
		fmt.Println("  And 30+ more languages (see LANGUAGES.md for full list)")
		os.Exit(1)
	}

	// Combine all arguments into a single string
	text := strings.Join(os.Args[1:], " ")
	// Detect the language
	info := getlang.FromString(text)

	// Display the results
	fmt.Println("=====================================================")
	fmt.Println("Text analyzed:")
	fmt.Printf("\"%s\"\n", text)
	fmt.Println("=====================================================")
	fmt.Printf("Detected Language: %s\n", info.LanguageName())
	fmt.Printf("ISO 639-1 Code: %s\n", info.LanguageCode())
	fmt.Printf("Native Name: %s\n", info.SelfName())
	fmt.Printf("Confidence: %.2f%%\n", info.Confidence()*100)
	fmt.Println("======================================================")
}
