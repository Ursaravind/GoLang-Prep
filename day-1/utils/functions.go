package utils

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