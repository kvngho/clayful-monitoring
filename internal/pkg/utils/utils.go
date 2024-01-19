package utils
func IsEqual(a, b map[string]struct{}) bool {
	if len(a) != len(b) {
		return false
	}
	for key := range a {
		if _, exists := b[key]; !exists {
			return false
		}
	}
	return true
}