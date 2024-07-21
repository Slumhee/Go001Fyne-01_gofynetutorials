package main

import "fmt"

// 定义 Animal 接口
type Animal interface {
    Speak() string
}

// 定义 Dog 结构体
type Dog struct {
    name string
}

// 定义 Cat 结构体
type Cat struct {
    name string
}

// Dog 实现 Animal 接口的 Speak 方法
func (d Dog) Speak() string {
    return "Woof!"
}

// Cat 实现 Animal 接口的 Speak 方法
func (c Cat) Speak() string {
    return "Nya!"
}

// 一个函数，接受 Animal 类型的参数
func MakeAnimalSpeak(a Animal) {
    fmt.Println(a.Speak())
}

func main() {
    dog := Dog{name: "Inu"}
    cat := Cat{name: "Neko"}

    // Dog 和 Cat 类型都实现了 Animal 接口
    MakeAnimalSpeak(dog) // 输出: Woof!
    MakeAnimalSpeak(cat) // 输出: Nya!
}