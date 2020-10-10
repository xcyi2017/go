package mylib

func __merge(arr []int, l int, mid int, r int) {
	var temp = make([]int, 0, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-l] = arr[i]
	}
	i := l
	j := mid + 1
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = temp[j-l]
			j++
		} else if j > r {
			arr[k] = temp[i-l]
			i++
		} else if arr[i-l] < arr[j-l] {
			arr[k] = temp[i-l]
			i++
		} else {
			arr[k] = temp[j-l]
			j++
		}
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func MergeSortBU(arr []int, n int) {
	for sz := 1; sz <= n; sz += sz { // 归并操作的次数，改变sz的大小
		for j := 0; j+sz < n; j += sz + sz {
			// 对arr[j...sz+j-1]和arr[j+sz....j+sz+sz-1]的元素进行交换
			if arr[j] > arr[j+sz] {
				__merge(arr, j, j+sz-1, min(j+sz+sz-1, n-1))
			}
		}
	}
}

func partition(arr []int, l int, r int) int {
	i := l
	for k := l + 1; k <= r; k++ {
		if arr[l] >= arr[k] {
			i++
			arr[i], arr[k] = arr[k], arr[i]
		}
	}
	arr[i], arr[l] = arr[l], arr[i]
	return i
}

func __quickSort(arr []int, l int, r int) {
	if l < r {
		mid := partition(arr, l, r)
		__quickSort(arr, l, mid-1)
		__quickSort(arr, mid+1, r)
	}
}
func QuickSort(arr []int, n int) {
	__quickSort(arr, 0, n-1)
}

//插入排序实现

func Insertition(arr []int, n int) {
	n--
	for i := 1; i <= n; i++ {
		v := arr[i]
		var j int
		for j = i; j >= 1; j-- {
			if arr[j-1] > v {
				arr[j] = arr[j-1]
			} else {
				break
			}
		}
		arr[j] = v
	}
}

// shell sort
func ShellSort(arr []int, n int) {
	n--
	for gap := n / 2; gap >= 1; gap = gap / 2 { // 步长设定
		for i := gap; i <= n; i = i + gap { // 分组
			v := arr[i]
			var j int
			for j = i; j >= 1; j = j - gap {
				if arr[j-gap] > v {
					arr[j] = arr[j-gap]
				} else {
					break
				}
			}
			arr[j] = v
		}
	}
}

// 选择排序
func SelectionSort(arr []int, n int) {
	n--
	for i := 0; i <= n-1; i++ {
		index := i
		for j := i; j <= n-1; j++ {
			if arr[j] > arr[j+1] {
				index = j + 1
			}
		}
		arr[i], arr[index] = arr[index], arr[i]
	}
}

// 冒泡排序，递归版
func BubbleSort(arr []int, l int, r int) {
	if l < r {
		for k := l; k+1 <= r; k++ {
			if arr[k] > arr[k+1] {
				arr[k], arr[k+1] = arr[k+1], arr[k]
			}
		}
		BubbleSort(arr, l, r-1)
	}
}

// 冒泡排序，非递归版，带优化
func BubbleSort_(arr []int, n int) {
	n--
	for i := 0; i < n; i++ {
		flag := true
		for flag {
			for j := 0; j < n-i; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					flag = false
				}
			}
		}
	}
}
