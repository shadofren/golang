package sortings


func HeapSort(list []int) {
	var size int;
	size = len(list);
	build_max_heap(list, size);
	for i := size - 1; i > 0; i-- {
		// swap the largest item to the back
		temp := list[0];
		list[0] = list[i];
		list[i] = temp;
		
		// re max-heapify
		max_heapify(list, 0, i);
	}
}

func build_max_heap(list []int, size int ) {
	var index int;
	index = size / 2;
	for index >= 0 {
		max_heapify(list, index, size);
		index--;
	}
}

func max_heapify(list []int, pos int, size int) {
	left := pos * 2 + 1;
	right := left + 1;
	
	largest := pos;
	
	if (left < size && list[largest] < list[left]) {
		largest = left;
	}

	if (right < size && list[largest] < list[right]) {
		largest = right;
	}

	if (largest != pos) {
		temp := list[pos];
		list[pos] = list[largest];
		list[largest] = temp;
		max_heapify(list, largest, size);
	}
}