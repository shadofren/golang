package sortings

import "fmt";

func QuickSort(list []int) {
	quicksort(list, 0, len(list)-1);
}

func quicksort(list []int, left int, right int) {
	if (right <= left) {return;}
		
	pivot := partition(list, left, right);
	quicksort(list, left, pivot - 1);
	quicksort(list, pivot + 1, right);
}

func partition(list []int, left int, right int) int {
	pivot_val := list[right];
	i := left-1;
	for j := left; j < right; j++ {
		if (list[j] < pivot_val) {
			i++;
			temp := list[i];
			list[i] = list[j];
			list[j] = temp;
		}
	}
	// swap the pivot with i+1 (first value that is more than pivot value)
	list[right] = list[i+1];
	list[i+1] = pivot_val;
	return i+1;
}

func TestQuickSort(list []int) {
	QuickSort(list);
	fmt.Println("Quick Sort", list);
}