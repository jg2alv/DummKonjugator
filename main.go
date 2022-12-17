package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Willkommen aus DummKonjugator.\nBitte, tippen Sie ein Verb: ")
	reader := bufio.NewReader(os.Stdin)
	verbs, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	if strings.Contains(verbs, ",") {
		for _, verb := range strings.Split(verbs, ",") {
			conjugate(verb)
		}

		return
	}
	
	conjugate(verbs)
}

func conjugate(verb string) {
	verb = strings.Trim(verb, " ")
	var base string

	if strings.Contains(verb, "en\n") {
		base = strings.TrimSuffix(verb, "en\n")
	} else if strings.Contains(verb, "en") {
		base = strings.TrimSuffix(verb, "en")
	} else {
		fmt.Printf("Das Verb in '%v' wurde nicht gefunden!", verb)
		return
	}

	fmt.Printf("ich        %ve\n", base)
	fmt.Printf("du         %vst\n", base)
	fmt.Printf("Sie        %ven\n", base)
	fmt.Printf("er/sie/es  %vt\n", base)
	fmt.Printf("wir        %ven\n", base)
	fmt.Printf("ihr        %vt\n", base)
	fmt.Printf("Sie        %ven\n", base)
	fmt.Printf("sie        %ven\n\n", base)
}
