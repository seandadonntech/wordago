package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var amount int
	var createFile bool
	rand.Seed(time.Now().UnixNano())

	fmt.Println("ENTER THE NUMBER OF WORDS TO GENERATE:")
	fmt.Scan(&amount)

	var includeUppercase, includeLowercase, includeNumbers, includeSymbols bool
	fmt.Println("Include Uppercase letters? (true/false):")
	fmt.Scan(&includeUppercase)
	fmt.Println("Include Lowercase letters? (true/false):")
	fmt.Scan(&includeLowercase)
	fmt.Println("Include Numbers? (true/false):")
	fmt.Scan(&includeNumbers)
	fmt.Println("Include Symbols? (true/false):")
	fmt.Scan(&includeSymbols)
	fmt.Println("CREATE FILE? (true/false):")
	fmt.Scan(&createFile)

	charSets := buildCharSets(includeUppercase, includeLowercase, includeNumbers, includeSymbols)

	var words []string
	for i := 0; i < amount; i++ {
		word := generateRandomWord(charSets)
		words = append(words, word)
	}

	if createFile {
		file, err := os.Create("generated_words.txt")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		for _, word := range words {
			_, err := file.WriteString(word + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}

		fmt.Println("Generated words have been written to 'generated_words.txt'.")
	} else {
		fmt.Println("Generated Words:")
		for _, word := range words {
			fmt.Println(word)
		}
	}
}

func buildCharSets(uppercase, lowercase, numbers, symbols bool) []string {
	charSets := []string{}
	if uppercase {
		charSets = append(charSets, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	if lowercase {
		charSets = append(charSets, "abcdefghijklmnopqrstuvwxyz")
	}
	if numbers {
		charSets = append(charSets, "1234567890")
	}
	if symbols {
		charSets = append(charSets, "!@#$%&'()*+")
	}

	return charSets
}

func generateRandomWord(charSets []string) string {
	var word string
	for i := 0; i < 10; i++ {
		charSet := charSets[rand.Intn(len(charSets))]
		randomChar := charSet[rand.Intn(len(charSet))]
		word += string(randomChar)
	}

	return word
}
