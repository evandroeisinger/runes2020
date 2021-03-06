package main

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/runenames"
)

const first, last = ' ', unicode.MaxRune

func main() {
	rangeNames := []string{}
	singleWordChars := []rune{}
	rangeCounts := make(map[string]int)
	fmt.Println("unicode.MaxRune = ", unicode.MaxRune)
	fmt.Println("runenames.UnicodeVersion = ", runenames.UnicodeVersion)
	firstNamed, lastNamed := rune(-1), unicode.MaxRune
	firstUnnamed := rune(-1)
	uniqueCount := 0
	for char := first; char <= last; char++ {
		name := runenames.Name(char)
		if len(name) == 0 {
			if firstUnnamed == rune(-1) {
				firstUnnamed = char
			}
			continue
		}
		rangeCounts[name]++
		if rangeCounts[name] == 2 {
			rangeNames = append(rangeNames, name)
		}
		if name[0] != '<' {
			uniqueCount++ // names that don't start with <
			if firstNamed == -1 {
				firstNamed = char
			}
			lastNamed = char
			if !strings.ContainsRune(name, ' ') {
				singleWordChars = append(singleWordChars, char)
			}
		}
	}
	fmt.Println("Repeated character names (with counts):")
	for _, name := range rangeNames {
		count := rangeCounts[name]
		if count > 1 {
			fmt.Printf("%6d\t%s\n", count, name)
		}
	}
	fmt.Println(strings.Repeat("_", 60))
	fmt.Printf("%6d\tcharacters with unique names\n", uniqueCount)
	fmt.Printf("first:\t%U\t%q\t%s\n", firstNamed, firstNamed, runenames.Name(firstNamed))
	fmt.Printf(" last:\t%U\t%q\t%s\n", lastNamed, lastNamed, runenames.Name(lastNamed))
	fmt.Printf("first character with no name: \t%U\n", firstUnnamed)
	/*
		fmt.Println("Single word character names < U+2FFF:")
		for _, char := range singleWordChars {
			if char < 0x2FFF {
				fmt.Printf("%U\t%c\t%s\n", char, char, runenames.Name(char))
			}
		}
	*/
}
