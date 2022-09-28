package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	_ = words

	// Print the byte and rune length of the strings
	// Hint: Use len and utf8.RuneCountInString
	for _, s := range words {
		fmt.Printf("%s: %d bytes, %d runes", s, len(s), utf8.RuneCountInString(s))
		println()
		fmt.Printf("\tbytes   : % x\n", s)
		fmt.Print("\trunes   :")
		for _, r := range s {
			fmt.Printf("% x", r)
		}
		fmt.Println()
		fmt.Print("\trunes   :")
		for _, r := range s {
			fmt.Printf("%q", r)
		}
		fmt.Println()

		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		r, size := utf8.DecodeRuneInString(s)
		fmt.Printf("\tfirst   : %q (%d bytes)\n", r, size)

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		r, size = utf8.DecodeLastRuneInString(s)
		fmt.Printf("\tlast    : %q (%d bytes)\n", r, size)

		// Slice and print the first two runes of the strings
		_, first := utf8.DecodeRuneInString(s)
		_, second := utf8.DecodeRuneInString(s[first:])
		fmt.Printf("\tfirst 2 : %q\n", s[:first+second])

		// Slice and print the last two runes of the strings
		_, last1 := utf8.DecodeLastRuneInString(s)
		_, last2 := utf8.DecodeLastRuneInString(s[:len(s)-last1])
		fmt.Printf("\tlast 2  : %q\n", s[len(s)-last2-last1:])

		// Convert the string to []rune
		// Print the first and last two runes
		rs := []rune(s)
		fmt.Printf("\tfirst 2 : %q\n", string(rs[:2]))
		fmt.Printf("\tlast 2  : %q\n", string(rs[len(rs)-2:]))
	}

}
