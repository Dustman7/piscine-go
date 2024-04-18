package piscine

func SplitWhiteSpaces(s string) []string {
	var out []string = []string{}
	var current string = ""

	for _, char := range s {
		if char == '\n' || char == '\t' || char == ' ' {
			if current != "" {
				out = append(out, current)
				current = ""
			}
		} else {
			current += string(char)
		}
	}

	if current != "" {
		out = append(out, current)
	}

	return out
}
