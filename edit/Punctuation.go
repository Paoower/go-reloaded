package edit

import (
	"regexp"
	"strings"
)

func EditPunctuation(text string) string {
	// Définir l'expression régulière pour les ponctuations individuelles
	Punc := regexp.MustCompile(`\s*([.,!?:;])\s*`)
	// Définir l'expression régulière pour les ponctuations multiples (par exemple, ... ou !?)
	reMultiplePunc := regexp.MustCompile(`\s*([.]{3}|!?\.{2,}|[!?]{2,})\s*`)

	// Remplacer les ponctuations individuelles par le bon format
	text = Punc.ReplaceAllString(text, "$1 ")

	// Remplacer les ponctuations multiples par le bon format
	text = reMultiplePunc.ReplaceAllStringFunc(text, func(p string) string {
		// Enlever les espaces autour des ponctuations multiples
		return strings.TrimSpace(p)
	})

	// Définir l'expression régulière pour corriger les espaces entre les ponctuations
	reFixSpaces := regexp.MustCompile(`([.,!?:;])\s+([.,!?:;])`)
	// Enlever l'espace entre les ponctuations qui ne devraient pas en avoir
	text = reFixSpaces.ReplaceAllString(text, "$1$2")

	// Enlever les espaces en début et fin de texte
	text = strings.TrimSpace(text)

	return text
}
