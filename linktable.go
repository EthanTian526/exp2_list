package main

import "fmt"

// 定义节点
type ListNode struct {
	data interface{}
	next *ListNode
}

// 定义链表
type LinkTable struct {
	length   int
	headNode *ListNode
}

// 带有头节点的单链表的初始化
func InitList() *LinkTable {
	node := new(ListNode)
	L := new(LinkTable)
	L.headNode = node
	L.length = 0
	return L
}

// 判断链表是否为空
func (list *LinkTable) IsNull() bool {
	if list.length == 0 {
		return true
	} else {
		return false
	}
}

// 得到链表长度
func (list *LinkTable) GetLength() int {
	return list.length
}

// 遍历链表
func (list *LinkTable) ShowList() {
	if !list.IsNull() {
		cur := list.headNode.next
		for {
			fmt.Printf("%v\t", cur.data)
			if cur.next != nil {
				cur = cur.next
			} else {
				break
			}
		}
		fmt.Printf("\n")
	}
}

// 查找指定值是否存在于链表中
func (list *LinkTable) IsInElem(v interface{}) bool {
	if list.IsNull() {
		fmt.Println("err")
		return false
	} else {
		pre := list.headNode
		for pre != nil {
			if pre.data == v {
				return true
			}
			pre = pre.next
		}
		return false
	}
}

// 单链表头插法
func (list *LinkTable) AddElem(v interface{}) {
	node := &ListNode{data: v}
	node.next = list.headNode.next
	list.headNode.next = node
	list.length++
	return
}

// 单链表尾插法
func (list *LinkTable) AppendElem(v interface{}) {
	node := &ListNode{data: v}
	cur := list.headNode
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = node
	list.length++
	return
}

// 删除指定值结点
func (list *LinkTable) RemoveElem(v interface{}) {
	pre := list.headNode
	for pre.next != nil {
		if pre.next.data == v {
			pre.next = pre.next.next
			list.length--
			return
		} else {
			pre = pre.next
		}
	}
	fmt.Println("fail")
	return
}

func main() {
	L := InitList()
	msg := []int{1, 2, 3, 4, 5, 6}
	// 利用尾插法由数组创建链表
	for i := range msg {
		L.AppendElem(msg[i])
	}
	L.ShowList()
	fmt.Println(L.GetLength())
	// 头插法插入结点
	L.AddElem(0)
	L.ShowList()
	// 删除指定结点
	L.RemoveElem(2)
	L.ShowList()
	// 查询结点是否在链表中
	fmt.Println(L.IsInElem(1))
	fmt.Println(L.GetLength())
}
