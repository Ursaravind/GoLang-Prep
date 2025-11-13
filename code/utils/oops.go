package utils

import (
	"fmt"
	"math"
)

type Book struct {
	Name   string
	Price  float32
	Author string
	Copies int
}

func (book Book) BookDeatails() {

	fmt.Printf("Book name : %s author : %s price : %.2f , copies : %d", book.Name, book.Author, book.Price, book.Copies)

}

type Shape interface {
	Area() float64
	Perimeter() float64
}


type Circle struct{
	Radius float64
}
// type Rectangle struct{
// 	length,width float64
// }

func (c Circle) Area() float64{
	return math.Pi*c.Radius*c.Radius
}

func (c Circle) Perimeter() float64{
	return  2*math.Pi*c.Radius
}
