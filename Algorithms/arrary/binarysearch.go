package arrary

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
