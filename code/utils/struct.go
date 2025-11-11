package utils

// import "fmt"

type Address struct {
	Name    string
	City    string
	State   string
	Mobile  int64
	Pincode int
}

// func Structure() {
// 	person := Address{
// 		Name:    "Ancy",
// 		Age:     21,
// 		Address: "ascinx",
// 	}
// 	fmt.Println(person)
// }

func (p Address) GetName() (string,string){
	return p.Name,p.City
}