func generate(numRows int) [][]int {
	res := [][]int{{1}}

	for i := 0; i < numRows-1; i++ {
		temp := append([]int{0}, res[len(res)-1]...)
		temp = append(temp, 0)
		row := []int{}

		for j := 0; j < len(res[len(res)-1])+1; j++ {
			row = append(row, temp[j]+temp[j+1])
		}

		res = append(res, row)
	}

	return res

}