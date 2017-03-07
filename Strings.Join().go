// Join concatenates the elements of a to create a single string. The separator string
// sep is placed between elements in the resulting string.
func Join(a []string, sep string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}
        // n is the length of the underlying array that will hold the result string
	// total len(a)-1 seps are needed to fill out the gaps
	n := len(sep) * (len(a) - 1)
        // plus total length of all the elements
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

        // allocate n-byte long slice to store the result string
	b := make([]byte, n)

        // copy the first slice element to the result bytes
        // the rest will be copied as sep-element pair
        // bp is the starting cursor for the remaining available space
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
        // convert byte slice to string
	return string(b)
}