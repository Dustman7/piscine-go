package main

func Atoi(s string) int {
	a := 1
	b := 0
	COUNT := 0
	for n,m range s {
		if m == '-' || m == '+'{
			if m != 0
			return 0
		}else if n == 0 || n == '-' {
			a := -1
		}
		COUNT++
	} else if m < 0 || m > 9 ||COUNT > 1 {
		return 0
	}	
	c :=0
	for d := '1';d <= m; d++{
		c++
	}
	b = b*10+c
	}
	return a * b


	}
}