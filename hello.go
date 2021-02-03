package main

import (
    "fmt"
    "./pkgs/pkg2"
    "./pkgs/linkedList"
)

func main() {
    fmt.Println(pkg2.Hello())
    // 初始化 list
    list := linkedList.List{}

    list.Append(1)
    list.Append(2)
    list.Append(3)
    list.Append(4)
    list.Append("a")
    list.Append("b")
    list.Append("c")
    list.Append("d")
    
    fmt.Printf("链表 list 的长度为 %d\n", list.Length())
    fmt.Print("链表List当前数值为：")
    list.ShowList()
    fmt.Println()
    fmt.Println()
    //2. 从头部添加元素
    fmt.Println("注意： =======从头部添加元素，beforeHead")
    fmt.Println()
    list.Add("before")
    list.Append("append")
    
    fmt.Printf("链表 list 的长度为 %d\n", list.Length())
    fmt.Print("链表List当前数值为：")
    list.ShowList()
    fmt.Println()

    // 3. 插入
    fmt.Println("链表 list 第 2 位插入: Rodger")
    list.Insert(1, "Rodger")
    list.ShowList()
    contain := list.Contain("a")
    
    fmt.Println()
    fmt.Println()
    fmt.Println()

    fmt.Println("链表包含 a：", contain)
    list.Remove("b")
    fmt.Println("链表移除 b:", list.Contain("b"))
    list.ShowList()

    fmt.Println()
    fmt.Println("链表移除位置 = 2 的节点：")
    list.RemoveAtIndex(2)
    list.ShowList()

    fmt.Println()
    



}

