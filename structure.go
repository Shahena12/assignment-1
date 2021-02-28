
package main 

import "fmt"
 
type Address struct { 
	Name string 
	city string 
	Pincode int
} 

func main() { 

 
	var a Address 
	fmt.Println(a) 

	 
	a1 := Address{"shahena", "r.knagar", 605007} 

	fmt.Println("Address1: ", a1) 

	
	a2 := Address{Name: "reshma", city: "chennai", 
								Pincode: 605001} 

	fmt.Println("Address2: ", a2) 

	
	a3 := Address{Name: "mumbai"} 
	fmt.Println("Address3: ", a3) 
} 
