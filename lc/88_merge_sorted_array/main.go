package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	fmt.Println("Before merge:")
	fmt.Println("nums1:", nums1)
	fmt.Println("nums2:", nums2)

	merge(nums1, m, nums2, n)

	fmt.Println("\nAfter merge:")
	fmt.Println("nums1:", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
    backup := make([]int, len(nums1))
    copy(backup, nums1)

    i1 := 0 
    i2 := 0

    for i := range nums1 {
        
        if i2 == n {
            for j := i; j < len(nums1); j++ {
                nums1[j] = backup[i1]
                i1++
            }
            return
        }
        
        if i1 < m && backup[i1] < nums2[i2] {
            nums1[i] = backup[i1]
            i1++
        } else {
            nums1[i] = nums2[i2]
            i2++
        }
    }
}

