package piscine

func TrimAtoi(s string) int {
	var neg bool = false
	var empty bool = true
	var res int = 0
	for _, i := range s {
		if empty && i == '-' {
			neg = true
		}
		if i >= '0' && i <= '9' {
			res *= 10
			res += int(i - 48)
			empty = false
		}
	}
	if empty {
		return 0
	} else {
		if neg {
			return -res
		} else {
			return res
		}
	}
}

/*func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}*/
