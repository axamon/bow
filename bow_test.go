package bow_test

import (
	"fmt"

	"github.com/axamon/bow"
)

func ExampleString2LemmiSet() {
	lemmi := bow.String2LemmiSet("Salve, il cliente lamenta problemi. Il Cliente è infuriato")
	for _, lemma := range lemmi.Strings() {
		fmt.Println(lemma)
	}
	// Unordered output:
	// infuriare
	// cliente
	// il
	// lamentare
	// problema
	// salva
}

func ExampleString2Tokens() {
	tokens := bow.String2Tokens("Salve, il cliente lamenta problemi. Il Cliente è infuriato")
	for token := range tokens {
		fmt.Println(token)
	}
	// Unordered output:
	// infuriare
	// cliente
	// il
	// lamentare
	// problema
	// salva
}
