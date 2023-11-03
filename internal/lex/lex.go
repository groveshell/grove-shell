package lex

func Lex(input string) []string {
	var tokens []string
	current := ""
	var isString, isRawString bool

	for _, char := range input {
		switch char {
		case ' ', '\r', '\t':
			if current != "" {
				if !isString {
					tokens = append(tokens, current)
					current = ""
				} else {
					current += string(char)
				}
			}
		case '"':
			if isRawString {
				current += string(char)
			} else {
				isString = !isString
				current += string(char)
			}
		case '\'':
            if isString {
                current += string(char)
            } else {
                isRawString = !isRawString
                current += string(char)
            }
		case '#':
			if current == "" && !isString {
				return tokens
			} else {
				current += string(char)
			}
		default:
			current += string(char)
		}
	}

	if current != "" {
		tokens = append(tokens, current)
	}

	return tokens
}
