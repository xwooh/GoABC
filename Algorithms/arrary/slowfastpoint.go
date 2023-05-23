package arrary

func RemoveDuplicates(arr []int) int {
	/*
		给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

		考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：

		更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums的其余元素与 nums 的大小不重要。
		返回 k。
		判题标准:

		系统会用下面的代码来测试你的题解:

		int[] nums = [...]; // 输入数组
		int[] expectedNums = [...]; // 长度正确的期望答案

		int k = removeDuplicates(nums); // 调用

		assert k == expectedNums.length;
		for (int i = 0; i < k; i++) {
		    assert nums[i] == expectedNums[i];
		}
		如果所有断言都通过，那么您的题解将被 通过。


		示例 1：

		输入：nums = [1,1,2]
		输出：2, nums = [1,2,_]
		解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
		示例 2：

		输入：nums = [0,0,1,1,1,2,2,3,3,4]
		输出：5, nums = [0,1,2,3,4]
		解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。

		来源：力扣（LeetCode）
		链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-array
		著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
	*/

	if len(arr) <= 1 {
		return len(arr)
	}

	sp := 0
	for _, e := range arr[1:] {
		if e != arr[sp] {
			// 把这个元素挪到 sp 的后一位
			// 然后 sp 往前移
			sp++
			arr[sp] = e
		}
	}

	return sp + 1
}

func MoveZeros(nums []int) {
	/*
		给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

		请注意，必须在不复制数组的情况下原地对数组进行操作。

		示例 1:

		输入: nums = [0,1,0,3,12]
		输出: [1,3,12,0,0]
		示例 2:

		输入: nums = [0]
		输出: [0]


		来源：力扣（LeetCode）
		链接：https://leetcode.cn/problems/move-zeroes
		著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
	*/

	if len(nums) <= 1 {
		return
	}

	sp := 0
	for _, e := range nums {
		// 慢指针：遇到 0 停住
		// 快指针：遇到不为 0 的，将这个元素赋给慢指针所在位置 (如果此时快慢指针不在一起)
		if e != 0 {
			nums[sp] = e
			sp++
		}
	}
	// sp 往后都赋为0
	for ; sp < len(nums); sp++ {
		nums[sp] = 0
	}
}
