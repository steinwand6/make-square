package piscine

func ConvMino(m []rune) [][]int {
	results := make([][]int, 4)
	count := 0
	base_x, base_y := find_top_in_mino(m)
	x, y := base_x, base_y
	for y < 4 {
		x = base_x
		for x < 4 {
			if count > 3 {
				return results
			}
			if m[y*4+x] != '.' {
				results[count] = []int{x - base_x, y - base_y}
				count++
			}
			x++
		}
		y++
	}
	return results
}

func find_top_in_mino(m []rune) (int, int) {
	y := 0
	for y < 4 {
		x := 0
		for x < 4 {
			if m[x+4*y] != '.' {
				return x, y
			}
			x++
		}
		y++
	}
	return 0, 0
}
