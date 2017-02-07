package main

import (
	"common"
	"common/calc"
	"fmt"
)

func main() {
	fmt.Println(calc.Add2(1, 2))
	fmt.Println(calc.Add3(1, 2, 3))
	fmt.Println(calc.Sub2(1, 2))

	fmt.Println(common.GetValue())
	//a, b := common.GetValue()
	var a, b int = common.GetValue()
	fmt.Println(a, b)

	//_, 空白标识符(占位符)
	c, _ := common.GetValue()
	fmt.Println(c)

	_, f := common.GetValue()
	fmt.Println(f)

	fmt.Println(common.Sum(1, 2, 3))
	fmt.Println(common.Sum(1, 2, 3, 4))

	fmt.Println(common.Fact(3))

}
