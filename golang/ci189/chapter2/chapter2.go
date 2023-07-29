package chapter2

import (
	"fmt"
	"strconv"
	"strings"
)

func Chapter2() {
	c2q1()
	c2q2()
	c2q3()
	c2q4()
	c2q5()
	c2q6()
	c2q7()
	c2q8()
}

/* 連結リスト */
type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	head *Node
}


func (ls *LinkedList) Append(val int) {
	n := Node{}
	n.value = val

	if ls.head == nil {
		ls.head = &n
	} else {
		node := ls.head
		for node.next != nil {
			node = node.next
		}
		node.next = &n
	}
}


func (ls *LinkedList) RemoveAt(index int) {
	if index < 0 || ls.head == nil {
		return
	}
	pre := ls.head
	for i := 0; i < index - 1; i++ {
		pre = pre.next
		if pre == nil || pre.next == nil{
			return
		}
	}

	target := pre.next
	pre.next = target.next
	return
}


func (ls *LinkedList) GetAt(index int) *Node {
	if index < 0 || ls.head == nil {
		return nil
	}
	node := ls.head
	for i := 0; i < index; i++ {
		node = node.next
		if node == nil {
			return nil
		}
	}
	return node
}


func (ls *LinkedList) Display() {
	node := ls.head

	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}


/*
 * Q.2.1
 * 重複要素の削除
*/
func c2q1() {
	fmt.Println(Deduplicate([]int{1,5,3,2,1,4,5,2,3,6}))
}

//for
func Deduplicate0(sl []int) []int {
	ret := []int{}

	for i, x := range(sl) {
		if !Contains(sl[i + 1:], x) {
			ret = append(ret, x)
		}
	}
	return ret
}

//再帰
func Deduplicate(sl []int) []int {
	if len(sl) == 0 {
		return []int{}
	}

	if Contains(sl[1:], sl[0]) {
		return Deduplicate(sl[1:]) 
	}
	return append(Deduplicate(sl[1:]), sl[0])
}


func Contains(sl []int, x int) bool {
	for _, v := range sl {
		if v == x {
			return true
		}
	}
	return false
}


/* 
 * Q.2.2
 * 後ろからK番目を返す
単方向連結リスト
*/
/*
func c2q2() {
	fmt.Println(GetNthFromEnd([]int{1,5,3,2,1,4,5,2,3,6}, 3) == 5)
}

func GetNthFromEnd(sl []int, n int) int {
	return sl[len(sl) - 1 - n]
}
*/

func c2q2() {
	ls := LinkedList{}
	ls.Append(1)
	ls.Append(4)
	ls.Append(2)
	ls.Append(5)
	ls.Append(3)
	fmt.Println(ls.GetAtFromEnd(1).value == 5)
}


func (ls *LinkedList) GetAtFromEnd(index int) *Node {
	if index < 0 {
		return nil
	}
	c := 0
	return ls.getAtFromEnd(ls.head, index, &c)
}


func (ls *LinkedList) getAtFromEnd(node *Node, index int, c *int) *Node {
	if node == nil {
		return nil
	}
	ptr := ls.getAtFromEnd(node.next, index, c)

	*c++

	if *c == index + 1 {
		return node
	}
	return ptr
}


/* 
 * Q.2.3
 * 間の要素を削除
単方向連結リスト
*/
func c2q3() {
	ls := LinkedList{}
	ls.Append(1)
	ls.Append(4)
	ls.Append(2)
	ls.Append(5)
	ls.Append(3)

	ls.Remove(5)
	ls.Display()
}

func (ls *LinkedList) Remove(elem int) {
	node := ls.head
	if node == nil {
		return 
	}

	for node.next != nil {
		if node.next.value == elem {
			node.next = node.next.next
			return
		}
		node = node.next
	}
}


/* 
 * Q.2.4
 * リストの分割
単方向連結リスト
指定した要素より小さい値が前に来るように並び替え
ただし、指定した要素より小さいものが全て前にあれば、大きいものが前にあっても良い
入力: 3->5->8->10->2->1 
出力(例): 3->1->2->10->5->8
*/
func c2q4() {
	ls := LinkedList{}
	ls.Append(3)
	ls.Append(5)
	ls.Append(8)
	ls.Append(10)
	ls.Append(1)
	ls.Append(2)

	ls.SortSplitBy(5)
	ls.Display()
}

func (ls *LinkedList) SortSplitBy(elem int) {
	node := ls.head
	if node == nil {
		return
	}
	
	for node.next != nil {
		next := node.next
		if next.value < elem {
			node.next = next.next
			next.next = ls.head
			ls.head = next
		} else {
			node = node.next
		}
	}
}


/* 
 * Q.2.5
 * リストで表された2数の和
入力 (6->1->7), (2->9->5)  ... 617 + 295
出力 (9->1->2) ... 912
*/
func c2q5() {
	ls1 := LinkedList{}
	ls1.Append(6)
	ls1.Append(1)
	ls1.Append(7)
	
	ls2 := LinkedList{}
	ls2.Append(2)
	ls2.Append(9)
	ls2.Append(5)
	
	ls := SumOfLinkedList(ls1, ls2)
	ls.Display()
}


func SumOfLinkedList(ls1, ls2 LinkedList) LinkedList {
	num := LinkedListToNumber(ls1) + LinkedListToNumber(ls2)
	return NumberToLinkedList(num)
}


func LinkedListToNumber(ls LinkedList) int {
	node := ls.head
	if node == nil {
		return 0
	}

	strnum := ""
	for node != nil {
		strnum += strconv.Itoa(node.value)
		node = node.next
	}
	num, _ := strconv.Atoi(strnum)
	return num
}


func NumberToLinkedList(num int) LinkedList {
	ls := LinkedList{}
	strnums := strings.Split(strconv.Itoa(num), "")

	for _, s := range strnums {
		n, _ := strconv.Atoi(s)
		ls.Append(n)
	}
	return ls
}


/* 
 * Q.2.6
 * 回文
連結リストが回文か判定
*/
func c2q6() {
	ls1 := LinkedList{}
	ls1.Append(1)
	ls1.Append(2)
	ls1.Append(1)
	
	ls2 := LinkedList{}
	ls2.Append(1)
	ls2.Append(2)
	ls2.Append(3)
	
	fmt.Println(IsPalindromeLinkedList(ls1) && !IsPalindromeLinkedList(ls2))
}


func IsPalindromeLinkedList(ls LinkedList) bool {
	l := CloneLinkedList(ls)
	l.ReverseLinkedList()
	return EqualLinkedList(l, ls)
}


func (ls *LinkedList) ReverseLinkedList() {
	pre := ls.head

	if pre == nil {
		return
	}

	node := pre.next
	pre.next = nil
	for node != nil {
		next := node.next
		node.next = pre
		pre = node 
		node = next
	}

	ls.head = pre
}


func CloneLinkedList(origin LinkedList) LinkedList {
	clone := LinkedList{}

	if origin.head == nil {
		return clone
	}

	on := origin.head
	cnpre := &Node{value: on.value}
	clone.head = cnpre

	for on.next != nil {
		on = on.next
		cn := &Node{value: on.value}
		cnpre.next = cn
		cnpre = cn
	}

	return clone
}


func EqualLinkedList(ls1, ls2 LinkedList) bool {
	if ls1.head == nil && ls2.head == nil {
		return true
	}else if ls1.head == nil || ls2.head == nil {
		return false
	}

	return equalLinkedList(ls1.head, ls2.head)
}


func equalLinkedList(n1, n2 *Node) bool {
	if n1.next == nil && n2.next == nil {
		return true
	} else if n1.next == nil || n2.next == nil {
		return false
	} else if n1.next.value != n2.next.value{
		return false
	}

	return equalLinkedList(n1.next, n2.next)
}


/* Q.2.7
 * 共通するノード
 共通はそのノードの参照が一致していること
 1つ目の連結リストのk番目のノードと、2つ目の連結リストのj番目のノードが一致する時、共通
 */
 func c2q7() {

 	node := &Node{value :2}
 	
	ls1 := LinkedList{}
	ls1.head = &Node{value :1}
	ls1.head.next = &Node{value :2}

	ls2 := LinkedList{}
	ls2.head = &Node{value :1}
	ls2.head.next = node

	ls3 := LinkedList{}
	ls3.head = &Node{value :1}
	ls3.head.next = node

	fmt.Println(IsCommonLinkedLists(ls2, ls3) && !IsCommonLinkedLists(ls1, ls2))
}


func IsCommonLinkedLists(ls1, ls2 LinkedList) bool {

	if ls1.head == nil && ls2.head == nil {
		return true
	}else if ls1.head == nil || ls2.head == nil {
		return false
	}

	return isCommonLinkedLists(ls1.head, ls2.head, ls2.head)
}


func isCommonLinkedLists(n1, n2 *Node, n2head *Node) bool {
	if n1 == n2 {
		return true
	}
	if n1 == nil {
		return false
	}
	if n2 == nil {
		return isCommonLinkedLists(n1.next, n2head, n2head)
	}
	
	return isCommonLinkedLists(n1, n2.next, n2head)
}


/* Q.2.8
 * ループの検出
循環を含む連結リストにおいて、
ノードの次へのポインタが以前に出現したノードと一致しているノードを検出
 */
 func c2q8() {
 	n := &Node{value :10}
 	
	ls := LinkedList{}

	ls.head = &Node{value :1}
	n1 := ls.head
	n1.next = &Node{value :3}
	n2 := n1.next
	n2.next = n
	n.next = &Node{value :5}
	n3 := n.next
	n3.next = &Node{value :4}
	n4 := n3.next
	n4.next = n 
	
	fmt.Println(FindBeginningOfCircular(ls).value == 10)
}


func FindBeginningOfCircular(ls LinkedList) *Node {
	head := ls.head
	node := ls.head
	c := 0
	for node != nil {
		n := head
		for i := 0; i < c; i++ {
			if n == node {
				return node
			}
			n = n.next
		}
		c += 1
		node = node.next 
	}
	return nil
}