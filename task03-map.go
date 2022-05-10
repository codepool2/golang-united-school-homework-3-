package homework

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here

	result = make([]string, len(input))
	for key, element := range input {
		result[key] = element
	}
	return
}
