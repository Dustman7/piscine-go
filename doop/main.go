package main

import (
	"os"
)

var (
	INT64_MAX int64 = 9223372036854775807
	INT64_MIN int64 = -9223372036854775808
)

// ERRORS
//  0 = NO ERROR
//  1 = PARSE ERROR
//  2 = OVERFLOW ERROR
//  3 = UNDERFLOW ERROR

func IntToString(n int64) string {
	var number []rune
	var negative bool = false

	if n == 0 {
		number = append(number, '0')
	}

	if n < 0 {
		n = -n
		negative = true
	}

	for nb := n; nb > 0; nb /= 10 {
		digit := (nb % 10) + '0'
		number = append(number, rune(digit))
	}

	if negative {
		number = append(number, '-')
	}

	for i, j := 0, len(number)-1; i < j; i, j = i+1, j-1 {
		number[i], number[j] = number[j], number[i]
	}

	return string(number)
}

func ParseValue(arg string, status *int) int64 {
	var out int64 = 0
	var negative bool = false
	var start int = 0

	runes := [](rune)(arg)
	length := len(runes)

	if runes[0] == '-' {
		negative = true
		start = 1
	}

	for i := start; i < length; i++ {
		if runes[i] < '0' || runes[i] > '9' {
			*status = 1
			return -1
		}

		digit := int64(runes[i] - '0')

		if negative {
			if out >= 0 {
				out = (-out * 10)
				out = Sub64Safe(out, digit, status)
			} else {
				out = (out * 10)
				out = Add64Safe(out, -digit, status)
			}

			if *status != 0 {
				return -1
			}
		} else {
			out = (out * 10)
			out = Add64Safe(out, digit, status)
			if *status != 0 {
				return -1
			}
		}
	}

	return out
}

func ParseOperator(arg string, status *int) rune {
	runes := [](rune)(arg)
	if len(runes) != 1 {
		*status = 1
		return -1
	}
	first := runes[0]
	return first
}

func Add64Safe(left int64, right int64, status *int) int64 {
	// Check if the right is a positive or negative number
	// Choose to check integers overflows or underflows
	if right > 0 {
		// Check overflows
		if left > INT64_MAX-right {
			*status = 2
			return -1
		}
	} else {
		// Check underflows
		if left < INT64_MIN-right {
			*status = 3
			return -1
		}
	}
	return (left + right)
}

func Sub64Safe(left int64, right int64, status *int) int64 {
	// Check if the right is a positive or negative number
	// Choose to check integers overflows or underflows
	if right > 0 {
		// Check underflows
		if left < INT64_MIN+right {
			*status = 2
			return -1
		}
	} else {
		// Check overflows
		if left > INT64_MAX+right {
			*status = 3
			return -1
		}
	}
	return (left - right)
}

func Mul64Safe(left int64, right int64, status *int) int64 {
	x := left * right
	if left != 0 && x/left != right {
		*status = 3
	}
	return x
}

func PrintStr(s string) {
	os.Stdout.Write([]byte(s))
}

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		return
	}

	v1s, ops, v2s := 0, 0, 0

	v1 := ParseValue(args[0], &v1s)
	op := ParseOperator(args[1], &ops)
	v2 := ParseValue(args[2], &v2s)

	if v1s != 0 || ops != 0 || v2s != 0 {
		return
	}

	var status int = 0
	var v0 int64 = 0

	if op == '*' {
		v0 = Mul64Safe(v1, v2, &status)
		if status != 0 {
			return
		}
	} else if op == '/' {
		if v2 == 0 {
			PrintStr("No division by 0\n")
			return
		}
		v0 = v1 / v2
	} else if op == '+' {
		v0 = Add64Safe(v1, v2, &status)
		if status != 0 {
			return
		}
	} else if op == '-' {
		v0 = Sub64Safe(v1, v2, &status)
		if status != 0 {
			return
		}
	} else if op == '%' {
		if v2 == 0 {
			PrintStr("No modulo by 0\n")
			return
		}
		v0 = v1 % v2
	} else {
		return
	}

	os.Stdout.Write([]byte(IntToString(v0)))
	os.Stdout.Write([]byte(string('\n')))
}
