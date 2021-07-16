# leetcode_median_two_sorted_arrays

This repository is for solved the problem for [median-of-two-sorted-arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)

## problem description

Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

## Observation
```
two arrays are both sorted

i.e for i,j in 0..nums1.length-1 , nums1[i] < nums1[j] if i <= j

the median of an array  arr = (arr[Math.floor((arr.length-1)/2)] + arr[Math.floor(arr.length/2)])/2

we could just find out two indexes in nums1, nums2, nums1Idx, nums2Idx

such that half = (nums1.length + nums2.length)/2 , nums1Idx + nums2Idx = half + 2

and nums1[nums1Idx] <= nums2[nums2Idx+1], nums2[nums2Idx] <= nums1[nums1Idx+1]

namely, use binary Search the nums1Idx, nums2Idx

cost O(Min(nums1.length, nums2.length))
```