package arrays

func IsUnique(phrase string) bool {
	var letters = map[byte]bool{0x00: false}

	for i := 0; i < len(phrase); i++ {
		if !letters[phrase[i]] {
			letters[phrase[i]] = true
		} else {
			return false
		}
	}
	return true
}
