package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const Len = "len"
const Amount = "amount"

type Check struct {
	Title string
	Min   int
	Max   int
}

func main() {
	checks := []Check{
		{Title: Len, Min: 6, Max: 128},
		{Title: Amount, Min: 0, Max: 1000},
	}
	attributes := GeAttributes(checks)
	passwords := Generate(attributes)
	fmt.Println("Your passwords:", passwords)
}

func GeAttributes(checks []Check) (attributes map[string]int) {
	for _, check := range checks {
		entry := 0
		for {
			fmt.Printf("Enter %s of your passwords\n", check.Title)
			_, err := fmt.Scanf("%d", &entry)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			if entry < check.Min || entry > check.Max {
				fmt.Printf(
					"%s of the password should be between %d and %d!\n",
					strings.Title(check.Title),
					check.Min,
					check.Max,
				)
				continue
			} else {
				if attributes == nil {
					attributes = make(map[string]int)
				}
				attributes[check.Title] = entry
				break
			}
		}
	}

	return attributes
}

func Generate(pAttributes map[string]int) (passwords []string) {
	for i := 0; i < pAttributes[Amount]; i++ {
		var runes []rune
		for j := 0; j < pAttributes[Len]; j++ {
			r := rune(rand.Intn(94) + 33)
			runes = append(runes, r)
		}
		passwords = append(passwords, string(runes))
	}

	return passwords
}