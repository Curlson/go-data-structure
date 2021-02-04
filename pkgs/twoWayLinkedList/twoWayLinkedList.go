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

func (this *List) Length() int {
	return this.length
}

func (this *List) IsEmpty() bool {
	if this.headNode == nil {
		return true
	}
	return false
}

// 创建: 头插法
func (this *List) Add(data Object) (node *Node) {
	node = &Node{Data: data}
	if this.headNode == nil {
		this.tailNode = node
	} 
	node.Next, node.Pre= this.headNode, this.headNode.Pre	// 首几点，编程头茬 node 的 Next
	this.headNode.Pre = node	// 当前 node 称为 原 head 的 Pre 
	this.headNode = node

	this.length++
	return
}

// 创建：尾插法
func (this *List) Append(data Object) (node *Node) {
	// 创建节点
	node = &Node{Data: data}
	tail := this.tailNode

	// 空链表
	if this.IsEmpty() {
		this.headNode = node
	} else {
		// 更新指针
		node.Pre, node.Next = tail, tail.Next
		tail.Next = node
	}

	this.tailNode = node
	this.length++

	return
}

// 删除：指定 index
func (this *List) RmoveAt(index int) bool {
	// 超出长度：默认成功
	if this.length < Abs(index) || index == 0 {
		fmt.Println("index 超出范围 或者 不能为 0")
		return false
	}
	
	var (
		count, pos int
		cur *Node = this.headNode
	)

	if index > 0 {
		pos = index - 1
	} else {
		pos = this.length + index
	}

	for cur != nil &&  count < pos {
		cur = cur.Next
		count ++
	}
	
	if cur.Pre != nil {
		cur.Pre.Next = cur.Next
	}
	if cur.Next != nil {
		cur.Next.Pre = cur.Pre
	}

	this.length--
	return true
}

// 删除： 值 value
func (this *List) Rmove(data Object) bool {
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
	return true
}

// 插入
func (this *List) Insert(index int, data Object) bool {

	if this.IsEmpty() {
		this.Append(data)
		this.length++
		return true
	}

	var (
		count, pos int
		cur *Node = this.headNode
		node *Node = &Node{Data: data}
	)

	
	if index > 0 {
		pos = index - 1
	} else {
		pos = this.length + index
	}

	// 头插和尾插
	for cur != nil && count < pos {
		cur = cur.Next
		count ++
	}

	// TODO 关注一下
	if cur == nil {
		return false
	}

	node.Next = cur
	node.Pre = cur.Pre
	if cur.Pre != nil {
		cur.Pre.Next = node
	} else {
		this.headNode = node
	}
	// 只能放最后
	cur.Pre = node
	this.length++
	return true
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

	var cur *Node = this.headNode

	for cur != nil{
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

func Abs(data int) int {
	if data < 0 {
		data = -data
	}
	return data
}