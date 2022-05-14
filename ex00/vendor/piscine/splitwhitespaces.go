package piscine

func SplitWhiteSpaces(s string) []string {
	return split(s, "\n\n")
}

func replaceRune(s string, before, after rune) string {
	runes := []rune(s)
	result := ""
	for _, r := range runes {
		if r == before {
			result += string(after)
		} else {
			result += string(r)
		}
	}
	return result
}

func mergeWhiteSpaces(s string) string {
	runes := []rune(s)
	result := ""
	flg := false
	for _, r := range runes {
		if flg && r == ' ' {
			continue
		}
		flg = r == ' '

		result += string(r)
	}
	return result
}

func split(str string, sp string) []string {
	result := make([]string, 0)
	index := Index(str, sp)
	for index != -1 {
		result = append(result, removeRune(str[0:index], '\n'))
		str = str[index+2:]
		index = Index(str, sp)
	}
	result = append(result, removeRune(str, '\n'))
	return result
}

func removeRune(s string, rm rune) string {
	runes := []rune(s)
	result := ""
	for _, r := range runes {
		if r == rm {
			continue
		} else {
			result += string(r)
		}
	}
	return result
}

func getByteSize(s string) (size int) {
	for size = range s + " " {
	}
	return size
}

func Index(s string, toFind string) int {
	size_s := getByteSize(s)
	size_f := getByteSize(toFind)
	if size_s < size_f {
		return -1
	}
	for b := range s {
		if toFind == s[b:b+size_f] {
			return b
		}
		if size_s <= size_f+b {
			break
		}
	}
	return -1
}
