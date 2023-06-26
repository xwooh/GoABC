package arrary

import "fmt"

func SearchTheLeftTarget(nums []int, target int) int {
	// 寻找左侧边界的二分搜索
	// 比如 nums: 1 2 2 3 4 5, target: 2
	// 此时的返回应该是 1

	left, right := 0, len(nums)-1

	// 探索边界为 [left, right]
	for left <= right {
		// 防止 left 和 right 过大时，left + right 越界
		mid := left + (right-left)/2
		if nums[mid] < target {
			// 边界左移
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 因为要找最左侧的，所以找到一个之后要左移
			right = mid - 1
		}
	}

	// 如果 for 循环遍历完都没找到，此时 left = right + 1，而 right 为 len(nums) - 1
	// 此时 left = len(nums)，下面 nums[left] 会越界，这里需要提前判断下
	if left > len(nums)-1 {
		return -1
	}

	// 退出时的条件是 left = right + 1
	// 而上面 if nums[mid] == target 时，mid = right + 1
	//     所以此时 left == mid
	// 但是因为 nums[mid] > target 的时候，mid = right + 1
	//     所以还要再做次判断
	if nums[left] == target {
		return left
	}

	return -1
}

func SearchTheRightTarget(nums []int, target int) int {
	left, right := 0, len(nums)

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 因为要找最右侧的，所以找到一个之后要右移
			left = mid + 1
		}
	}

	// 因为要返回 right
	// right 的极端情况是一直往左移，此时 for 循环退出时 right = left - 1  => 0 - 1
	// 此时 right 是 -1 了，为了防止下面 nums[right] ?= target 的越界，需要提前判断下
	if right < 0 {
		return -1
	}

	// 退出时 left = right + 1
	// 找到时 left = mid + 1
	//   => right = mid
	if nums[right] == target {
		return right
	}

	return -1
}

func ShipWithinDays(weights []int, days int) int {
	/*
		传送带上的包裹必须在 days 天内从一个港口运送到另一个港口。

		传送带上的第 i个包裹的重量为weights[i]。每一天，我们都会按给出重量（weights）的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。

		返回能在 days 天内将传送带上的所有包裹送达的船的最低运载能力。


		示例 1：

		输入：weights = [1,2,3,4,5,6,7,8,9,10], days = 5
		输出：15
		解释：
		船舶最低载重 15 就能够在 5 天内送达所有包裹，如下所示：
		第 1 天：1, 2, 3, 4, 5
		第 2 天：6, 7
		第 3 天：8
		第 4 天：9
		第 5 天：10

		请注意，货物必须按照给定的顺序装运，因此使用载重能力为 14 的船舶并将包装分成 (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) 是不允许的。
		示例 2：

		输入：weights = [3,2,2,4,1,4], days = 3
		输出：6
		解释：
		船舶最低载重 6 就能够在 3 天内送达所有包裹，如下所示：
		第 1 天：3, 2
		第 2 天：2, 4
		第 3 天：1, 4
		示例 3：

		输入：weights = [1,2,3,1,1], days = 4
		输出：3
		解释：
		第 1 天：1
		第 2 天：2
		第 3 天：3
		第 4 天：1, 1


		来源：力扣（LeetCode）
		链接：https://leetcode.cn/problems/capacity-to-ship-packages-within-d-days
		著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
	*/

	// 解析：要想把货成功运出去，那传送带的承载量必然在 [max(weights), sum(weights)] 这个区间里
	// 		 这样一想，这道题就转变为：在 [max(weights), sum(weights)] 这个升序区间内查找一个值（传送带承载量），
	//		 						能满足在 days 内成功把这批货运出去。
	//		 	题目要求找到最小承载量，所以可以通过二分查找法查找最左边界值。

	// 判断 t 承载量能否在 days 内运送完
	check := func(t int, days int) int {
		// bool 是否正好 days 运完
		fmt.Printf("===========\nt=%d, days=%d\n", t, days)
		days--
		left := t
		for _, e := range weights {
			if left >= e {
				// 还能运
				left -= e
				fmt.Printf("%d ", e)
			} else {
				// 不能运了，转到下一天
				fmt.Printf("\n")
				fmt.Printf("%d ", e)
				days--
				left = t  // 恢复运载量
				left -= e // 并运掉这批货
			}
		}
		fmt.Printf("\n=====剩余天数: %d=====\n", days)

		return days
	}

	max := func(ns []int) int {
		var m int
		for _, e := range ns {
			if e > m {
				m = e
			}
		}
		return m
	}
	sum := func(ns []int) int {
		var s int
		for _, e := range ns {
			s += e
		}
		return s
	}

	left, right := max(weights), sum(weights)
	for left <= right {
		mid := left + (right-left)/2

		r := check(mid, days)
		if r > 0 {
			// 还剩天数，说明运多了，mid 降降
			right = mid - 1
		} else if r < 0 {
			// 不行的话，意味着 mid 要提一提
			left = mid + 1
		} else if r == 0 {
			// 正好运完， mid 尝试降降
			right = mid - 1
		}
	}

	// left = right + 1
	// mid = right + 1
	if left > sum(weights) {
		return -1
	}

	if check(left, days) >= 0 {
		return left
	}

	return -1
}
