package homework

func average(input [15]float32) (result float32) {

	result = 0

	for _, element := range input {
		result += element
	}

	result = result / float32(len(input))

	return
}
