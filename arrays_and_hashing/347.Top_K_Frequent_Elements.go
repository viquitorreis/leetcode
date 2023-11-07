func topKFrequent(nums []int, k int) []int {
	hash := make(map[int]int)

	for _, val := range nums {
		hash[val]++
	}

	bucket := make([][]int, len(nums)+1)
	for number, frequency := range hash {
		// colocando no array de array bucket a partir do count -> values
		// index => frequencia
		// dnetro do array => número com aquela quantidade de ocorrências
		bucket[frequency] = append(bucket[frequency], number)
	}

	result := make([]int, 0)
	// como nosso bucket está armazenando os elementos de menor ocorrencia
	// para os de maior ocorrencia, podemos iterar do ultimo elemento do bucket
	// até o primeiro elemento do bucket, inserindo dentro de um array
	// após fazer isso podemos retornar um slice com quantidade de elementos
	// que queremos
	for i := len(bucket) - 1; i > 0 && k > 0; i-- {
		if len(bucket[i]) != 0 {
			result = append(result, bucket[i]...)
		}
	}

	return result[:k]

}
