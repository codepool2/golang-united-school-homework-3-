package homework

func reverse(input []int64) (result []int64) {

	reverseHelper(input, 0, len(input)-1)

	result = input
	return
}

func reverseHelper(input []int64, startIndex int, endIndex int) {

	for startIndex < endIndex {
		temp := input[startIndex]
		input[startIndex] = input[endIndex]
		input[endIndex] = temp
		startIndex++
		endIndex--
	}

}
