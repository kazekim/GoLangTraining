package stringutil

func reverseTwo(s string) string {
	r := []rune(s)
	j := 0
	for i := 0; i < len(r)/2; i = i + 1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// this demonstrates how an unexported function
// can be used by an exported function in the same package
