package main

func main() {
	nums := []int{0,1,2,2,3,0,4,2}
	val := 2
	result := removeElement(nums, val)
	println("Result:", result)
}

func removeElement(nums []int, val int) int {
	var arr []int

	for _, n := range nums {
		if n != val {
			arr = append(arr, n)
		}
	}

	copy(nums, arr)

	return len(arr)
}
