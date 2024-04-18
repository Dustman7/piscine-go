package piscine

func Compact(ptr *[]string) int {
	var s []string
	i := 0
	for _, tfef := range *ptr {
		if tfef != "" {
			s = append(s, tfef)
			i++
		}
	}
	*ptr = s

	return i
}
