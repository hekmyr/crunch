package main

func main() {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {
	var arr []int

	hist := make(map[int]bool)

	for _, n := range nums {
		if !hist[n] {
			hist[n] = true
			println(n)
			arr = append(arr, n)
		}
	}

	copy(nums, arr)

	return len(arr)
}
