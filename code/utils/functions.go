package utils

import "fmt"

// type 1 declaration
func Add(a int , b int) int{
	return a+b
}
// type 2 declaration
func Sub(a,b int) int{
	return  a-b
}

func Swap(x,y string)(string,string){
	return  y,x
}
func Split(sum int)(x,y int){
	x = sum-10
	y=x-5
	return x,y
}

// anonymous function
func AnonyMous(){
	func (str string){
		fmt.Println(str)
	}("Hello im from anonymous function")

	name := func (name string)  {
		fmt.Println(name)
	}
	name("aravind")
}

// returning anonymous functions
func AnonyMous2() func(name string) string{
	myf := func (name string) string  {
		return  "returning from an"+ name
	}
	return  myf
}

// func Modify(num int){
// 	num = 100
// }

// function with multiple return values
func Myfunc(a,b int)(int ,int ,int,int){
	return  a+b,a-b,a*b,a/b
}