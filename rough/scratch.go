package main

import (
	"fmt"
	"strings"
)

func runAnon(anonFunc func(a, b string) string) {
	fmt.Println(anonFunc("abc", "def"))
}

type Tank interface {
	Volume() float64
}

type Car interface {
	Fuel() float64
}

type myTank struct {
	width  float64
	height float64
}

func (mt myTank) Volume() float64 {
	return mt.width * mt.height
}

func (mt myTank) Fuel() float64 {
	return mt.width * mt.height * 0.5
}

func checkInterface(a interface{}) {
	val, ok := a.(Tank)
	fmt.Println(val, ok)
}

func MainK() {

	mx := myTank{
		width: 10, height: 20,
	}

	var mt Tank = mx
	var mc Car = mx

	fmt.Println(mt.Volume(), mc.Fuel())

	checkInterface(mt)

	anon := func(a, b string) string {
		return a + "<><><><><><><><>" + b
	}

	runAnon(anon)

	arr := []int{2, 10, 4, 3, 1}

	fmt.Println(arr)
	fmt.Println(len(arr))

	fmt.Println(strings.Count("abca", ""))

	var x = 100

	var ptr *int = &x

	*ptr = 200

	fmt.Println(x)
}
