package chapter3

func Chapter3() {}


/* スタック */
type Stack struct {
	top *Node
	cap int
	len int
}


type Node struct {
	data int
	min int
	next *Node
}


func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	}
	top := s.top
	s.top = s.top.next
	s.len -= 1
	return top.data, true 
}


func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	}
	return s.top.data, true
}


func (s *Stack) IsEmpty() bool {
	return s.top == nil
}


/*
 * Q.3.2
 * 最小値を返すスタック
ただし、push, pop, minはO(1)
*/

func (s *Stack) Push(item int) {
	n := &Node{data: item}
	n.next = s.top
	if s.IsEmpty() {
		n.min = item
	} else if s.top.min >= item {
		n.min = item
	} else {
		n.min = s.top.min 
	}
	s.len += 1
	s.top = n
} 


func (s *Stack) Min() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	}
	return s.top.min, true
} 


/*
 * Q.3.3
 * 積み上がっている皿
データがいっぱいになったら新しいスタックを作るデータ構想SetOfStacks
*/
type SetOfStacks struct {
	top *StackNode
}

type StackNode struct {
	stack *Stack
	next *StackNode
}


func (ss *SetOfStacks) IsEmpty() bool {
	return ss.top == nil
}

func (ss *SetOfStacks) Peek() (int, bool) {
	if ss.IsEmpty() {
		return -1, false
	}

	return ss.top.stack.Peek()
}

func (ss *SetOfStacks) Push(item int) {
	if ss.IsEmpty() || ss.top.stack.len >= ss.top.stack.cap {
		//1スタックの容量3
		s := &Stack{cap: 3}
		s.Push(item)
		sn := &StackNode{stack: s}
		sn.next = ss.top
		ss.top = sn
	} else {
		ss.top.stack.Push(item)
	}
}


func (ss *SetOfStacks) Pop() (int, bool) {
	if ss.IsEmpty() {
		return -1, false
	}

	ret, ok := ss.top.stack.Pop()
	
	if ss.top.stack.len == 0 {
		ss.top = ss.top.next
	}

	return ret, ok	
}


/* Q.3.4
 * スタックでキュー
2つのスタックを使ってキューを実装
*/
type MyQueue struct {
	stack0 *Stack
	stack *Stack
}


func (q *MyQueue) IsEmpty() bool {
	return q.stack0.IsEmpty() && q.stack.IsEmpty()
}


func (q *MyQueue) Peek() (int, bool)  {
	if q.IsEmpty() {
		return -1, false
	}
	if q.stack.IsEmpty() {
		 q.setMyQueue()
	}

	return q.stack.Peek()
}


func (q *MyQueue) Add(item int) {
	q.stack0.Push(item)
}


func (q *MyQueue) Remove() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}

	if q.stack.IsEmpty() {
		 q.setMyQueue()
	}

	return q.stack.Pop()
}


func (q *MyQueue) setMyQueue() {
	for !q.stack0.IsEmpty() {
		x, _ := q.stack0.Pop()
		q.stack.Push(x)
	}
}