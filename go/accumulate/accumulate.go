package accumulate

func Accumulate(collection []string, operation func(c string) string) []string {
	result := []string{}
	for r := range collection {
		result = append(result, operation(collection[r]))
	}
	return result

}
