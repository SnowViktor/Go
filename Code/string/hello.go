package string

const (
	EnglishHelloPrefix = "Hello, "
	SpanishHelloPrefix = "Hola, "
	ChineseHelloPrefix = "哈囉，"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return GreetingPrefix(language) + name
}

func GreetingPrefix(language string) (prefix string) {
	switch language {
	case "English":
		prefix = EnglishHelloPrefix
	case "Spanish":
		prefix = SpanishHelloPrefix
	case "Chinese":
		prefix = ChineseHelloPrefix
	default:
		prefix = EnglishHelloPrefix
	}
	return
}
