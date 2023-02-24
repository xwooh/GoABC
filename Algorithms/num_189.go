package main

func rotate(nums []int, k int) {
	/*
	   一次次操作的步骤:
	       1. 取出最后一位数
	       2. 从 n-1 位数开始，依次往后移动一位
	       3. 第 1 位数移动完成之后将上面取出的最后一位数放在第 1 位上

	       length := len(nums)
	       for i := 0; i < k; i++ {
	           tmp := nums[length - 1]
	           for pos := length - 2; pos >= 0; pos-- {
	               nums[pos + 1] = nums[pos]
	           }
	           nums[0] = tmp
	       }

	   直接一次性操作的步骤:
	       1. 取出末尾的 k 个元素
	       2. 从 n-k 位元素开始依次往后移动 k 位
	       3. 移动完成之后将之前取出的 k 个元素按原本顺序填充到前面 k 个空位上
	*/

	length := len(nums)

	// 避免旋转次数超过当前数组长度
	k = k % length

	// 取出末尾的 k 个元素
	var tmp []int
	for i := k - 1; i >= 0; i-- {
		tmp = append(tmp, nums[length-1-i])
	}

	for i := length - 1 - k; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	for idx, v := range tmp {
		nums[idx] = v
	}
}

func main() {
	rotate([]int{-1, 100, 3, 99}, 5)
}
