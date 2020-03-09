package stringsu

func Contains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if needle == item {
			return true
		}
	}
	return false
}

func Filter(stack []string, cb func(string) bool) []string {
	result := []string{}
	for _, item := range stack {
		if cb(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map(stack []string, cb func(string) string) []string {
	result := []string{}
	for _, item := range stack {
		result = append(result, cb(item))
	}
	return result
}

func Remove(stack []string, search ...string) []string {
	return Filter(stack, func(s string) bool {
		return !Contains(search, s)
	})
}
