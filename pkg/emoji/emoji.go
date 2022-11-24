package emoji

import (
	ea "emojiaffixer/pkg/affix"
	"fmt"
	"regexp"
	"strings"
)

var alphabets = regexp.MustCompile(`[a-zA-Z]`)
var symbols = regexp.MustCompile(`[!|?]`)
var digits = regexp.MustCompile(`[0-9]`)

func match(regexp *regexp.Regexp, char string) bool {
	return regexp.MatchString(char)
}

func GetEmojiExpression(char string, props ea.Property) string {
	var (
		prefix = props.Prefix
		suffix = props.Suffix
	)

	if char == " " {
		return char
	}

	var emoji string

	if match(alphabets, char) {
		emoji = strings.ToLower(char)
	}

	if match(symbols, char) {
		emoji = translateSymbol(char)
	}

	if match(digits, char) {
		emoji = translateDigit(char)
	}

	if emoji != "" {
		return buildEmoji(emoji, prefix, suffix)
	}

	// There's no matching, returns char as is
	return buildEmoji(char, prefix, suffix)
}

func translateSymbol(s string) string {
	switch s {
	case "!":
		return "exclamation"
	case "?":
		return "question"
	default:
		panic(fmt.Sprintf("Expected: Symbol. Actual: %s.", s))
	}
}

func translateDigit(d string) string {
	switch d {
	case "1":
		return "one"
	case "2":
		return "two"
	case "3":
		return "three"
	case "4":
		return "four"
	case "5":
		return "five"
	case "6":
		return "six"
	case "7":
		return "seven"
	case "8":
		return "eight"
	case "9":
		return "nine"
	case "0":
		return "zero"
	default:
		panic(fmt.Sprintf("Expected: Digit. Actual: %s.", d))
	}
}

func buildEmoji(char string, prefix string, suffix string) string {
	return fmt.Sprintf(":%s%s%s:", prefix, char, suffix)
}
