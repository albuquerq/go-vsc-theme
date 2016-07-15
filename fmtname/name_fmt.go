package fmtname

import (
	"bytes"
	"strings"
	"unicode"
)

// Normalize normaliza File Names para Display Names.
func Normalize(s string) string {
	nome := strings.Replace(s, "-", " ", -1)
	nome = strings.Replace(nome, "_", " ", -1)
	nome = strings.TrimSpace(nome)
	nome = SplitCamelCase(nome)
	nome = strings.Title(nome)
	return nome
}

// SplitCamelCase separa palavras de nomes em camelCase, ou PascalCase com espaços.
func SplitCamelCase(s string) string {
	b := bytes.Buffer{}
	for i, letra := range s {
		if unicode.IsUpper(letra) {
			if i < len(s)-1 {
				if unicode.IsLower(rune(s[i+1])) {
					if i > 1 {
						if s[i-1] != ' ' {
							b.WriteRune(' ')
						}
					}
				}
			}
		}
		b.WriteRune(letra)
	}
	return b.String()
}

// ToLowerJoin unifica palavras separadas por espaço e transforma para lowercase.
func ToLowerJoin(s string) string {
	nome := strings.ToLower(s)
	nome = strings.Replace(nome, " ", "", -1)
	return nome
}
