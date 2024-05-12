package do

func IndexOf[T comparable](a []T, e T) int {

	for i, v := range a {
		if v == e {
			return i
		}
	}

	return -1
}
