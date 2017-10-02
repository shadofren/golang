package others

import "fmt"

func swap(arr []int, x, y int) {
	tmp := arr[x];
	arr[x] = arr[y];
	arr[y] = tmp;
}

func permutation(arr []int, i int, length int) {
	if (length == i) {
		fmt.Println(arr);
	}
	for j := i; j < length; j++ {
		swap(arr, i, j);
		permutation(arr, i+1, length);
		swap(arr, i, j);
	}
}

func TestPermutation() {
	var length int = 5;
	var arr = make([]int, length);
	for i := 0; i < length; i++ {
		arr[i] = i;
	}
	permutation(arr, 0, length);
}