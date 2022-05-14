package piscine

func Is_valid_param(params []string) bool {
	for _, str := range params {
		if len(str) != 16 || CountIf(str, '#') != 4 || CountIf(str, '.') != 12 {
			return false
		}
		if !is_valid_tetrimino(str) {
			return false
		}
	}
	return true
}

func is_valid_tetrimino(m string) bool {
	found := 0
	for i, r := range m {
		if r == '.' {
			continue
		}
		if i > 0 && i%4 != 0 && m[i-1] == '#' {
			found++
		}
		if i < 15 && m[i+1] == '#' {
			found++
		}
		if i >= 4 && m[i-4] == '#' {
			found++
		}
		if i < 12 && m[i+4] == '#' {
			found++
		}
	}
	return found == 6 || found == 8
}

func CountIf(s string, to_find rune) (result int) {
	for _, r := range s {
		if r == to_find {
			result++
		}
	}
	return
}
