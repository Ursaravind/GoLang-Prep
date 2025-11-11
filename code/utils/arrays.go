package utils

import "fmt"

func ArrayDemo() {

	// var scores [4]int
	// scores[0] = 100
	// fmt.Println(scores)
	// one dimensional
	scores := [5]int{10,20,23,30,50}
	names := [3]string{"aravind", "prashanth", "shiva"}
	// fmt.Println(names)
	for index, value := range names {
		fmt.Println(index, value)
	}
	// ingoring the index with blank variable

	for _,score := range scores{
		fmt.Println(score)
	}

	// traditional for loop
	for  i := 0;i<len(names);i++{
		fmt.Println(names[i])
	}

	// multidimensional array

	var arr1 [2][2] int
	arr1[0][0] = 100
	arr1[0][1] = 200

	fmt.Println(arr1)
	arr2 := [2][2]int{{100,200},{0,30}}
	fmt.Println(arr2)

	arr3 := [...]int{20,304,420,2034,23032}
	fmt.Println(arr3)
	fmt.Println(arr1==arr2)

	slice := arr3[1:4]
	fmt.Println(slice) 
	var slice1 = []int{12,3,4,5,6}
	fmt.Println(slice1)

}
