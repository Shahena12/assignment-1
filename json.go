package main 

import ( 
	"fmt"
	"encoding/json"
) 

type Human struct{ 
 
	Name string 
	Age int
	Address string 
} 

func main() { 
	
	human1 := Human{"shahena", 23, "puducherry"} 
	
	human_enc, err := json.Marshal(human1) 
	
	if err != nil { 
		
		fmt.Println(err) 
	} 

	fmt.Println(string(human_enc)) 
	
	human2 := []Human{ 
		{Name: "reshma", Age: 23, Address: "chennai"}, 
		{Name: "afreen", Age: 20, Address: "cuddalore"}, 
		{Name: "asrin", Age: 24, Address: "coimbatore"}, 
	} 
	 
	human2_enc, err := json.Marshal(human2) 
		
		if err != nil { 
	
			fmt.Println(err) 
		} 
		 
	fmt.Println() 
		fmt.Println(string(human2_enc)) 
} 
