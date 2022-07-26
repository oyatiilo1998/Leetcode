package main

import "fmt"

func main() {
	merge([]int{1, 2, 5, 0, 0, 0}, 3, []int{2, 4, 6}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var p1, p2 int = m - 1, n - 1

	for i := n + m - 1; i >= 0; i-- {
		if p2 < 0 {
			break
		}
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[i] = nums1[p1]
			p1--
		} else {
			nums1[i] = nums2[p2]
			p2--
		}
	}
	fmt.Print(nums1)
}
