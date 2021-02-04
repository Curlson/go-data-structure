package main

import (
    "fmt"
    "./pkgs/pkg2"
    // "./pkgs/twoWayLinkedList"
)

type Object interface{}

func main() {
    fmt.Println(pkg2.Hello())
    // 初始化 list
    // list := twoWayLinkedList{}   

    var (
        goon bool = true
        data Object
        isAppend bool
    )

    for goon {
        var goonStr string
        fmt.Print("请输入节点的值, 使用 enter 结束：")
        fmt.Scanln(&data)

        if isAppend {
            fmt.Println("【头】插法插入", data)
        } else {
            fmt.Println("【尾】插法插入", data)
        }
        fmt.Println("是否继续？yes/no (default: yes)")
        fmt.Scanln(&goonStr)
        if goonStr != "yes" {
            goon = false
        } 
        if !goon {
            fmt.Println("恭喜：输入完成！")
            break
        }
    }


}

