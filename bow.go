package bow

import (
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/axamon/lemmatizer"
	"github.com/axamon/stringset"
)

var lemmat *lemmatizer.Lemmatizer

func init() {

	initlemmatizer, err := lemmatizer.New("it")

	if err != nil {
		log.Println("lemmatizer non disponibile ", err.Error())
	}
	lemmat = initlemmatizer
	return
}

var alphareg, _ = regexp.Compile(`[^a-z]+`)

// Lemmizza restituisce il lemma di una parola italiana
func Lemmizza(parola string) (lemma string, err error) {

	lemma = lemmat.Lemma(parola)
	if parola == "" {
		log.Printf("Nessuna corrispondenza er %s\n", parola)
	}
	//fmt.Println(lemma)

	return lemma, err

}

// String2Tokens receives a string and outputs a slice of lemmatized italian
// tokens.
func String2Tokens(str string) (tokens map[string]int) {
	tokens = make(map[string]int)

	alphareg, _ := regexp.Compile(`[^a-z]+`)

	lineunica := strings.ToLower(str)
	lineunica = alphareg.ReplaceAllString(lineunica, " ")
	line := strings.Split(lineunica, " ")
	sort.Strings(line)

	for _, field := range line {

		token, err := Lemmizza(field)
		if err != nil {
			continue
		}
		tokens[token]++
	}

	return tokens
}

func String2LemmiSet(str string) (lemmiSet *stringset.StringSet) {

	lemmiSet = stringset.NewStringSet()
	lineunica := strings.ToLower(str)
	lineunica = alphareg.ReplaceAllString(lineunica, " ")
	line := strings.Split(lineunica, " ")
	sort.Strings(line)

	for _, field := range line {

		lemma, err := Lemmizza(field)
		if err != nil {
			continue
		}
		lemmiSet.Add(lemma)
	}

	return lemmiSet
}

// Bow receives a slice of a slice of strings and outputs its tokens map
func Bow(dati [][]string, lung, count int) (tokens map[string]int, NumDocs int) {

	tokens = make(map[string]int)
	NumDocs = len(dati)

	for _, line := range dati {

		lineunica := strings.Join(line, " ")

		lineunica = strings.ToLower(lineunica)

		lineunica = alphareg.ReplaceAllString(lineunica, " ")

		line = strings.Split(lineunica, " ")

		sort.Strings(line)

		for _, field := range line {

			token, err := Lemmizza(field)
			if err != nil {
				continue
			}
			tokens[token]++
		}

		for token, occorrenze := range tokens {
			if occorrenze <= count {
				delete(tokens, token)
			}

			if len(token) <= lung {
				delete(tokens, token)
			}
		}

	}
	return
}
