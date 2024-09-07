func twoSum(nums []int, target int) []int {
    var res []int

    for i := 0; i < len(nums); i++ {
        
        if len(res) == 2 {
            break;
        }
        for j := i+1; j < len(nums) ; j++ {
            sum := nums[i] + nums[j]

            if (sum) == target {
                res = append(res, i)
                res = append(res, j)
                break;
            }
        }
    }

    return res
}