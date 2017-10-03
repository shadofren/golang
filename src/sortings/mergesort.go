package sortings

import "fmt"

func sort(arr []int, low, high int) {
	if (low == high) {
		return;
	} else {
		var mid int;
		mid = low + (high - low) / 2;
		sort(arr, low, mid);
		sort(arr, mid+1, high);
		merge(arr, low, mid+1, high+1);
	}
}

func merge(arr []int, i, j, k int) {
	var p1, p2 int;
	p1 = i;
	p2 = j;
	for p1 != p2 && p2 != k{
		if (arr[p1] > arr[p2]) {
			temp := arr[p2];
			for m := p2; m > p1; m-- {
				arr[m] = arr[m-1];
			}
			arr[p1] = temp;
			p2++;
		}
		p1++;
	}
}

func TestMergeSort(list []int) {
	sort(list, 0, len(list)-1);
	fmt.Println(list);
}