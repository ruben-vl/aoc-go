package stringconv

func DigitsFromString(s string) []int {
	out := make([]int, 0, len(s))
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			out = append(out, int(ch-'0'))
		}
	}
	return out
}
