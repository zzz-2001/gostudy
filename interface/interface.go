package main

import "fmt"

type Usber interface {
	start(string)
	stop(string)
}

type Phone struct {
	Name string
}

func (p Phone) start(a string) {
	fmt.Println(p.Name, a)
}

func (p Phone) stop(a string) {
	fmt.Println(p.Name, a)
}

func main() {
	p := Phone{
		Name: "华为",
	}

	var p1 Usber
	p1 = p

	p1.start("启动！")
	// p1.stop("关机！")

}
