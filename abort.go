package piscine

func Abort(a, b, c, d, e int) int {
	if a > b {
		u := a
		a = b
		b = u
	}
	if a > c {
		u := a
		a = c
		c = u
	}
	if a > d {
		u := a
		a = d
		d = u
	}
	if a > e {
		u := a
		a = e
		e = u
	}
	if b > c {
		u := b
		b = c
		c = u
	}

	if b > d {
		u := b
		b = d
		d = u
	}
	if b > e {
		u := b
		b = e
		e = u
	}
	if c > d {
		u := c
		c = d
		d = u
	}
	if c > e {
		u := c
		c = e
		e = u
	}
	if d > e {
		u := d
		d = e
		e = u
	}

	return c
}
