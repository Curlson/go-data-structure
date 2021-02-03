package linkedList

import "fmt"

type Object interface{}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node
}

// 1. 判空 O(1)
func (this *List) IsEmpty() bool {
	if this.headNode == nil {
		return true
	}
	return false
}

// 2. 求长 O(n)
func (this *List) Length() (count int)  {
	// 获取头节点，原来的头节点不能改变，否则将丢失链表
	cur := this.headNode

	for cur != nil {
		// 头节点不为空，则
		count ++
		// 对地址进行逐个位移
		cur = cur.Next
	}
	return
}

// 3. 头插法 O(1)
func (this *List) Add(data Object) (node *Node) {
	// init 新 节点
	node = &Node{Data: data}
	node.Next = this.headNode
	this.headNode = node

	return
}

// 4. 尾插法	O(n)
func (this *List) Append(data Object) (node *Node) {
	// init node
	node = &Node{Data: data}
	if this.IsEmpty() {
		this.headNode = node
	} else {
		// 获取头部 pointer
		cur := this.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
	return
}

// 插入 O(n)
func (this *List) Insert(index int, data Object) {
	// 小于 0 头插法
	if index <= 0 {
		this.Add(data)
	} else if index > this.Length() {
		// 如果 index > 长度，则进行尾插法
		this.Append(data)
	} else {
		pre := this.headNode
		count :=0
		for count < (index - 1) {
			pre = pre.Next
			count ++
		}

		node := &Node{Data: data}
		// cur 指向 index-1 的位置
		node.Next = pre.Next
		pre.Next = node

	}
	}

	// 删除: 指定 data 的节点 O(n)
	func (this *List) Remove(data Object) {
		pre := this.headNode

		if pre.Data == data {
			this.headNode = pre.Next
		} else {
			for pre.Next != nil {
				if pre.Next.Data == data {
					pre.Next = pre.Next.Next
				} else {
					pre = pre.Next
				}
			}
		}
	}

// 
func (this *List) RemoveAtIndex(index int) {
	pre := this.headNode

	if index <= 0 {
		this.headNode = pre.Next
	} else if index > this.Length() {
		fmt.Println("超出链表长度")
		return
	} else {
		count := 0
		for count != (index-1) && pre.Next != nil {
			count ++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}

func (this *List) Contain(data Object) bool {
	cur := this.headNode

	for cur != nil {
		if cur.Data == data {
			return true
		}
		cur = cur.Next
	}
	return false
}

// 遍历
func (this *List) ShowList() {
	if !this.IsEmpty() {
		cur := this.headNode

		for {
			fmt.Printf("\t%v", cur.Data)
			if cur.Next != nil {
				cur = cur.Next
			} else {
				break
			}
		}
	}
}






















