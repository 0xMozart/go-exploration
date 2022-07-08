func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int, len(nums))

	for _, num := range nums {
		counter[num] += 1
	}

	bubble := make([][]int, len(nums)+1)

	for k, v := range counter {
		bubble[v] = append(bubble[v], k)
	}

	res := make([]int, 0)

	for i := len(nums); i >= 0; i-- {
		for _, ans := range bubble[i] {
			if len(res) == k {
				return res
			}

			res = append(res, ans)
		}
	}

	return res
}
