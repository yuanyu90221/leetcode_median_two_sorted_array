package sol

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return SearchMedianSortedArrays(&nums1, &nums2)
	}
	return SearchMedianSortedArrays(&nums2, &nums1)
}

func SearchMedianSortedArrays(large *[]int, small *[]int) float64 {
	largeLen := len(*large)
	smallLen := len(*small)
	if largeLen == 0 {
		return float64(0)
	}
	if smallLen == 0 {
		return float64((*large)[(largeLen-1)/2]+(*large)[largeLen/2]) / 2
	}
	total := largeLen + smallLen
	half := total / 2
	searchLeft := 0
	searchRight := smallLen - 1
	smallIdx := 0
	largeIdx := 0
	smallLeft := math.MinInt32
	smallRight := math.MaxInt32
	largeLeft := math.MinInt32
	largeRight := math.MaxInt32
	for searchLeft <= searchRight+1 {
		if searchLeft+searchRight >= 0 {
			smallIdx = (searchRight + searchLeft) / 2
		} else {
			smallIdx = -1
		}
		largeIdx = half - smallIdx - 2

		if smallIdx >= 0 {
			smallLeft = (*small)[smallIdx]
		} else {
			smallLeft = math.MinInt32
		}
		if smallIdx+1 < smallLen {
			smallRight = (*small)[smallIdx+1]
		} else {
			smallRight = math.MaxInt32
		}
		if largeIdx >= 0 {
			largeLeft = (*large)[largeIdx]
		} else {
			largeLeft = math.MinInt32
		}
		if largeIdx+1 < largeLen {
			largeRight = (*large)[largeIdx+1]
		} else {
			largeRight = math.MaxInt32
		}
		if smallLeft <= largeRight && largeLeft <= smallRight {
			if total%2 == 1 {
				return float64(Min(smallRight, largeRight))
			} else {
				return float64(Min(smallRight, largeRight)+Max(smallLeft, largeLeft)) / 2
			}
		} else if smallLeft > largeRight {
			searchRight = smallIdx - 1
		} else if smallRight < largeLeft {
			searchLeft = smallIdx + 1
		}
	}
	return float64(0) // not found
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
