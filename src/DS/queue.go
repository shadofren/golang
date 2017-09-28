package DS

import "fmt";

type Queue struct {
	List []int;
	First int;
	Last int;
}

func (q *Queue) IsFull() bool {
	if (q.First - q.Last == 1 || q.First - q.Last == len(q.List)-1) {
		return true;
	} else {
		return false;
	}
}

func (q *Queue) IsEmpty() bool {
	return q.First == q.Last;
}

func (q *Queue) EnQueue(val int) string{
	if (q.IsFull()) {
		return "Queue is full!";
	} else {
		q.List[q.Last] = val;
		q.Last = (q.Last + 1) % len(q.List);
		return "Successfully added.";
	}
}

func (q *Queue) DeQueue() int {
	if (q.IsEmpty()) {
		return -1;
	} else {
		val := q.List[q.First];
		q.First = (q.First + 1) % len(q.List);
		return val;
	}
}

func TestQueue() {
	queue := Queue{make([]int, 5), 0, 0};

	fmt.Println(queue);
	queue.EnQueue(1);
	fmt.Println(queue);
	queue.EnQueue(2);
	fmt.Println(queue);
	fmt.Println(queue.DeQueue());
	queue.EnQueue(3);
	fmt.Println(queue);
	fmt.Println(queue.DeQueue());
	fmt.Println(queue);
	fmt.Println(queue.DeQueue());
	fmt.Println(queue);
	fmt.Println(queue.DeQueue());

	fmt.Println(queue);
	fmt.Println(queue.EnQueue(4));
	fmt.Println(queue);
	fmt.Println(queue.EnQueue(5));
	fmt.Println(queue);
	fmt.Println(queue.EnQueue(6));
	fmt.Println(queue);
	fmt.Println(queue.EnQueue(7));
	fmt.Println(queue);
	fmt.Println(queue.EnQueue(8));
	fmt.Println(queue);
	fmt.Println(queue.EnQueue(9));
	fmt.Println(queue);
}