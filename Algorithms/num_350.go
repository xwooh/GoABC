package main

import "fmt"

/*

给你两个整数数组 nums1 和 nums2，请你以数组形式返回两数组的交集。
返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。
可以不考虑输出结果的顺序。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/intersection-of-two-arrays-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

示例 1：
```
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
```

示例 2:
```
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]
```

*/

func intersect(nums1 []int, nums2 []int) []int {
	/*
	   两个数组取交集，并返回的交集包含最小的相同元素
	   比如：a = [1, 2, 2, 2]; b = [1, 1, 2, 2]
	   最后返回是：[1, 2, 2]

	   解题思路：
	   1. 算出其中一个数组每个元素的个数，构造成一个 `map`
	   2. 新建一个数组作为最后的 `result`
	   3. 遍历另一个数组，如果这个元素存在于上面的 `map` 中 & 对应的 `value > 0` 则添加到 `result` 中
	   4. 返回 `result`
	*/

	var result []int
	var m = make(map[int]int)
	for _, val := range nums1 {
		m[val]++
	}

	for _, val := range nums2 {
		if _, ok := m[val]; ok && m[val] > 0 {
			result = append(result, val)
			m[val]--
		}
	}

	return result
}

func main() {
	// [1, 2, 2]
	fmt.Println(intersect([]int{1, 2, 2, 2}, []int{1, 1, 2, 2}))
}
