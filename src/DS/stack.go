package DS

import "fmt";

type Stack struct {
	List []int;
	Last int;
}

func (s *Stack) Push(val int) {
	if (len(s.List) > s.Last) {
		s.List[s.Last] = val;
	} else {
		s.List = append(s.List, val);
	}
	s.Last++;
}

func (s *Stack) IsEmpty() bool {
	return s.Last == 0;
}

func (s *Stack) Pop() int {
	if (!s.IsEmpty()) {
		s.Last--;
		return s.List[s.Last];
	}
	return -1;
}

func TestStack() {
	stack := Stack{make([]int, 0), 0};
	(&stack).Push(0);
	fmt.Println(stack.List, len(stack.List), cap(stack.List));
	(&stack).Push(1);
	fmt.Println(stack.List, len(stack.List), cap(stack.List));
	(&stack).Push(2);
	fmt.Println(stack.List, len(stack.List), cap(stack.List));
}