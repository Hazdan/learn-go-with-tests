package iteration

func Repeat(word string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += word
	}

	return repeated
}
