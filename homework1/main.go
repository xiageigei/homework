package main

import (
	"fmt"
	"sort"
)

// 任务一
func main() {
	// 回文数
	fmt.Println(isPalindrome(123321))

	// 只出现一次的数字
	fmt.Println(singleNumber([]int{1, 2, 3, 4, 5, 5, 4, 3, 1, 2, 6}))

	// 有效的括号
	fmt.Println(isValid("([{}])"))

	// 最长公共前缀
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

	// 加1
	digits := []int{9, 9, 9}
	fmt.Println(plusOne(digits))

	// 删除重复项
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))

	// 合并
	arr := [][]int{
		{4, 5},
		{1, 4},
	}
	fmt.Println(merge(arr))

	// 两数之和
	target := 99
	nums = []int{98, 2, 1}
	fmt.Println(twoSum(nums, target))
}

// 回文数
func isPalindrome(param int) bool {
	if param < 0 {
		return false
	}
	xx := param
	i := 0
	for xx > 0 {
		i = i*10 + xx%10
		xx /= 10
	}
	return param == i
}

// 只出现一次的数字
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	for i := range m {
		if m[i] == 1 {
			return i
		}
	}
	return 0
}

// 有效的括号
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []rune{}

	for _, v := range s {
		switch v {
		case '(', '{', '[':
			// 左括号入栈
			stack = append(stack, v)
		case ')', '}', ']':
			// 右括号：检查栈是否为空或栈顶不匹配
			if len(stack) == 0 || stack[len(stack)-1] != m[v] {
				return false
			}
			// 弹出栈顶
			stack = stack[:len(stack)-1]
		}
	}

	// 栈为空则都已匹配
	return len(stack) == 0
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// 加一
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}

	}

	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}

// 删除重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	q := 0
	p := 1
	for p < len(nums) {
		if nums[q] != nums[p] {
			q++
			nums[q] = nums[p]
		}
		p++
	}
	return q + 1
}

// 合并
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for _, current := range intervals[1:] {
		last := merged[len(merged)-1]
		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}
	return merged
}

// 两数之和
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, x := range nums {
		complement := target - x
		if _, ok := m[complement]; ok {
			return []int{complement, x}
		}
		m[x] = i
	}
	return nil
}
