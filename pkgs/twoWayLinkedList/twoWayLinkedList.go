package twoWayLinkedList

import "fmt"

type Object interface{}

type Node struct {
	Data Object
	Pre, Next *Node
}

type List struct {
	headNode *Node
	tailNode *Node
	length int
}

func (this *List) IsEmpty() bool {
	return this.length == 0
}

// 创建: 头插法
func (this *List) Add(data Object) (node *Node) {
	node = &Node{Data: data}
	head := this.headNode
	if head == nil {
		this.headNode = node
		this.tailNode = node
	} else {
		node.Next = head	// 首几点，编程头茬 node 的 Next
		node.Pre = head.Pre
		head.Pre = node	// 当前 node 称为 原 head 的 Pre 
		this.headNode = node
	}
	this.length++
	return
}

// 创建：尾插法
func (this *List) Append(data Object) (node *Node) {
	// 创建节点
	node = &Node(Data: data)
	head := this.headNode
	// 空链表
	if head == nil {
		this.headNode, this.tailNode = node, node
	} else {
		tail := this.tailNode

		// 更新指针
		node.Pre = tail
		node.Next = tail.Next
		tail.Next = node

		// 更新尾节点
		this.tailNode = node
	}
	this.length++
	return
}

func Abs(data int) int {
	if data < 0 {
		return -data
	}
	return data
}

// 删除：指定 index
func (this *List) RmoveAt(index int) bool {

	// 超出长度：默认成功
	if this.length < Abs(index) + 1 {
		return true
	}
	
	if index >= 0 {
		cur := this.headNode
		var count int

		for count < index {
			cur = cur.Next
			count ++
		}
		if cur.Pre != nil {
			cur.Pre.Next = curNext
		}
		if cur.Next != nil {
			cur.Next.Pre = curPre
		}
	} else {
		cur := this.tailNode

		count := 1
		for count < Abs(index) {
			cur = cur.Pre
			count ++
		}
		if cur.Pre != nil {
			cur.Pre.Next = cur.Next
		}
		if cur.Next != nil {
			cur.Next.Pre = cur.Pre
		}
	} 
	this.length++
}

// 删除： 值 value
func (this *List) Rmove(data Object) {
	cur := this.headNode
	for i:=0; i<this.length; i++ {
		if cur.Data == data {
			// 当前位 headNode
			if cur.Pre == nil {
				this.headNode = cur.Next
				this.headNode.Pre = cur.Pre
			} else {
				cur.Pre.Next = cur.Next
				cur.Next.Pre = cur.Pre
			}
			this.length--
		}
	}
}

// 插入
func (this *List) Insert(index int, data Object) bool {

	if this.IsEmpty() {
		this.Append(data)
		return
	}

	node := &Node{Data: data}
	// 头插和尾插
	if index >= 0 {
		count := 0
		cur := this.headNode

		for cur && count < index {
			cur = cur.Next
			count ++
		}

		// TODO 关注一下
		if !cur || count > index {
			return false
		}

		node.Next = cur
		node.Pre = cur.Pre
		cur.Pre = node
		if cur.Pre != nil {
			cur.Pre.Next = node
		}
		this.length++
	} else {
		count := 1
		cur := this.tailNode
		
		for cur && count < Abs(index) {
			cur = cur.Pre
			count ++
		}

		if !cur || count > Abs(index) {
			return false
		}

		node.Next = cur.Next
		node.Pre = cur
		// 更新旧节点
		if cur.Next {
			cur.Next.Pre = node
		}
		cur.Next = node

	}



}

// 取值: 指定 index
func (this *List) GetNodeAt(index int) Object {

	if index > this.length {
		fmt.Println("index 超过 length 不合法")
		return 0
	}

	cur := this.headNode
	for i:=1; i<index; i++ {
		cur = cur.Next
	}

	return cur.Data
}

// 查找: 指定 value 是否存在
func (this *List) Contain(data Object) bool {
	cur := this.headNode
	for i:=0; i<this.length; i++ {
		
		if cur.Data == data {
			return true
		}
		cur = cur.Next
	}
	return false
}

// 打印
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























