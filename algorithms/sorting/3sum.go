func threeSum(nums []int) [][]int {
    if len(nums) < 3 { return [][]int{} }
    res, hash := [][]int{}, make(map[[3]int][]int)
    sort.Ints(nums)
    for i := 0; i < len(nums); i++ {
        start, end := i + 1, len(nums) - 1
        for start < end {
            if sum := nums[i] + nums[start] + nums[end]; sum == 0 {
                triplet := [3]int{nums[i], nums[start], nums[end]}
                hash[triplet] = []int{nums[i], nums[start], nums[end]}
                start, end = start + 1, end - 1
            } else if sum < 0 {
                start++
            } else {
                end--
            }
        }
    }
    for _, triplet := range hash {
        res = append(res, triplet)
    }
    return res
}
