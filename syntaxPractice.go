package main

import (
	"fmt"
	"errors"
	"math"
)

func main() {
	typeOrClass()
//________________________________________________________________

	// result, err := sqrt(-2)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }
//________________________________________________________________
	// getKeyValueWithMap()
//________________________________________________________________

	// getIndexValue()
//________________________________________________________________

	// loop(20)
//________________________________________________________________

	// deleteFromDic()

//________________________________________________________________

	// dic()
//________________________________________________________________

	// slice(15)
//________________________________________________________________

	// array()
//________________________________________________________________

	// result := isBiggerThanTen(10)
	// fmt.Println(result)
//________________________________________________________________

	// sum()
//________________________________________________________________

	// fmt.Println("hello world");
}

func sum () {
	var x int = 5
	y :=  7
	fmt.Println( y + x)
}

func isBiggerThanTen (x int) string {
	if x > 10 {
		return("It's bigger than 10")
	} else if x == 10 {
		return("it's equal 10")
	} else {
		return("it's less than 10")
	}
}

func array () {
	a:= [5]int{1,2,3,4,5}
	fmt.Println(a)
}
func slice (x int) {
	a:= []int{1,2,3,4,5}
	a = append(a, x)
	fmt.Println(a)
}
func dic () {
	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12
	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])
}
func deleteFromDic () {
	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	delete(vertices,"triangle")
	fmt.Println(vertices)
}

func loop (x int) {
	for i := 0; i < x; i++ {
		fmt.Println(i)
	}
}

func getIndexValue () {
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}
}

func getKeyValueWithMap () {
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key:", key, "value", value)
	}
}
func sqrt (x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("the number must be bigger than 0")
	}
	return math.Sqrt(x), nil
}

func typeOrClass () {
	type person struct {
		name string
		age int
	}
	p := person{name:"Su", age:29}
	fmt.Println(p)
	fmt.Println(p.age)
}