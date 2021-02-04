package main

import (
    "fmt"
    "./pkgs/pkg2"
    "./pkgs/twoWayLinkedList"
)

type Object interface{}

func main() {
    fmt.Println(pkg2.Hello())
    // 初始化 list
    list := twoWayLinkedList.List{}   


    var (
        data string
        goonStr string = "y"
        goon bool = true
        addType int
        index int
    )

    

    for goon {
        fmt.Print("请输入节点的值, 继续 ( rodger  0  y), 使用空格结束：")

        fmt.Scanln(&data, &addType, &goonStr)

        fmt.Printf("输入了 %v 和 %v", data, goonStr)

        // 顺序插入：尾差法：追加
        if addType == 0{
            list.Append(data)
        } else if addType == 1{
            list.Add(data)
        } else {
            // 插入
            fmt.Println("请输入，新元素要插入的位置（放到第一个位置就填 1， 倒数第一个，就写 -1）")
            fmt.Scanln(&index)
            res := list.Insert(index, data)
            fmt.Println("插入结果：", res)
        }
        fmt.Println("最新链表如下： length= ", list.Length())
        list.ShowList()
        fmt.Println("\n========================")
        if goonStr != "y" {
            goon = false
        } 
        if !goon {
            fmt.Println("恭喜：输入完成！")
            break
        }
    }

}

