package main

func main() {
	nums := []int{0,0,1,1,1,1,2,3,3}
	removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {
	var arr []int

	record := make(map[int]int)

	for _, n := range nums {
		if record[n] < 2 {
			arr = append(arr, n)
			record[n]++
		}
	}

	copy(nums, arr)

	return len(arr)
}
