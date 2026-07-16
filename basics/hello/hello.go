package main

import (
	"fmt"
)

// constants
const (
	englishPrefix    = "hello, "
	portuguesePrefix = "ola, "
	spanishPrefix    = "hola, "
	germanyPrefix    = "hallo, "
	frenchPrefix     = "bonjour, "
	italianPrefix    = "ciao, "
)

// domain code -> o que de fato interessa
func hello(name string, lang string) string {
	if name == "" {
		return englishPrefix + "world"
	}
	return manageLang(lang) + name
}

func manageLang(lang string) string {
	switch lang {
	case "english":
		return englishPrefix
	case "portuguese":
		return englishPrefix
	case "spanish":
		return englishPrefix
	case "germany":
		return englishPrefix
	case "french":
		return englishPrefix
	case "italian":
		return englishPrefix
	default:
		return ""
	}
}

func helper(name string, lang string) string {
	if name == "" {
		return "what's your name?"
	} else if lang == "" {
		return "what language do you speak?"
	} else if lang == "" && name == "" {
		return "troll"
	} else {
		return ""
	}
}

func main() {
	fmt.Println(hello("ghu", "portuguese"))
}
