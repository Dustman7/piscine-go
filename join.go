package piscine

func Join(strs []string, sep string) string {
	concat := ""
	for i, sac := range strs {
		concat += sac
		if i != len(strs)-1 {
			concat += sep
		}
	}
	return concat
}
