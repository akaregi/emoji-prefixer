package main

import (
	ea "emojiaffixer/pkg/affix"
	em "emojiaffixer/pkg/emoji"
	pr "emojiaffixer/pkg/printer"
	"flag"
	"fmt"
	"strings"
)

func main() {
	prefix := flag.String("p", "", "Prefix to be attached to message.")
	suffix := flag.String("s", "", "Suffix to be attached to message.")

	// NOTE: delimiter's default value is ZWSP(Zero Witch Space).
	delimiter := flag.String("d", "​", "Delimiter to concatenate the text")
	verbose := flag.Bool("v", false, "Verbose mode?")

	flag.Parse()

	affixProps := ea.NewAffixProps(prefix, suffix)
	printer := pr.NewPrinterProps(verbose, delimiter)

	// GO GO GO
	var text = strings.Split(strings.Join(flag.Args(), " "), "")
	var emojis []string

	for _, t := range text {
		emoji := em.GetEmojiExpression(t, affixProps)
		emojis = append(emojis, emoji)
	}

	// TODO: Implement this line on printer
	fmt.Printf("%s\n", strings.Join(emojis, printer.Delimiter))
}
