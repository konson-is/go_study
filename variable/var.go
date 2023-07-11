package main

import (
	"fmt"
	"os/user"
	"time"
)

func init() {
	fmt.Println("Init!")
}

func bazz() {
	fmt.Println("Bazz")
}

func foo() {
	// ショートの場合関数の外では宣言できない
	xi := 1
	var xf32 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xi)

}

func main() {
	//bazz()
	fmt.Println("Hello world", time.Now())
	fmt.Println(user.Current())

	// 変数宣言
	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "test"
		t, f bool    = true, false
	)
	fmt.Println(i, f64, s, t, f)
	foo()

}
