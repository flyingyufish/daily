package util

// IsHave  对于无序数组判断是否存在某个数组
func IsHave(arr []int, num int) bool {
	for _, a := range arr {
		if a == num {
			return true
		}
	}
	return false
}

// DescBinSearch 降序二分发查找
// descarr 是降序排列的
func DescBinSearch(descarr []int, s int) int {
	floor := 0
	top := len(descarr) - 1
	for {
		if floor > top {
			break
		}
		mid := (floor + top) / 2
		if descarr[mid] > s {
			floor = mid + 1
		} else if descarr[mid] < s {
			top = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// AscBinSearch 升序二分发查找
func AscBinSearch(ascArr []int, num int) int {
	floor := 0
	top := len(ascArr) - 1
	for {
		if floor > top {
			break
		}
		mid := (floor + top) / 2
		if ascArr[mid] < num {
			floor = mid + 1
		} else if ascArr[mid] > num {
			top = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
