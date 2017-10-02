package search

import "fmt"

func BinarySearchI(arr []int, val int) int {
	var mid, low, high int;
	low = 0;
	high = len(arr) - 1;
	for low <= high {
		mid = low + (high-low)/2;
		if (arr[mid] == val) {
			return mid;
		} 
		if (arr[mid] > val) {
			high = mid - 1;
		} else {
			low = mid + 1;
		}
	}
	return -1;
}

func TestBinarySearchI() {
	length := 10;
	val := 11;
	var arr = make([]int, length);
	for i := 0; i < length; i++ {
		arr[i] = i;
	}
	fmt.Println(BinarySearchI(arr, val));
}

func BinarySearchR(arr []int, low, high, val int) int {
	var mid int = low + (high-low)/2;
	if (mid == val) {
		return mid;
	}
	if (low == high) {
		return -1;
	}	
	if (arr[mid] > val) {
		return BinarySearchR(arr, low, mid - 1, val);
	} else {
		return BinarySearchR(arr, mid + 1, high, val);
	}
}

func TestBinarySearchR() {
	length := 10;
	val := 11;
	var arr = make([]int, length);
	for i := 0; i < length; i++ {
		arr[i] = i;
	}
	fmt.Println(BinarySearchR(arr, 0, length - 1,val));
}