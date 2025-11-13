package main

import (
	"Go-Prep/code/utils"
	"fmt"
)

func init() {
	fmt.Println("hello init function")
}

func main() {
	// it stores the default values of the respective data types in the variables
	// utils.DataTypes()
	// utils.ControlStatements()
	// utils.Loops()
	// utils.SwitchDemo()
	// utils.Display("Hello main")
	// go utils.Display("Hello , Gourotine!")
	// fmt.Println(utils.Split(20))
	// num := 20
	// fmt.Println("before overwritten of num",num)
	// utils.Modify(num)
	// fmt.Println("After overwritten of num",num)
	// defer utils.Mul(5, 2)
	// defer utils.Mul(5,2)
	// defer utils.Mul(7,2)
	// fmt.Println(utils.Myfunc(10, 20))
	// utils.Mul(10, 2)
	// var a utils.Address
	// fmt.Println(a)
	// a1 := utils.Address{Name: "Aravind", City: "Palwancha", State: "Telangana", Mobile: 9347733508, Pincode: 507155}
	// fmt.Println(a1.Name)
	// a1.Name = "jessica"
	// a2 := &utils.Address{Name: "john",City: "tornato",State: "texas",Mobile: 192840127,Pincode: 23412}
	// fmt.Println(a2.City)
	// fmt.Println(a1)

	// utils.Structure()
	// name := utils.GetName()
	// myAddress := utils.Address{
	// 	Name: "Aravind",
	// 	City: "Texas",
	// }
	// name,city := myAddress.GetName()
	// fmt.Println(name,city)
	// utils.ArrayDemo()
	// utils.FilePath()
	// book1 := utils.Book{Name: "Fountain Head",Author: "Ayan Rand",Price: 54.02,Copies:20}
	// book1.BookDeatails()
	var s utils.Shape
	s = utils.Circle{Radius:5}
	fmt.Println(s.Area(),s.Perimeter())

}
