func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int) 

    for _,v := range nums {
        m[v]++
    }
   var countArr [][]int
	for i, num := range m {
		countArr = append(countArr, []int{i, num})
	}
    //sorting using custom comparator in descending order 
	sort.Slice(countArr, func(i, j int) bool {
		return countArr[i][1] > countArr[j][1]
	})
	var res []int
	for i := 0; i < k; i++ {
		res = append(res, countArr[i][0])
	}
	return res

   
}