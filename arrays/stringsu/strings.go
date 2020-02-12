package stringsu

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
