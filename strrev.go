package piscine

func StrRev(s string) (Result string) { // Result can be set after return key word
	for _, t := range s {
		// Result if not present on the bottom line, will only count one letter every time */
		Result = string(t) + Result
	}
	return /* HERE */
}
