//https://leetcode-cn.com/problems/3sum/
package main

import "fmt"

func main() {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{3, -2, 1, 0}
	// nums := []int{-2, 0, 0, 2, 2}
	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(threeSum(nums))

}

//排序 + 双指针
func threeSum(nums []int) [][]int {
	sortNums(nums)
	returnArray := make([][]int, 0)

	for k := 0; k < len(nums)-2; k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i := k + 1
		j := len(nums) - 1

		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			if i > k+1 && nums[i] == nums[i-1] || sum < 0 {
				i++
				continue
			}
			if j < len(nums)-1 && nums[j] == nums[j+1] || sum > 0 {
				j--
				continue
			}

			if sum == 0 {
				returnArray = append(returnArray, []int{nums[k], nums[i], nums[j]})
				i++
				j--
			}
		}
	}

	return returnArray
}

func sortNums(nums []int) {
	for i := 0; i < len(nums); i++ {
		j := len(nums) - 1
		for i < j {
			if nums[i] > nums[j] {
				tmp := nums[j]
				nums[j] = nums[i]
				nums[i] = tmp
			}
			j--
		}
	}
}

//没做出来
func threeSumMap(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	returnArray := make([][]int, 0)
	mapping := make(map[int]int)
	for _, num := range nums {
		mapping[num] = mapping[num] + 1
	}

	chkMapping := make(map[string]string)

	for mk, mv := range mapping {
		fmt.Println(mk, mv)
	}

	for k, num := range nums {
		if k+1 == len(nums) {
			break
		}

		val := 0 - (num + nums[k+1])

		if _, ok := mapping[val]; ok {
			if (num == val && mapping[num] < 2) || (nums[k+1] == val && mapping[nums[k+1]] < 2) || (num == val && nums[k+1] == val && mapping[num] < 3) {
				continue
			}
			tmpArray := []int{num, nums[k+1], val}
			mKey := sortKey(tmpArray)
			if _, okk := chkMapping[mKey]; !okk {
				chkMapping[mKey] = mKey
				returnArray = append(returnArray, tmpArray)
			}
		}
	}
	return returnArray
}

func sortKey(nums []int) string {
	for i := 0; i < len(nums); i++ {
		j := len(nums) - 1
		for i < j {
			if nums[i] > nums[j] {
				tmp := nums[j]
				nums[j] = nums[i]
				nums[i] = tmp
			}
			j--
		}
	}
	key := ""
	for _, num := range nums {
		key += fmt.Sprintf("%d", num)
	}
	return key
}
